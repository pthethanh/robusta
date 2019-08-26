package limiter

import "time"

type (
	MemStore struct {
		visitors map[string]*Visitor
	}
)

func NewMemStore() *MemStore {
	return &MemStore{
		visitors: make(map[string]*Visitor),
	}
}

func (s *MemStore) Set(k string, v *Visitor) error {
	s.visitors[k] = v
	return nil
}

func (s *MemStore) Get(k string) (*Visitor, error) {
	return s.visitors[k], nil
}

func (s *MemStore) Del(k string) error {
	delete(s.visitors, k)
	return nil
}

func (s *MemStore) Clean(maxAge time.Duration) error {
	for ip, v := range s.visitors {
		if time.Now().Sub(v.lastSeen) > maxAge {
			_ = s.Del(ip)
		}
	}
	return nil
}
