package main

import "fmt"
import "time"
import "strconv"

func isPalindrome(v int) bool {
	str := strconv.Itoa(v)
	right := len(str) - 1
	left := 0
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func maxPalindrome(maxfact int) int {

	max := -1
	minlow := 0

	// establish a low and high value pair as the values that create the product.
	for low := maxfact; low >= minlow; low-- {
		for high := maxfact; high >= low; high-- {
			prod := low * high
			if prod <= max {
				// no need to keep looking for things that can't possibly be larger than previous max.
				break
			}
			if isPalindrome(prod) {
				max = prod
				// limit how far we search back to things that would exceed the smallest factor
				minlow = max / maxfact
			}
		}
	}

	return max
}

func main() {
	start := time.Now()

	maxp := maxPalindrome(999)

	elapsed := time.Since(start)
	fmt.Println(maxp)
	fmt.Println("Programme execution time %s", elapsed)
}
