package singleton

import (
	"sync"
)

type singleton struct {

}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	if instance == nil {
		instance = &singleton{}   // not safe
	}
	return instance
}


