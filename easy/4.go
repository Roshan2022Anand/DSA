package easy

import (
	"fmt"
)

// thinking -999 | 9ms Beats78.23%
func firstUniqChar(s string) int {

	for i := 0; i < len(s); i++ {
		exist := false
		for j := 0; j < len(s); j++ {
			if i != j {
				if s[i] == s[j] {
					exist = true
					break
				}
			}
		}
		if !exist {
			return i
		}
	}

	return -1
}

func Run1stUniqChar() {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))
	fmt.Println(firstUniqChar("aabb"))
}

// Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.

// Example 1:
// Input: s = "leetcode"
// Output: 0

// Explanation:
// The character 'l' at index 0 is the first character that does not occur at any other index.

// Example 2:
// Input: s = "loveleetcode"
// Output: 2

// Example 3:
// Input: s = "aabb"
// Output: -1
