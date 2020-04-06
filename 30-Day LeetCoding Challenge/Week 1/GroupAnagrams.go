/*
Given an array of strings, group anagrams together.

Example:
Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
Output:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]

Note:
All inputs will be in lowercase.
The order of your output does not matter.
*/

package week1

/*
Tried to solve without implement sort Interface.
This is not the best solution :) (See alternative for better performance.)
*/
func groupAnagrams(strs []string) [][]string {
	result := [][]string{}
	for _, s := range strs {
		found := false
		for i, sArr := range result {
			if areAnagrams(s, sArr[0]) {
				found = true
				result[i] = append(result[i], s)
				break
			}
		}
		if !found {
			result = append(result, []string{s})
		}
	}
	return result
}

func areAnagrams(x, y string) bool {
	if len(x) != len(y) {
		return false
	}
	runesX := getRunesMap(x)
	runesY := getRunesMap(y)
	for _, r := range x {
		if runesX[r] != runesY[r] {
			return false
		}
	}

	for _, r := range y {
		if runesX[r] != runesY[r] {
			return false
		}
	}

	return true
}

func getRunesMap(x string) map[rune]int {
	runesX := map[rune]int{}
	for _, runeX := range x {
		_, ok := runesX[runeX]
		if ok {
			runesX[runeX]++
		} else {
			runesX[runeX] = 1
		}
	}
	return runesX
}

/*
Alternative Solution:

type bytes []byte

func (b bytes) Len() int {
    return len(b)
}

func (b bytes) Less(i, j int) bool {
    return b[i] < b[j]
}

func (b bytes) Swap(i, j int) {
    b[i], b[j] = b[j], b[i]
}

func groupAnagrams(strs []string) [][]string {
    var ret [][]string
    hm := make(map[string][]string)
    for _, str := range strs {
        key := bytes(str)
        sort.Sort(key)
        str1 := string(key)
        hm[str1] = append(hm[str1], str)
    }
    for _, words := range hm {
        ret = append(ret, words)
    }
    return ret
}
*/
