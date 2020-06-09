package main

import "fmt"

type T struct {
	Value string
}

func main() {
	//tm := make(map[string]T)
	//
	//tm["cat"].Value = "this is a cat"
	//tm["cat"].Value = "this is anoher cat"
	//
	//fmt.Println(tm["cat"].Value)

	s2 := make(map[string]*int)
	n := 1
	s2["chenchao"] = &n
	fmt.Println(*s2["chenchao"])
}

//
//func main() {
//	tm := make(map[string]T)
//
//	tm["cat"] = T{ //copy 操作
//		Value: "this is a cat",
//	}
//	//tm["cat"].Value = "this is anoher cat" //修改操作 编译失败，
//
//	fmt.Println(tm["cat"].Value)
//
//	tm2 := make(map[string]*T)
//	tm2["cat01"] = &T{ //赋值操作
//		Value: "this is a cat",
//	}
//	tm2["cat01"].Value = "this is anoher cat" //修改操作
//	fmt.Println(tm2["cat01"].Value)
//
//	tm3 := make(map[string]int)
//	tm3["cat02"] = 2 //赋值操作
//	tm3["cat03"] = 3
//	fmt.Println(tm3)
//}

//上面的代码会编译失败，因为在go中 map中的赋值属于值copy，
//就是在赋值的时候是把T的完全复制了一份，复制给了map。而在go语言中，是不允许将其修改的。
//但是如果map的Value为int，是可以修改的，因为修改map中的int属于赋值的操作。
