package pivot

import "fmt"

func SumArray(nums []int) int {
	var totalSum int
	for _, num := range nums {
		totalSum += num
	}
	return totalSum
}

func pivotIndex(nums []int) int {
	rightSum := SumArray(nums) - nums[0]
	if rightSum == 0 {
		return 0
	}
	var sumleft int
	sumleft = nums[0]
	for i, value := range nums[1:] {
		if sumleft == rightSum-value {
			return i + 1
		}
		sumleft += value
		rightSum -= value
	}
	return -1
}

func PivotTest() {
	fmt.Println(pivotIndex([]int{2, 1, -1}))
}
