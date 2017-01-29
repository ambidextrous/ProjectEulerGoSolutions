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

// Improved thanks to advice from rolfl:
// http://codereview.stackexchange.com/questions/153787/parallel-brute-force-solution-for-project-euler-4-largest-palindrome-product

package main

import "fmt"
import "time"
import "strconv"
import "sync"

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

func getPals(i int, c chan int) {

	for j := i - 1; j > 0; j-- {

		candidate := i * j

		s := strconv.Itoa(candidate)

		if isPalindrome(s) {

			c <- candidate
		}
	}
}

var wg sync.WaitGroup

func main() {

	start := time.Now()
	//c := make(chan int)

	lowerLimit := 100
	upperLimit := 1000
	count := upperLimit - 1 - lowerLimit

	palindromes := make(chan int, 1024)

	var wg sync.WaitGroup
	wg.Add(count)

	for i := 0; i < count; i++ {
		go func(val int) {
			getPals(val, palindromes)
			wg.Done()
		}(i + lowerLimit)
	}

	go func() {
		// wait for all routines to complete
		wg.Wait()
		// close the channel to allow the terminal range
		close(palindromes)
	}()

	highestPal := 0
	for candidate := range palindromes {
		if candidate > highestPal {
			highestPal = candidate
		}
	}

	elapsed := time.Since(start)
	fmt.Println(highestPal)
	fmt.Println("Programme execution time %s", elapsed)
}
