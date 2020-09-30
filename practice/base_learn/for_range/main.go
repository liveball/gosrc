package main

func main() {
	b := []int{1, 2, 3}

	bb := make([]*int, 0, 1)
	for k, v := range b {
		_ = &v    //8 line
		_ = &b[k] //9 line
		//bb = append(bb, &v)
		bb = append(bb, &b[k])
	}

	//for _, v := range bb {
	//	fmt.Println(*v)
	//}
}
