package main

func main() {
	b := []int{1}

	bb := make([]*int, 0, 1)
	for k, v := range b {
		_ = &v
		_ = &b[k]
		// bb = append(bb, &v)
		bb = append(bb, &b[k])
	}

	// for _, v := range bb {
	// 	fmt.Println(*v)
	// }
}
