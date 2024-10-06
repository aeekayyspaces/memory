// Package base - This is the base package for shared primitives.
package base

import (
	"sync"
)

type SingletonI interface{
	GetInstance() SingletonI
}

type Singleton struct {
	Name	string		// The name of the singleton
	Data	map[string]interface{}	// A map of data within the object
}

var instance *Singleton

// sync.Once to ensure that instance is initialized only once
var once sync.Once

// GetInstance returns the singleton instance
func (s *Singleton) GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func (s *Singleton) GetData() map[string]interface{} {
	return s.GetInstance().Data
}
