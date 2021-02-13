// 使用懒惰模式的单例模式，使用双重检查加锁保证线程安全

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
