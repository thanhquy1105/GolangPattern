package singleton

/* Singleton interface */
type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton

// var once sync.Once

func init() {
	instance = &singleton{count: 100}
}

/*GetInstance function return object*/
func GetInstance() Singleton {
	// once.Do(func() {
	// 	time.Sleep(time.Second)
	// 	instance = &singleton{count: 100}
	// })
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
