package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func reverse(s string) string {
	var reversed []byte
	for i := len(s) - 1; i >= 0; i-- {
		reversed = append(reversed, s[i])
	}
	return string(reversed)
}

func IsPalindrome(s string) bool {
	if utf8.RuneCountInString(s) <= 1 {
		return true
	}
	joined := strings.Join(strings.Split(s, " "), "")
	middle := len(joined) / 2
	half1 := joined[:middle]
	half2 := ""
	if len(joined)%2 == 0 {
		half2 = joined[middle:]
	} else {
		half2 = joined[middle+1:]
	}
	return half1 == reverse(half2)
}

func main() {
	fmt.Println(IsPalindrome("racecar"))
	fmt.Println(IsPalindrome("a man a plan a canal panama"))
}
