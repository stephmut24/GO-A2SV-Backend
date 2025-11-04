package main

import (
	"fmt"
	"strings"
	"unicode"
)

func IsPalindrome(text string) bool {
	// Normalize input: lowercase and remove non-alphanumeric characters
	var cleaned strings.Builder
	for _, r := range strings.ToLower(text) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			cleaned.WriteRune(r)
		}
	}

	str := cleaned.String()
	n := len(str)

	// Two-pointer check
	for i := 0; i < n/2; i++ {
		if str[i] != str[n-1-i] {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama")) // true
	fmt.Println(IsPalindrome("racecar"))                        // true
	fmt.Println(IsPalindrome("Hello"))                          // false
}