package easy

import "fmt"

// first try 0ms Beats 100.00%
func twoSum(nums []int, target int) []int {
	nC := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if j, ok := nC[nums[i]]; ok {
			return []int{j, i}
		}
		nC[target-nums[i]] = i
	}
	return []int{0, 0}
}

func RunTwoSum() {
	fmt.Println(" Easy : TWO SUM :")
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // [0,1]
	fmt.Println(twoSum([]int{3, 2, 4}, 6))      // [1,2]
	fmt.Println(twoSum([]int{3, 3}, 6))         // [0,1]
}
