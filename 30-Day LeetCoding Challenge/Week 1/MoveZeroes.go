/*
Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Example:
Input: [0,1,0,3,12]
Output: [1,3,12,0,0]

Note:
	1. You must do this in-place without making a copy of the array.
	2. Minimize the total number of operations.
*/

package week1

func moveZeroes(nums []int)  {
	index := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if index == -1 {
				index = i
			}
			continue
		}
		if index != -1 {	
			nums[index] = nums[i]
			nums[i] = 0
			index++
		}
	}
}