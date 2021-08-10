package main

import "fmt"

func main()  {
	res:=letterCombinations("23")

	fmt.Println(res)
}

func letterCombinations(digits string) []string {
	return letterCombinations1(digits)
}

var (
	letterMap = []string{
		"",     //0
		"",     //1
		"abc",  //2
		"def",  //3
		"ghi",  //4
		"jkl",  //5
		"mno",  //6
		"pqrs", //7
		"tuv",  //8
		"wxyz", //9
	}
	res = []string{}
)

func letterCombinations1(digits string) []string {
	if digits == "" {
		return []string{}
	}

	res = []string{}
	findCombination(&digits, 0, "")
	return res
}

func findCombination(digits *string, index int, s string) {
	if index == len(*digits) {
		res = append(res, s)
		return
	}

	num := (*digits)[index]
	letter := letterMap[num-'0']

	fmt.Println(111, index, num, num-'0', letter)
	for i := 0; i < len(letter); i++ {
		findCombination(digits, index+1, s+string(letter[i]))
	}

	return
}
