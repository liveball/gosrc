package main

import (
	"bytes"
	"strings"
	"unicode/utf8"
)

func main(){
	a:="Test00002叮当猫7251"
	println("len",len(a))

	println("len rune",len([]rune(a)))
	println("bytes.Count", bytes.Count([]byte(a),nil) - 1)
	println("strings.Count",strings.Count(a,"")-1)
	println("utf8.RuneCountInString",utf8.RuneCountInString(a))


	if utf8.RuneCountInString(a)>16{//最大长度为16个字符，当超过16个时，展示为15个字符...
		runes := []rune(a)
		a=string(runes[0:15])+"..."
	}
	println(a)
	println(utf8.RuneCountInString(a))
}
