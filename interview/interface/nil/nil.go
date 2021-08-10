package main


func main() {
	var a interface{} = nil
	var b *int = nil

	isNil(a)
	isNil(b)
}

func isNil(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}
