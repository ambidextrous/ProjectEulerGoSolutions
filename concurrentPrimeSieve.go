// A concurrent prime sieve
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

	ch := make(chan int) // Create a new channel.
	go Generate(ch)      // Launch Generate goroutine.

	for i := 0; i < 1000; i++ {

		prime := <-ch
		fmt.Println(prime)
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}

	elapsed := time.Since(start)
	fmt.Printf("Programme execution tool %s", elapsed)
	fmt.Println()
}
