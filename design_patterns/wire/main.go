package main

import "fmt"

type WeightVC struct {
	MaxWeight int64 `json:"maxweight" default:"20000"`
}

func main() {
	var w WeightVC
	fmt.Println(w.MaxWeight)
}
