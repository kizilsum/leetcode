namespace CSharpSolutions.Easy
{
    /// <summary>
    /// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
    /// You may assume that each input would have exactly one solution, and you may not use the same element twice.
    ///
    /// Example:
    /// Given nums = [2, 7, 11, 15], target = 9,
    /// Because nums[0] + nums[1] = 2 + 7 = 9,
    /// return [0, 1].
    /// </summary>
    public partial class Solution
    {
        public static int[] TwoSum(int[] nums, int target)
        {
            for (int index1 = 0; index1 < nums.Length - 1; index1++)
            {
                for (int index2 = index1 + 1; index2 < nums.Length; index2++)
                {
                    if (nums[index1] + nums[index2] == target)
                    {
                        return new[] {index1, index2};
                    }
                }
            }

            return new[] {-1, -1};
        }
    }
}