package containdublicate

import "fmt"

func containsDuplicate(nums []int) bool {
	duplicateMap := make(map[int]bool)
	for _, value := range nums {
		if _, ok := duplicateMap[value]; ok {
			return true
		}
		duplicateMap[value] = true
	}
	return false
}

func TestContainDuplicate() {
	var nums = []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums))
}
