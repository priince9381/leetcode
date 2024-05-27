package sumsorted

import "fmt"

func twoSum(nums []int, target int) []int {
	i := 0
	j := len(nums) - 1
	var tempSum int
	for i < j {
		tempSum = nums[i] + nums[j]
		if tempSum > target {
			j--
		} else if tempSum < target {
			i++

		} else if tempSum == target {
			return []int{i + 1, j + 1}
		}
	}
	return []int{}
}

func TestTwoSum() {
	inputArray := []int{2, 7, 11, 15}
	fmt.Println(twoSum(inputArray, 9))
}
