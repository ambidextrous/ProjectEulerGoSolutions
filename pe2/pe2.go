package main

import "fmt"

// Programme to calculate the sum of the even numbers
// in the Fibonacci series below a given integer value.
func fibonacci() func() (int, bool) {

	reachedTotal := false
	sumEven := 0
	first := 0
	second := 1

	return func() (int, bool) {

		ret := first
		first, second = second, first+second

		if ret%2 == 0 {

			sumEven += ret
		}

		if ret > 4000000 {

			reachedTotal = true
		}

		return sumEven, reachedTotal
	}

}

func main() {

	f := fibonacci()

	prevNum := 0
	newNum := 0
	reachedTotal := false

	i := 0

	for {

		prevNum = newNum

		newNum, reachedTotal = f()

		if reachedTotal {

			fmt.Println(prevNum)

			break
		}

		i++
	}
}
