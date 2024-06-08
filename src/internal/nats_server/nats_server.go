package nats_server

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"go-nuts/database"
	"go-nuts/internal/cache"
	"log"
	"sync"
)

type Nats struct {
	Con  *nats.Conn
	Chan chan *nats.Msg
}

func New() Nats {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	return Nats{Con: nc, Chan: make(chan *nats.Msg, 64)}
}

func (nt Nats) Subscribe() {
	if _, err := nt.Con.ChanSubscribe("foo", nt.Chan); err != nil {
		log.Fatal(err)
	}
}

func (nt Nats) Reader(s *sync.Mutex, cache *cache.Cache, db database.Database) {
	serializedData := make(map[string]interface{})
	for i := range nt.Chan {
		s.Lock()
		data := i.Data
		if err := json.Unmarshal(data, &serializedData); err != nil {
			log.Fatalf("Some errors with data format %s", err.Error())
		}
		db.Insert(serializedData["order_uid"].(string), data)
		cache.Set(serializedData["order_uid"].(string), serializedData)
		fmt.Printf("Received a message: %s\n", serializedData)
		s.Unlock()
	}
}
