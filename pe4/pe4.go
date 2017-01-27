//Project Euler: Problem 4

//Largest palindrome product

//A palindromic number reads the same both ways. The largest
//palindrome made from the product of two 2-digit numbers is
//9009 = 91 × 99.

//Find the largest palindrome made from the product of two
//3-digit numbers.

package main

import "fmt"
import "time"

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	if s[0] != s[len(s)-1] {
		return false
	}
	s = s[:len(s)-1]
	s = s[1:]
	return isPalindrome(s)
}

func main() {

	start := time.Now()

	highestPal := 0

	lowerLimit := 100

	upperLimit := 1000

	for i := upperLimit - 1; i >= lowerLimit; i-- {

		for j := i - 1; j > 0; j-- {

			candidate := i * j

			s := fmt.Sprintf("%d", candidate)

			if isPalindrome(s) {

				if candidate > highestPal {

					highestPal = candidate
				}
			}
		}
	}

	elapsed := time.Since(start)

	//num := 1000 * 999
	//s := fmt.Sprintf("%d", num)
	//output := isPalindrome(s)
	//fmt.Println(output)
	//output = isPalindrome("101")
	//fmt.Println(output)
	//output = isPalindrome("9")
	fmt.Println(highestPal)
	fmt.Println("Programme execution time %s", elapsed)
}
