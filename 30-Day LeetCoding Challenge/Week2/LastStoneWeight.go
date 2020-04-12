/*
We have a collection of stones, each stone has a positive integer weight.
Each turn, we choose the two heaviest stones and smash them together.  Suppose the stones have weights x and y with x <= y.  The result of this smash is:

If x == y, both stones are totally destroyed;
If x != y, the stone of weight x is totally destroyed, and the stone of weight y has new weight y-x.
At the end, there is at most 1 stone left.  Return the weight of this stone (or 0 if there are no stones left.)

Example 1:
Input: [2,7,4,1,8,1]
Output: 1
Explanation:
We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
we combine 1 and 1 to get 0 so the array converts to [1] then that's the value of last stone.

Note:
	1. 1 <= stones.length <= 30
	2. 1 <= stones[i] <= 1000
*/

package week2

func lastStoneWeight(stones []int) int {
	s1, s2, stones := getHeaviestTwoStones(stones)
	r := s1 - s2
	for s1 != 0 && s2 != 0 {
		stones = append(stones, r)
		s1, s2, stones = getHeaviestTwoStones(stones)
		r = s1 - s2
	}
	return r
}

func getHeaviestTwoStones(stones []int) (int, int, []int) {
	s1, s2 := 0, 0
	i1, i2 := -1, -1
	for i, val := range stones {
		if val > s1 {
			s2 = s1
			s1 = val
			i2 = i1
			i1 = i
		} else if val > s2 {
			s2 = val
			i2 = i
		}
	}

	if i1 != -1 {
		stones[i1] = 0
	}
	if i2 != -1 {
		stones[i2] = 0
	}

	return s1, s2, stones
}
