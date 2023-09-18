package inmemory

func (s *Storage) Set(key string, val any) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[key] = val
}
