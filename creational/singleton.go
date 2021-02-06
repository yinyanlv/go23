package creational

import (
	"sync"
)

const count = 100

type Singleton struct {
}

var singleton *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})

	return singleton
}
