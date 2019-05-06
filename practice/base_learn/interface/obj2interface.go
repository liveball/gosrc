package main

import "fmt"

func main() {
	NewSayService("").Hello()
}

type SayService interface {
	Hello()
}

type sayService struct {
	serviceName string
}

func NewSayService(serviceName string) SayService {
	if len(serviceName) == 0 {
		serviceName = "greeter"
	}
	return &sayService{
		serviceName: serviceName,
	}
}

func (c *sayService) Hello() {
	fmt.Println(c.serviceName)
}
