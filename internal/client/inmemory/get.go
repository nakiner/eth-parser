package inmemory

func (s *Storage) Get(key string) (any, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.store[key]
	return val, ok
}
