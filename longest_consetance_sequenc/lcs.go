package longestconsetancesequenc

import (
	"fmt"
	"slices"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slices.Sort(nums)
	maxSubSequenc := 0
	currentSubSeq := 0
	for i := 0; i < len(nums)-1; i++ {
		fmt.Println(nums[i+1], nums[i])
		if nums[i+1] == nums[i] {
			continue
		}
		if nums[i+1]-nums[i] == 1 || nums[i+1]-nums[i] == 0 {
			currentSubSeq += 1
		} else {
			currentSubSeq = 0
		}
		if currentSubSeq > maxSubSequenc {
			maxSubSequenc = currentSubSeq
		}

	}
	if maxSubSequenc == 0 {
		return 0
	}
	return maxSubSequenc + 1

}

func Testlongestconsetancesequenc() {

	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	output := longestConsecutive(nums)
	fmt.Println(output)
}
