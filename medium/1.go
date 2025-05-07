package medium

import (
	"fmt"
)

// first try | 203ms Beats 5.03% very bad
func checkAnagram(comp1, comp2 string) bool {

	if len(comp1) != len(comp2) {
		return false
	}

	var c1Count [26]int
	var c2Count [26]int

	for i := 0; i < len(comp2); i++ {
		c1Count[comp1[i]-'a']++
		c2Count[comp2[i]-'a']++
	}

	return c1Count == c2Count
}
func groupAnagrams(strs []string) [][]string {

	result := [][]string{}

	for i := 0; i < len(strs); i++ {
		inserted := false
		for j := 0; j < len(result); j++ {
			if checkAnagram(result[j][0], strs[i]) {
				result[j] = append(result[j], strs[i])
				inserted = true
			}
		}

		if !inserted {
			result = append(result, []string{strs[i]})
		}
	}

	return result
}

// a little GOlang way | 8ms Beats 81.40%
func groupAnagrams2(strs []string) [][]string {
	tCount := map[string][]string{}
	result := [][]string{}

	for i := 0; i < len(strs); i++ {
		strCount := make([]byte, 26)

		for j := 0; j < len(strs[i]); j++ {
			strCount[strs[i][j]-'a']++
		}

		adrs := string(strCount)
		tCount[adrs] = append(tCount[adrs], strs[i])
	}

	for _, set := range tCount {
		result = append(result, set)
	}
	return result
}

func RunGrpAnagram() {
	fmt.Println("MEDIUM : GROUP ANAGRAM")

	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{""}))

	fmt.Println(groupAnagrams2([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams2([]string{""}))
}

/**
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

Example 1:
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Explanation:
There is no string in strs that can be rearranged to form "bat".
The strings "nat" and "tan" are anagrams as they can be rearranged to form each other.
The strings "ate", "eat", and "tea" are anagrams as they can be rearranged to form each other.

Example 2:
Input: strs = [""]
Output: [[""]]

Example 3:
Input: strs = ["a"]
Output: [["a"]]

*/
