package mcp

import (
	"net/http"

	"helloblog/internal/config"
	"helloblog/internal/svc"
)

// Server 是 MCP Server 的包装
// 详细实现需要引入 modelcontextprotocol/go-sdk
// 这里提供基础骨架，预留扩展
type Server struct {
	svcCtx *svc.ServiceContext
}

func NewServer(svcCtx *svc.ServiceContext) *Server {
	return &Server{svcCtx: svcCtx}
}

func NewHTTPHandler(s *Server, cfg config.MCPConfig) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		postgresStatus := "ok"
		sqlDB, err := s.svcCtx.DB.DB()
		if err != nil || sqlDB.Ping() != nil {
			postgresStatus = "error"
		}

		redisStatus := "disabled"
		if s.svcCtx.Redis != nil {
			redisStatus = "ok"
			if err := s.svcCtx.Redis.Ping(r.Context()).Err(); err != nil {
				redisStatus = "error"
			}
		}

		w.Write([]byte(`{"postgres":"` + postgresStatus + `","redis":"` + redisStatus + `"}`))
	})

	return mux
}
