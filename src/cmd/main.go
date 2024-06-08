package main

import (
	"go-nuts/database"
	"go-nuts/internal/cache"
	"go-nuts/internal/config"
	"go-nuts/internal/http_server"
	"go-nuts/internal/nats_server"
	"sync"
)

func main() {
	cfg := config.MustLoad()
	cache := cache.New()
	db := database.New(&cfg.DBConfig)
	cache.LoadData(db)

	nats := nats_server.New()
	go nats.Subscribe()
	go nats.Reader(&sync.Mutex{}, cache, db)

	// TODO http handler
	server := http_server.New(cfg.HttpServerConfig, cache)
	server.Run()
}

// TODO Dockerfile
// TODO README
// TODO Hide secrets
// TODO Make a video report
