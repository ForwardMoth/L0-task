package cache

import (
	"encoding/json"
	"fmt"
	"go-nuts/database"
	"log"
	"sync"
)

type Cache struct {
	sync.RWMutex
	Data map[string]interface{}
}

func New() *Cache {
	return &Cache{Data: map[string]interface{}{}}
}

func (c *Cache) Set(uuid string, data interface{}) {
	c.Lock()
	c.Data[uuid] = data
	c.Unlock()
}

func (c *Cache) Get(uuid string) (interface{}, error) {
	c.RLock()
	defer c.RUnlock()
	if _, ok := c.Data[uuid]; !ok {
		return nil, fmt.Errorf("data with uuid %v isn't found", uuid)
	}
	return c.Data[uuid], nil
}

func (c *Cache) LoadData(db database.Database) {
	data, err := db.Select()
	if err != nil {
		log.Fatal(err)
	}

	serializedData := make(map[string]interface{})
	for _, row := range data {
		if err = json.Unmarshal(row, &serializedData); err != nil {
			log.Fatalf("error unmarshal json %v", err.Error())
		}
		c.Set(serializedData["order_uid"].(string), serializedData)
	}
	log.Print("loading cache is done...")
}
