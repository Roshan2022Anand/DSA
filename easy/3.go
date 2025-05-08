package easy

import (
	"fmt"
)

// fastet one ever solved | 11ms Beats91.77%
func containsDuplicate(nums []int) bool {
	dupNum := make(map[int]int8, len(nums))

	for i := 0; i < len(nums); i++ {
		if _, ok := dupNum[nums[i]]; ok {
			return true
		}
		dupNum[nums[i]] = 1
	}

	return false
}

func RunContainDuplicate() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}

/**
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

Example 1:
Input: nums = [1,2,3,1]
Output: true

Example 2:
Input: nums = [1,2,3,4]
Output: false

*/
