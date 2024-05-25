package plusone

import "fmt"

func plusOne(digits []int) []int {
	j := len(digits) - 1
	reminder := 1
	for j >= 0 && reminder >= 0 {
		tempsum := reminder + digits[j]
		if tempsum > 9 {
			tempsum = tempsum % 10
			reminder = 1
		} else {
			reminder = -1
		}
		digits[j] = tempsum
		j--
	}
	if reminder > 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

func TestplusOne() {
	nums := []int{1, 2, 3}

	fmt.Println(plusOne(nums))
}
