package singleton

import (
	"sync"
)

type singletonGoIdiomatic struct {

}

var ins *singletonGoIdiomatic
var one sync.Once

func GetIns() *singletonGoIdiomatic {
	once.Do(func() {
		ins = &singletonGoIdiomatic{}
	})
	return ins
}


