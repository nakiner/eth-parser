package inmemory

import "sync"

type Storage struct {
	store map[string]any
	mu    sync.RWMutex
}

func New() *Storage {
	return &Storage{
		store: make(map[string]any),
	}
}
