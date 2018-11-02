package main

import (
	"fmt"

	"container/heap"
	"container/list"
	"container/ring"
)

func main() {

	heapTest()
	fmt.Println("----------------")
	listTest()
	fmt.Println("----------------")
	ringTest()
}

//heap提供了接口，需要自己实现如下方法
type Heap []int

//构造的是小顶堆，大顶堆只需要改一下下面的符号
func (h *Heap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Len() int {
	return len(*h)
}

func (h *Heap) Pop() interface{} {
	x := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return x
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Heap) Remove(idx int) interface{} {
	h.Swap(idx, h.Len()-1)
	return h.Pop()
}

func heapTest() {

	//创建一个heap
	h := &Heap{}

	heap.Init(h)
	//向heap中插入元素
	h.Push(5)
	h.Push(2)
	h.Push(1)
	h.Push(8)
	h.Push(4)
	h.Push(6)
	h.Push(2)

	//输出heap中的元素，相当于一个数组，原始数组
	fmt.Println(h)

	//这里必须要reheapify，建立好堆了
	heap.Init(h)

	//小顶堆对应的元素在数组中的位置
	fmt.Println(h)

	//移除下标为5的元素，下标从0开始
	h.Remove(5)

	//按照堆的形式输出
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	fmt.Println()
}

func listTest() {
	//创建一个双向链表
	ls := list.New()

	//向双向链表中插入26个小写字母
	for i := 97; i < 123; i++ {
		ls.PushFront(i) //PushFront()代表从头部插入，同样PushBack()代表从尾部插入
	}

	//遍历双向链表ls中的所有字母
	for it := ls.Front(); it != nil; it = it.Next() {
		fmt.Printf("%c ", it.Value)
	}
	fmt.Println()
}

func ringTest() {
	//创建10个元素的闭环
	r := ring.New(10)

	//给闭环中的元素赋值
	for i := 1; i <= r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}

	//循环打印闭环中的元素值
	r.Do(
		func(p interface{}) {
			fmt.Printf("%v,", p.(int)) //1,2,3,4,5,6,7,8,9,10
		})
	println()
	//获得当前元素之后的第5个元素
	r5 := r.Move(5)

	fmt.Println("r5.Value:", r5.Value) //6
	fmt.Println("r.Value:", r.Value)

	//链接当前元素r与r5，相当于删除了r与r5之间的元素
	r1 := r.Link(r5)
	fmt.Println("r1.Value:", r1.Value)

	r.Do(
		func(r interface{}) {
			fmt.Printf("%v,", r.(int)) //1,6,7,8,9,10
		})
	println()
}
