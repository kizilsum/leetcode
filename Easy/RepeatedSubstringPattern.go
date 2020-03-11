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
	length := len(s)
loop:
	for i := 2; i < length/2; i++ {
		if length%i != 0 {
			continue
		}

		subLength := length / i

		for j := 0; j < subLength; j++ {
			for k := 0; k < i; k++ {
				if s[i*j+k] != s[k] {
					continue loop
				}
			}
		}
		return true
	}
	return false
}
