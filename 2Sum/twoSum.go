package twoSum

import "fmt"

func twoSum(nums []int, target int) []int {
	valueIndexMap := make(map[int]int)
	for i, value := range nums {
		if _, ok := valueIndexMap[target - value]; ok {
			return []int{valueIndexMap[target - value], i}
		}
		valueIndexMap[value] = i
	}
	return []int{}
}

func TestTwoSum() {
	nums := []int{3, 3}
	target := 6
	fmt.Println(twoSum(nums, target))

}
