package main

import "fmt"

func main()  {
	slice:=[]int{11}

	//myMap:=make(map[int]*int)

	for _,v:=range slice {
		fmt.Println(&v)
	}

	//fmt.Println(myMap)
}
