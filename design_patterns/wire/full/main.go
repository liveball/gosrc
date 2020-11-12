package main

import (
	"gosrc/design_patterns/wire/full/injector"
	"gosrc/go/src/fmt"
	"gosrc/go/src/log"
)

func main()  {
	fn,err:=injector.UserLoader()
	if err!=nil{
		log.Fatal(err)
	}

	user:=fn(123)

	fmt.Println(user.Age)
}
