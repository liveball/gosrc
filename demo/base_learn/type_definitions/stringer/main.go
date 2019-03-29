package main

import (
	"fmt"
)

type Stringer interface {
	String()
	test()
}

// type Celsius float64

// func (c Celsius) String() string {
// 	return strconv.FormatFloat(float64(c), 'f', 1, 64) + "C"
// }

type Day int

var dayName = []string{"星期一", "星期二", "星期三", "星期四", "星期五", "星期六", "星期日"}

func (day Day) String() string {
	return dayName[day]
}
func (day Day) test() string {
	return dayName[4]
}

func main() {
	fmt.Println(Day(0))
	// fmt.Println(Day(1), "温度是:", Celsius(20.6))
}
