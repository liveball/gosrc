package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "abcdabcdabcd"
	// println(Replace(s, "ab", "bb", 1)) //strings.Count(s, old)  可以统计到ab 在s中出现了3次

	println(ReplaceAll(s, "ab", "bb"))
}

// ReplaceAll returns a copy of the string s with all
// non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune string.
func ReplaceAll(s, old, new string) string {
	return Replace(s, old, new, -1)
}

// Replace returns a copy of the string s with the first n
// non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune string.
// If n < 0, there is no limit on the number of replacements.
func Replace(s, old, new string, n int) string {
	if old == new || n == 0 {
		return s // avoid allocation
	}

	// Compute number of replacements.
	if m := strings.Count(s, old); m == 0 { //计算子串在字串中出现的次数
		return s // avoid allocation  未出现直接返回
	} else if n < 0 || m < n { // -1 表示出现的所有次数，真正出现的次数小于传入的次数则使用真正出现的次数
		n = m
	}

	// Apply replacements to buffer.
	t := make([]byte, len(s)+n*(len(new)-len(old)))
	w := 0
	start := 0
	for i := 0; i < n; i++ { //执行n次替换
		j := start
		if len(old) == 0 {
			if i > 0 {
				_, wid := utf8.DecodeRuneInString(s[start:])
				j += wid
			}
		} else {
			j += strings.Index(s[start:], old) //计算每次被替换子串的起始索引
		}

		w += copy(t[w:], s[start:j]) //计算每次拷贝

		w += copy(t[w:], new) //计算新串被拷贝的次数

		// fmt.Println(t[w:], string(s[start:j]))
		start = j + len(old) //计算最后一个被替换之后剩余字符串起始索引
		fmt.Println("w=", w, "start=", start, "j=", j)
	}

	// fmt.Println(111, string(t), w, t[w:], s[start:])

	w += copy(t[w:], s[start:])

	// fmt.Println(222, string(t), w, t[w:], s[start:])
	return string(t[0:w])
}
