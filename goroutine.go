package main

import (
	//"fmt"
	"fmt"
)

func longestPalindrome(s string) string {
	var max int
	var maxSub string
	s2 := []rune(s)
	for i := 0; i < len(s2); i++ {
		for j := i+1; j <= len(s2); j++ {
			sub := s2[i:j]
			if isPalindrome(sub) && len(sub) > max {
				max = len(sub)
				maxSub = string(sub)
			}
		}
	}

	return maxSub
}

func isPalindrome(s []rune) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}

func main() {
	var s string = "hello"
	fmt.Println([]byte(s))
	//var err error = syscall.Errno(1)
	//fmt.Println(err.Error())
	//fmt.Println(err)
	//var s string = "cbb"
	//s2 := []rune(s)
	//fmt.Println(isPalindrome(s2))
	//fmt.Println(longestPalindrome(s))

	//test.add()
	// go running()

	// var input string
	// fmt.Scanln(&input)
}
