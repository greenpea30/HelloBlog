package main

import (
	"flag"
	"log"
	"net/http"

	"helloblog/internal/config"
	"helloblog/internal/infra/db"
	redisinfra "helloblog/internal/infra/redis"
	"helloblog/internal/mcp"
	"helloblog/internal/svc"
)

func main() {
	configPath := flag.String("config", "etc/config.yaml", "path to config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	database, err := db.NewPostgres(cfg.Database.DSN)
	if err != nil {
		log.Fatalf("connect postgres: %v", err)
	}

	redisClient := redisinfra.New(cfg.Redis.Addr, cfg.Redis.Password, cfg.Redis.DB)
	if redisClient != nil {
		defer redisClient.Close()
	}

	svcCtx := svc.NewServiceContext(cfg, database, redisClient)
	handler := mcp.NewHTTPHandler(svcCtx, cfg.MCP)

	addr := ":8081"
	log.Printf("[mcp] MCP Server listening on %s", addr)
	log.Printf("[mcp] Tools: health, list_posts, get_post, search_posts")
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("run mcp server: %v", err)
	}
}
