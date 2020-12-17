package main

import (
	"fmt"
	"sync"
)

type config struct {
	once sync.Once
	path string
}

func NewInstance(path string) (c *config) {
	o := sync.Once{}

	o.Do(func() {
		c = &config{
			path: path,
		}
	})

	return
}

func main() {
	conf := NewInstance("/a/b/c")
	fmt.Println(conf)
}
