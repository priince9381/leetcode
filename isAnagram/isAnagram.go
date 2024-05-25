package isanagram

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mS := [26]rune{}
	for _, value := range s {
		mS[value-'a'] += 1
	}
	for _, value := range t {
		mS[value-'a'] -= 1
		if mS[value-'a'] < 0 {
			return false
		}
	}
	return true
}

func TestisAnagram() {
	s := "anagramt"
	t := "nagaramd"
	fmt.Println(isAnagram(s, t))
}
