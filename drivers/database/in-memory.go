package database

import "sync"

type InMemoryDatabase struct {
	Map sync.Map
}

func (receiver *InMemoryDatabase) Set(key, value string) {
	receiver.Map.Store(key, value)
}

func (receiver *InMemoryDatabase) Get(key string) (any, bool) {
	return receiver.Map.Load(key)
}
