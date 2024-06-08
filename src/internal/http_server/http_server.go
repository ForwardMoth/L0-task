package http_server

import (
	"encoding/json"
	"go-nuts/internal/cache"
	"go-nuts/internal/config"
	"log"
	"net/http"
)

const PageNotFound = "404 Page is not found"

type Server struct {
	S     *http.Server
	Cache *cache.Cache
}

func New(cfg config.HttpServerConfig, cache *cache.Cache) Server {
	return Server{
		S: &http.Server{
			Addr:        cfg.Address,
			ReadTimeout: cfg.Timeout,
			IdleTimeout: cfg.IdleTimeout,
		},
		Cache: cache,
	}
}

func (s Server) Run() {
	http.HandleFunc("/order", s.GetOrder) // order?id=
	if err := s.S.ListenAndServe(); err != nil {
		log.Fatalf("error with http server connection %v", err.Error())
	}
	log.Print("Starting server on :8080")
}

func (s Server) GetOrder(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if data, err := s.Cache.Get(id); err != nil {
		w.Write([]byte(PageNotFound))
	} else {
		body, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}
	log.Printf("GET /order with id %v", id)
}
