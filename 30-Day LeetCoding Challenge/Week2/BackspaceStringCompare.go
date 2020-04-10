/*
Given two strings S and T, return if they are equal when both are typed into empty text editors. # means a backspace character.

Example 1:
Input: S = "ab#c", T = "ad#c"
Output: true
Explanation: Both S and T become "ac".

Example 2:
Input: S = "ab##", T = "c#d#"
Output: true
Explanation: Both S and T become "".

Example 3:
Input: S = "a##c", T = "#a#c"
Output: true
Explanation: Both S and T become "c".

Example 4:
Input: S = "a#c", T = "b"
Output: false
Explanation: S becomes "c" while T becomes "b".

Note:
1 <= S.length <= 200
1 <= T.length <= 200
S and T only contain lowercase letters and '#' characters.

Follow up:
Can you solve it in O(N) time and O(1) space?
*/

package week2

func backspaceCompare(S string, T string) bool {
	return getString(S) == getString(T)
}

func getString(S string) string {
	runes := []rune{}
	counter := 0
	for i := len(S) - 1; i >= 0; i-- {
		if S[i] == byte('#') {
			counter++
		} else if counter > 0 {
			counter--
		} else {
			runes = append(runes, rune(S[i]))
		}
	}
	return string(runes)
}
