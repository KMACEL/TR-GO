package singleton

import "sync"

type secureSingleton struct {
	sayac int
	sync.RWMutex
}

var secureInstance *secureSingleton

func SecureGetInstance() *secureSingleton {
	if secureInstance == nil {
		secureInstance = new(secureSingleton)
	}
	return secureInstance
}

func (s *secureSingleton) Arttir() int {
	s.Lock()
	defer s.Unlock()
	s.sayac++
	return s.sayac
}
