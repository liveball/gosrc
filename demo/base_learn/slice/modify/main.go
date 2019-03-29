package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// noMake()
	// makeWithCap()
	// shareMake()

	// mycopy()
	modify()
	// del()
}

func noMake() {
	var sa []string
	fmt.Printf("addr:%p \t\tlen:%v content:%v\n", sa, len(sa), sa)
	for i := 0; i < 10; i++ {
		sa = append(sa, fmt.Sprintf("%v", i))
		fmt.Printf("addr:%p \t\tlen:%v content:%v\n", sa, len(sa), sa)
	}
	fmt.Printf("addr:%p \t\tlen:%v content:%v\n", sa, len(sa), sa)
}

func makeWithCap() {
	var sa = make([]string, 0, 10)
	fmt.Printf("addr:%p \t\tlen:%v content:%v\n", sa, len(sa), sa)
	for i := 0; i < 10; i++ {
		sa = append(sa, fmt.Sprintf("%v", i))
		fmt.Printf("addr:%p \t\tlen:%v content:%v\n", sa, len(sa), sa)
	}
	fmt.Printf("addr:%p \t\tlen:%v content:%v\n", sa, len(sa), sa)
}

func shareMake() {
	var osa = make([]string, 0)
	sa := &osa
	for i := 0; i < 10; i++ {
		*sa = append(*sa, fmt.Sprintf("%v", i))
		fmt.Printf("addr of osa:%p,\taddr:%p \t content:%v\n", osa, sa, sa)
	}
	fmt.Printf("addr of osa:%p,\taddr:%p \t content:%v\n", osa, sa, sa)
}

func mycopy() {
	var sa = make([]string, 0)
	for i := 0; i < 10; i++ {
		sa = append(sa, fmt.Sprintf("%v", i))
	}

	var da = make([]string, 0, 10)
	var cc = 0
	cc = copy(da, sa)
	fmt.Printf("copy to da(len=%d)\t%v\n", len(da), da)

	da = make([]string, 5)
	cc = copy(da, sa)
	fmt.Printf("copy to da(len=%d)\tcopied=%d\t%v\n", len(da), cc, da)

	da = make([]string, 10)
	cc = copy(da, sa)
	fmt.Printf("copy to da(len=%d)\tcopied=%d\t%v\n", len(da), cc, da)

}

func modify() {
	var ss []string
	fmt.Printf("[ local print ]\t:\t length:%v\taddr:%p\tisnil:%v\n", len(ss), ss, ss == nil)
	myprint("func print", ss)

	//切片尾部追加元素append elemnt
	for i := 0; i < 10; i++ {
		ss = append(ss, fmt.Sprintf("s%d", i))
	}
	fmt.Printf("[ local print ]\t:\tlength:%v\taddr:%p\tisnil:%v\n", len(ss), ss, ss == nil)
	myprint("after append", ss)

	//删除切片元素remove element at index
	index := 5
	ss = append(ss[:index], ss[index+1:]...)
	myprint("after delete", ss)

	//在切片中间插入元素insert element at index;
	//注意：保存后部剩余元素，必须新建一个临时切片
	rear := append([]string{}, ss[index:]...)
	ss = append(ss[0:index], "inserted")
	ss = append(ss, rear...)
	myprint("after insert", ss)
}

func myprint(msg string, ss []string) {
	fmt.Printf("[ %20s ]\t:\tlength:%v\taddr:%p\tisnil:%v\tcontent:%v", msg, len(ss), ss, ss == nil, ss)
	fmt.Println()
}

func del() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Printf("remove before a %#v \n", (*reflect.SliceHeader)(unsafe.Pointer(&a)))

	removeByRef(3, a)
	// removeByPointer(3, &a)
}

func removeByRef(m int, s []int) {
	tt := make([]int, 10, 10)
	for i := 0; i < len(s); i++ {
		if m == s[i] {

			fmt.Println("1", s[:i])
			fmt.Println("2", s[i+1:])

			t := append(s[:i], s[i+1:]...) //依赖同一个底层数组

			//copy t
			count := copy(tt, t)
			fmt.Printf("tt val(%+v) tt ref(%p) count(%d)\n", tt, tt, count)

			// r := (*reflect.SliceHeader)(unsafe.Pointer(&s))
			// r.Data = uintptr((*reflect.SliceHeader)(unsafe.Pointer(&t)).Data)

			fmt.Printf("s val(%+v) element(%d) ref(%p) addr(%p) slice %#v \n",
				s, s[3], s, &s[3],
				(*reflect.SliceHeader)(unsafe.Pointer(&s)),
			)

			s = t //s 和 t同为slice，引用的底层数组的地址一样，所以此赋值操作无效

			fmt.Printf("t val(%+v) element(%d) ref(%p) addr(%p) slice %#v \n",
				t, t[3], t, &t[3],
				(*reflect.SliceHeader)(unsafe.Pointer(&t)),
			)

		}
	}
}

func removeByPointer(m int, s *[]int) {
	for i := 0; i < len(*s); i++ {
		if m == (*s)[i] {
			t := append((*s)[:i], (*s)[i+1:]...) //依赖同一个底层数组

			*s = t
			fmt.Printf("*s (%p) \n", *s)

			fmt.Printf("t val(%v) ref(%p) addr(%p)  slice %#v \n",
				t, t, &t,
				(*reflect.SliceHeader)(unsafe.Pointer(&t)),
			)
		}
	}
}
