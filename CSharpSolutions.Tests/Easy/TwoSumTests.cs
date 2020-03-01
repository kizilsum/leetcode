using CSharpSolutions.Easy;
using Xunit;

namespace CSharpSolutions.Tests.Easy
{
    public class TwoSumTests
    {
        [Theory]
        [InlineData(0, 0, 1, 0, 0)]
        [InlineData(9, 0, 1, 2, 7, 11, 15)]
        [InlineData(4, 0, 2, 2, 7, 2, 15)]
        [InlineData(11, 1, 2, 1, 4, 7, 9, 14, 31, 1)]
        [InlineData(6, 4, 8, 3, 4, 7, 8, 0, 12, 42, 15, 6, 1)]
        public void ShouldReturnCorrectIndicesForGivenInput(int target, int index1, int index2, params int[] array)
        {
            int[] result = Solution.TwoSum(array, target);

            Assert.Equal(result[0], index1);
            Assert.Equal(result[1], index2);
        }
    }
}