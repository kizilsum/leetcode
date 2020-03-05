/*
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.
A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

Example:
Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].

Note:
Although the above answer is in lexicographical order, your answer could be in any order you want.
*/

package medium_backtracking

func letterCombinations(digits string) []string {
	m := map[rune][]string{
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"}}
	result := make([]string, 0)
	for index, digit := range digits {
		temp := result
		if index == 0 {
			result = m[digit]
			continue
		}
		result = make([]string, 0)
		for i, tempVal := range temp {
			for j, mapVal := range m[digit] {
				result = append(result, tempVal+mapVal)
				_ = j
			}
			_ = i
		}
	}

	return result
}
