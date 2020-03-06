/*
Write an algorithm to determine if a number is "happy".
A happy number is a number defined by the following process:
Starting with any positive integer, replace the number by the sum of the squares of its digits,
and repeat the process until the number equals 1 (where it will stay),
or it loops endlessly in a cycle which does not include 1.
Those numbers for which this process ends in 1 are happy numbers.

Example:
Input: 19
Output: true

Explanation:
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1
*/

package medium

import (
	"math"
)

func isHappy(n int) bool {
	if n == 1 {
		return true
	}

	m := make(map[int]bool)

	for {
		n = getTotalForSquareOfDigits(n)

		if n == 1 {
			return true
		}

		found := m[n]

		if found {
			return false
		} else {
			m[n] = true
		}
	}
}

func getTotalForSquareOfDigits(n int) int {
	counter := 1
	total := 0
	for n != 0 {
		r := n % int(math.Pow(10, float64(counter)))
		n -= r
		total += int(math.Pow(float64(r/int(math.Pow(10, float64(counter-1)))), 2))
		counter++
	}
	return total
}
