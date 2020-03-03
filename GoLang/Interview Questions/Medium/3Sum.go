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

package medium

// TODO: Time Limit Exceeded! This is O(n^3), try better! Sort first and then calculate! Redo this problem.
func threeSum(nums []int) [][]int {
	length := len(nums)
	result := make([][]int, 0)
	for i := 0; i < length-2; i++ {
		for j := i + 1; j < length-1; j++ {
			for k := j + 1; k < length; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					result = appendToResult(result, nums[i], nums[j], nums[k])
				}
			}
		}
	}

	return result
}

func appendToResult(result [][]int, i int, j int, k int) [][]int {
	var arr []int
	var length = len(result)
	if i <= j && j <= k {
		arr = []int{i, j, k}
	} else if i <= k && k <= j {
		arr = []int{i, k, j}
	} else if j <= i && i <= k {
		arr = []int{j, i, k}
	} else if j <= k && k <= i {
		arr = []int{j, k, i}
	} else if k <= j && j <= i {
		arr = []int{k, j, i}
	} else {
		arr = []int{k, i, j}
	}

	for i := 0; i < length; i++ {
		if result[i][0] == arr[0] && result[i][1] == arr[1] && result[i][2] == arr[2] {
			return result
		}
	}

	return append(result, arr)
}
