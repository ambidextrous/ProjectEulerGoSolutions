// Project Euler: problem 3 solution
// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?

// Using a concurrent prime sieve adapted from code given in
// Rob Pike's 'Go Concurrency Patterns' talk:
// https://talks.golang.org/2012/concurrency.slide#1

package main

import "fmt"
import "time"

// Send the sequence 2,3,4, ... to channel 'ch'.
func Generate(ch chan<- int) {

	for i := 2; ; i++ {

		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {

	for {

		i := <-in // Recieve value from 'in',

		if i%prime != 0 {

			out <- i // Send 'i' to 'out'.
		}
	}
}

// The prime sieve: Daisy-chain Filter processes
func main() {

	start := time.Now()

	target := 600851475143 //957

	ch := make(chan int) // Create a new channel.
	go Generate(ch)      // Launch Generate goroutine.

	for i := 0; ; i++ {

		prime := <-ch

		if target%prime == 0 {

			oldTarget := target

			target = target / prime

			if target == 1 {

				target = oldTarget
				break
			}

		}

		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}

	elapsed := time.Since(start)
	fmt.Println("Highest prime factor = ", target)
	fmt.Printf("Programme execution tool %s", elapsed)
	fmt.Println()
}
