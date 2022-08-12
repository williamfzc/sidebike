package server

import cmap "github.com/orcaman/concurrent-map/v2"

type Store[T any] struct {
	cmap.ConcurrentMap[*T]
}

func createStore[T any]() Store[T] {
	return Store[T]{cmap.New[*T]()}
}

var machineStore Store[Machine]
var taskStore Store[Task]

func GetMachineStore() *Store[Machine] {
	return &machineStore
}

func GetTaskStore() *Store[Task] {
	return &taskStore
}

func init() {
	machineStore = createStore[Machine]()
	taskStore = createStore[Task]()
}
