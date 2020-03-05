/*
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:
The solution set must not contain duplicate triplets.

Example:
Given array nums = [-1, 0, 1, 2, -1, -4],
A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/

package medium_arrays

import "sort"

func threeSum(nums []int) [][]int {
	length := len(nums)

	if length < 3 {
		return nil
	}

	result := make([][]int, 0)
	sort.Ints(nums)

	for cursor, cursorVal := range nums[:length-2] {
		if cursor > 0 && nums[cursor] == nums[cursor-1] {
			continue
		}
		start := cursor + 1
		end := length - 1
		for start < end {
			sum := cursorVal + nums[start] + nums[end]
			if sum > 0 {
				end--
			} else if sum < 0 {
				start++
			} else {
				result = append(result, []int{cursorVal, nums[start], nums[end]})
				for start < end && nums[start] == nums[start+1] {
					start++
				}
				for start < end && nums[end] == nums[end-1] {
					end--
				}
				start++
				end--
			}
		}
	}
	return result
}
