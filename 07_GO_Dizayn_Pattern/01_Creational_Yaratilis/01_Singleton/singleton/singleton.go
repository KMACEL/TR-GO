package singleton

type singleton struct {
	sayac int
}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}
func (s *singleton) Arttir() int {
	s.sayac++
	return s.sayac
}
