package main

import (
	"flag"
	"log"

	"helloblog/internal/config"
	"helloblog/internal/infra/db"
	redisinfra "helloblog/internal/infra/redis"
	"helloblog/internal/server"
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

	httpServer := server.New(cfg, database, redisClient)
	if err := httpServer.Run(); err != nil {
		log.Fatalf("run http server: %v", err)
	}
}
