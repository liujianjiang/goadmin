package config

type AntsPool struct {
	//协程池容量大小
	Size int `yaml:"size"`
	//协程池最大可等待数量
	MaxBlockingTask int `yaml:"max_blocking_task"`
}
