/*
Given a non-empty string check if it can be constructed by taking a substring of it
and appending multiple copies of the substring together.
You may assume the given string consists of lowercase English letters only and its length will not exceed 10000.

Example 1:
Input: "abab"
Output: True
Explanation: It's the substring "ab" twice.

Example 2:
Input: "aba"
Output: False

Example 3:
Input: "abcabcabcabc"
Output: True
Explanation: It's the substring "abc" four times. (And the substring "abcabc" twice.)
*/

package easy

func repeatedSubstringPattern(s string) bool {
	mid := (len(s) / 2) + 1
	runes := []rune(s)
	for sub := 1; sub < mid; sub++ {
		if isSubRepeated(runes[0:sub], runes) {
			return true
		}
	}
	return false
}

func isSubRepeated(sub []rune, runes []rune) bool {
	for i := 0; i < len(runes); {
		for j := 0; j < len(sub); j++ {
			if i >= len(runes) {
				return false
			}
			if runes[i] != sub[j] {
				return false
			}
			i++
		}
	}
	return true
}
