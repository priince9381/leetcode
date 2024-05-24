package lastworldlength

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	t := strings.TrimSpace(s)
	m := strings.Split(t, " ")
	return len(m[len(m)-1])
}

func TestLastWord() {
	s := "Hello t      World  "
	fmt.Println(lengthOfLastWord(s))
}
