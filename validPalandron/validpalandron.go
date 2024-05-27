package validpalandron

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	fmt.Println(i, j)
	for i <= j {
		if unicode.IsDigit(rune(s[i])) || unicode.IsLetter(rune(s[i])) {
			if unicode.IsDigit(rune(s[j])) || unicode.IsLetter(rune(s[j])) {
				if unicode.IsDigit(rune(s[i])) || unicode.IsDigit(rune(s[j])) {
					if s[i] != s[j] {
						return false
					}
				}
				if !strings.EqualFold(strings.ToLower(string(s[i])), strings.ToLower(string(s[j]))) {
					return false
				}
				j--
				i++
			} else {
				j--
			}
		} else {
			i++
		}
	}
	return true
}

func TestPalindron() {
	s := "0P"
	fmt.Println(isPalindrome(s))
}
