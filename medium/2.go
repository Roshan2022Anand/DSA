package medium

import (
	"fmt"
)

// had full confidence in first try | 279ms Beats 7.12%
func lengthOfLongestSubstring(s string) int {

	sExist := map[byte]int{}
	maxLen := 0
	sLen := 0

	for i := 0; i < len(s); i++ {

		//check if the letter is repeating
		if j, ok := sExist[s[i]]; ok {
			sExist = map[byte]int{} // empty the map

			// set the greatest lenght
			if sLen > maxLen {
				maxLen = sLen
			}

			//to avoid infinite loop
			i--
			if j != i {
				i = j
			}

			sLen = 0
		} else {
			sExist[s[i]] = i
			sLen++
		}
	}

	if sLen > maxLen {
		maxLen = sLen
	}

	return maxLen
}

func RunLenOstr() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	// fmt.Println(lengthOfLongestSubstring2("abcabcbb"))
	// fmt.Println(lengthOfLongestSubstring2("bbbbb"))
	// fmt.Println(lengthOfLongestSubstring2("pwwkew"))
	// fmt.Println(lengthOfLongestSubstring2(" "))
	// fmt.Println(lengthOfLongestSubstring2("dvdf"))
	// fmt.Println(lengthOfLongestSubstring2("aab"))
}

/**
Given a string s, find the length of the longest substring without duplicate characters.

Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

*/
