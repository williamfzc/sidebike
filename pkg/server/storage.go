package server

import (
	lru "github.com/hashicorp/golang-lru"
)

type Store[T any] struct {
	*lru.Cache
}

var machineStore *Store[Machine]
var taskStore *Store[Task]

func GetMachineStore() *Store[Machine] {
	return machineStore
}

func GetTaskStore() *Store[Task] {
	return taskStore
}

func (store *Store[T]) GetWithType(key interface{}) (*T, bool) {
	ret, ok := store.Get(key)
	return ret.(*T), ok
}

func (store *Store[T]) GetAll() []*T {
	var ret []*T
	for _, each := range store.Keys() {
		item, ok := store.GetWithType(each)
		if ok {
			ret = append(ret, item)
		}
	}
	return ret
}

func init() {
	ret, _ := lru.New(128)
	machineStore = &Store[Machine]{ret}

	ret, _ = lru.New(512)
	taskStore = &Store[Task]{ret}
}
