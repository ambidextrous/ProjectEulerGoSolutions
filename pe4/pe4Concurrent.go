//Project Euler: Problem 4

//Largest palindrome product

//A palindromic number reads the same both ways. The largest
//palindrome made from the product of two 2-digit numbers is
//9009 = 91 × 99.

//Find the largest palindrome made from the product of two
//3-digit numbers.

// Note:
// Brute force solution launching 900 goroutines outperforms
// the same algorithms without goroutines by a factor of
// five: 14 ms vs. 72 ms.

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

func getHighestPal(i int, c chan int) {

	for j := i - 1; j > 0; j-- {

		candidate := i * j

		s := fmt.Sprintf("%d", candidate)

		if isPalindrome(s) {

			c <- candidate
		}
	}
}

func main() {

	start := time.Now()

	highestPal := 0

	lowerLimit := 100

	upperLimit := 1000

	c := make(chan int)

	for i := upperLimit - 1; i >= lowerLimit; i-- {

		go getHighestPal(i, c)
	}

	for i := 0; i < upperLimit-lowerLimit; i++ {

		candidate := <-c

		if candidate > highestPal {

			highestPal = candidate
		}
	}

	elapsed := time.Since(start)
	fmt.Println(highestPal)
	fmt.Println("Programme execution time %s", elapsed)
}
