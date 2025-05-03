package main

import "fmt"

/**
VALID ANAGRAM

given two strings s and t, return true if t is an anagram of s, and false otherwise

example 1:
	input : s = "anagram", t = "nagaram"
	output : true

example 2:
	input ; s= "rat", t = "car"
	output : false
*/

// first try | 10ms Beats 9.50% | 5.63MB Beats 11.86%
func validAnagram1(s, t string) bool {

	if len(s) != len(t) {
		return false
	}

	var sCount = map[string]int{}
	var tCount = map[string]int{}

	for i := 0; i < len(s); i++ {
		sCount[string(s[i])] += 1
		tCount[string(t[i])] += 1
	}

	for str, c := range sCount {
		if tCount[str] != c {
			return false
		}
	}

	return true
}

// improvised | 8ms Beats 15.27% | 4.63MB Beats 99.27%
func validAnagram2(s, t string) bool {

	if len(s) != len(t) {
		return false
	}

	sCoutn := map[byte]int{}

	for i := 0; i < len(s); i++ {
		sCoutn[s[i]]++
	}

	frq := len(t)
	for i := 0; i < len(t); i++ {
		tStr := t[i]
		if sCoutn[tStr] > 0 {
			frq--
			sCoutn[tStr]--
		} else {
			return false
		}
	}

	return true
}

// some more improvised | 7ms Beats 23.44%
func validAnagram3(s, t string) bool {

	if len(s) != len(t) {
		return false
	}

	sCoutn := map[byte]int{}

	for i := 0; i < len(s); i++ {
		sCoutn[s[i]] += 1
	}

	for i := 0; i < len(t); i++ {
		if sCoutn[t[i]] == 0 {
			return false
		}
		sCoutn[t[i]]--
	}

	return true
}

// finally | 0ms Beats 100.00% | 4.78MB Beats 71.92%
func validAnagram4(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var sCount [26]int
	var tCount [26]int

	for i := 0; i < len(s); i++ {
		sCount[s[i]-'a']++
		tCount[t[i]-'a']++
	}
	return tCount == sCount //it's straight forward to compare array
}

func main() {
	fmt.Println(validAnagram1("anagram", "nagaram"))
	fmt.Println(validAnagram2("a", "ab"))
	fmt.Println(validAnagram3("ggii", "eekk"))

	fmt.Println(validAnagram4("anagram", "nagaram"))
	fmt.Println(validAnagram4("rat", "car"))
	fmt.Println(validAnagram4("a", "ab"))
	fmt.Println(validAnagram4("ggii", "eekk"))
}
