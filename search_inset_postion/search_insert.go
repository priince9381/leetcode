package searchinsetpostion

import "fmt"

func searchInsert(nums []int, target int) int {
	i := 0
	j := len(nums) - 1
	for i < j {
		if nums[i] == target {
			return i
		}
		if nums[j] == target {
			return j
		}
		if nums[i] > target {
			return i
		}
		if nums[j] < target {
			return j + 1
		}
		i++
		j--
	}
	if nums[i] < target {
		return i + 1
	}
	return i
}

func TestSearchInsert() {
	nums := []int{1}
	target := 1
	fmt.Println(searchInsert(nums, target))
}
