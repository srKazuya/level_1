package main

import (
	"fmt"
	"sync"
)

type MyMap struct {
	mx    sync.Mutex
	store map[string]int
}

func NewMyMap() *MyMap {
	return &MyMap{
		store: make(map[string]int),
	}
}

func (m *MyMap) Get(key string) (int, bool) {
	m.mx.Lock()
	defer m.mx.Unlock()
	val, ok := m.store[key]
	return val, ok
}

func (m *MyMap) Set(key string, val int) {
	m.mx.Lock()
	defer m.mx.Unlock()
	m.store[key] = val
}

func main() {
	mappy := NewMyMap()
	mappy.Set("Age", 81)
	v, ok := mappy.Get("Age")
	if !ok {
		fmt.Println("no val")
	}
	fmt.Println(v)
}
