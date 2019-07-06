package main

import (
	"fmt"
	"sync"
)

// golang并发编程——安全传输引用和指针的方法

// 1、使用互斥量实现串行化访问
// 2、设定一个规则，一旦指针或者引用发送之后发送方就不会再访问它，然后让接收者来访问和释放指针或者引用
// 3、让所有导出的方法不能修改其值，所有修改其值的方法都不导出
type Stat struct{
	View int
}

//GetData test
func GetData(uid int) (res *Stat){
   var st *Stat
   st =new(Stat)
   st.View=uid
   res=st
   fmt.Printf("GetData uid=%d, res=%p\n", uid,res)
   return
}

func main(){
	var mystat *Stat

	var wg sync.WaitGroup

	for i := 1; i < 10; i++ {
		wg.Add(1)
		j:=i
		go func(){
			mystat = GetData(j)
	fmt.Printf("mystat=%+v,addr=%p\n", mystat,&mystat)
			wg.Done()
		}()
	}	

	wg.Wait()
}