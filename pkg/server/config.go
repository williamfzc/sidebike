package server

type Config struct {
	Port       int          `json:"port"`
	Debug      bool         `json:"debug"`
	InnerTasks []TaskDetail `json:"innerTasks"`
}
