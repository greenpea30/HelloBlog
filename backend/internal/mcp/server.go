package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"helloblog/internal/config"
	"helloblog/internal/dto"
	"helloblog/internal/svc"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func NewServer(svcCtx *svc.ServiceContext) *mcp.Server {
	server := mcp.NewServer(&mcp.Implementation{Name: "helloblog-mcp", Version: "0.1.0"}, nil)
	registerTools(server, svcCtx)
	return server
}

func NewHTTPHandler(svcCtx *svc.ServiceContext, cfg config.MCPConfig) http.Handler {
	server := NewServer(svcCtx)
	handler := mcp.NewStreamableHTTPHandler(func(*http.Request) *mcp.Server { return server }, &mcp.StreamableHTTPOptions{Stateless: true, JSONResponse: true})
	if len(cfg.AllowedOrigins) > 0 {
		origins := cfg.AllowedOrigins
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			for _, o := range origins {
				if o == origin || o == "*" { w.Header().Set("Access-Control-Allow-Origin", origin); break }
			}
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			if r.Method == http.MethodOptions { w.WriteHeader(204); return }
			handler.ServeHTTP(w, r)
		})
	}
	return handler
}

// ---------- Tools ----------

type healthArgs struct{}
type listPostsArgs struct { Page, PerPage int; OrderBy string }
type getPostArgs struct { ID int64 }
type searchArgs struct { Query string; PerPage int }

func registerTools(server *mcp.Server, svcCtx *svc.ServiceContext) {
	mcp.AddTool(server, &mcp.Tool{Name: "health", Description: "Check if HelloBlog MCP server is running"},
		func(ctx context.Context, req *mcp.CallToolRequest, _ healthArgs) (*mcp.CallToolResult, any, error) {
			return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: `{"status":"ok","service":"helloblog"}`}}}, nil, nil
		})

	mcp.AddTool(server, &mcp.Tool{Name: "list_posts", Description: "List blog posts with pagination"},
		func(ctx context.Context, req *mcp.CallToolRequest, args listPostsArgs) (*mcp.CallToolResult, any, error) {
			if args.Page <= 0 { args.Page = 1 }
			if args.PerPage <= 0 { args.PerPage = 10 }
			r := dto.PostListRequest{Page: args.Page, PageSize: args.PerPage, OrderBy: args.OrderBy}
			resp, err := svcCtx.Post.List(r)
			if err != nil { return nil, nil, err }
			return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: toJSON(resp)}}}, nil, nil
		})

	mcp.AddTool(server, &mcp.Tool{Name: "get_post", Description: "Get a blog post by ID"},
		func(ctx context.Context, req *mcp.CallToolRequest, args getPostArgs) (*mcp.CallToolResult, any, error) {
			resp, err := svcCtx.Post.GetByID(args.ID)
			if err != nil { return nil, nil, err }
			return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: toJSON(resp)}}}, nil, nil
		})

	mcp.AddTool(server, &mcp.Tool{Name: "search_posts", Description: "Full-text search blog posts"},
		func(ctx context.Context, req *mcp.CallToolRequest, args searchArgs) (*mcp.CallToolResult, any, error) {
			if args.PerPage <= 0 { args.PerPage = 10 }
			r := dto.SearchRequest{Query: args.Query, PageSize: args.PerPage}
			resp, err := svcCtx.Search.FullTextSearch(r)
			if err != nil { return nil, nil, err }
			return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: toJSON(resp)}}}, nil, nil
		})

	log.Println("[mcp] Registered tools: health, list_posts, get_post, search_posts")
}

func toJSON(v any) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil { return fmt.Sprintf(`{"error":"%v"}`, err) }
	return string(b)
}
