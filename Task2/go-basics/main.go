package main

import (
	"fmt"
	"unicode"
)
func isPalindrome(s string) bool {
	IsPalindrome := true
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			IsPalindrome = false
			break
		}
		left++
		right--
	}
	return IsPalindrome
}

func getFrequency(s string) map[rune]int {
	frequency := make(map[rune]int)
	for _, char := range s {
		if unicode.IsSpace(char) {
			continue
		}
		frequency[char]++
	}
	return frequency
}
func main() {
	
	// test the isPalindrome function
	palindromeTest := "racecar"
	if isPalindrome(palindromeTest) {
		fmt.Printf("%s is a palindrome\n", palindromeTest)
	} else {
		fmt.Printf("%s is not a palindrome\n", palindromeTest)
	}
	// test the getFrequency function
	frequencyTest := "hello world"
	frequency := getFrequency(frequencyTest)
	fmt.Printf("Character frequency in '%s':\n", frequencyTest)
	for char, count := range frequency {
		fmt.Printf("'%c': %d\n", char, count)	
	}
}