package main

import (
	"fmt"
)

// The thing about happy numbers is that you can easy define them recursively.
// - Base case: 1 is a happy number
// - Induction: A number is happy if the sum of the squares of its digits is
//   happy
// (The input numbers must be positive integers.)
// A number is unhappy if it loops forever instead of converging to 1.
// For example: 7 is happy if 7^2 = 49 is happy if 4^2 + 9^2 = 97 is happy if
// 9^2 + 7^2 = 130 is happy if 1^2 + 3^2 = 10 is happy if 1^2 + 0^2 = 1 is
// happy, which it is. Therefore, 7, 49, 97, 130 and 10 are all happy numbers.

// Because happiness is recursive, we can memoize it by tracking which numbers
// we have already marked as happy or unhappy. This is only really useful if
// we're checking if a long series of numbers is happy or not.
// However, it is also useful for seeing if there is a cycle. When first
// processing a number, its descendants won't have entries in the map. If we
// find that a descendent *is* in the map, we have encountered a cycle and the
// original number is unhappy.
// The base case (1 is happy) is included by default.
var happies = map[int]bool{ 1: true }


// sumDigitSquares simply takes an integer > 0 and computes the sum of the
// squares of its digits. This is accomplished by using %10 to get the lowest
// digit and /10 (integer division) to shift the number down.
func sumDigitSquares(i int) int {
	sum := 0
	for i != 0 {
		n := i % 10
		sum += n * n
		i = i / 10
	}
	return sum
}

// isHappy recursively determines the happiness of the input number.
// First it checks if the input is 1, which is automatically happy.
// Then it checks if the number has been computed before. If so, it just returns
// the previous computer happiness. If not, it determines the happiness of the
// number's descendent.
func isHappy(n int) bool {
	if n == 1 {
		return true
	}

	// A truly lovely Go idiom. Here, h is the happiness (true or false) and
	// e is the existence: true if n is already in the map or false
	// otherwise.
	h, e := happies[n]

	if !e {
		// setting n to false marks the number as entered in the map,
		// bypassing this block next time isHappy(n) is called.
		happies[n] = false
		// Then we go ahead and calculate the happiness of this number's
		// descendant
		happies[n] = isHappy(sumDigitSquares(n))
		return happies[n]
	}

	return h
}

func main() {
	happyCount := 0
	happyNums := make([]int, 8)
	for n := 1; happyCount < 8; n++ {
		if isHappy(n) {
			happyNums[happyCount] = n
			happyCount++
		}
	}
	fmt.Printf("The first %d happy numbers are:\n%d\n", happyCount, happyNums)
}
