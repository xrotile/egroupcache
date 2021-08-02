package slowdb

import (
	"sync"
	"time"
)

type DB struct {
	mu sync.Mutex
	db map[string]string
}

func (db *DB) Put(key string, value string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	time.Sleep(time.Second * 2)
	db.db[key] = value
}

func (db *DB) Get(key string) string {
	db.mu.Lock()
	defer db.mu.Unlock()
	return db.db[key]
}
