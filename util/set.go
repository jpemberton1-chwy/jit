package util

var exists = struct{}{}

type Set struct {
	m map[interface{}]interface{}
}

func NewSet() *Set {
	s := &Set{}
	s.m = make(map[interface{}]interface{})
	return s
}

func (s *Set) Add(value interface{}) *Set {
	s.m[value] = exists
	return s
}

func (s *Set) Remove(value interface{}) *Set {
	delete(s.m, value)
	return s
}

func (s *Set) Contains(value interface{}) bool {
	_, c := s.m[value]
	return c
}

func (s *Set) AsSlice() []interface{} {
	keys := []interface{}{}
	for key := range s.m {
		keys = append(keys, key)
	}
	return keys
}
