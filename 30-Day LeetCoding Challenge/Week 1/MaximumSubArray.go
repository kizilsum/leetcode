/*
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:
Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6

Explanation: [4,-1,2,1] has the largest sum = 6.
Follow up:

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
*/

package week1

func maxSubArray(nums []int) int {
	ans := -2147483648
	a := 0
	for i := 0; i < len(nums); i++ {
		a += nums[i]
		ans = max(ans, a)
		a = max(a, 0)
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}