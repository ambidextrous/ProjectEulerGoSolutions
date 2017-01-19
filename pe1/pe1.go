/*
Project Euler

Multiples of 3 and 5
Problem 1

If we list all the natural numbers below 10 that are multiples of
3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/

package main

import (
    "fmt"
    "time"
)

func sumMults(a, b, lim int) (sum int) {

    m := make(map[int]int)

    totA := 0

    for i := a; i < lim; i += a {

        totA += i
        m[i] = i
    }

    totB := 0

    for j := 0; j < lim; j += b {
        
        if _, ok := m[j]; ok == false {

            totB += j
        }
    }

    return totA + totB
}

func main() {
    start := time.Now()
    a, b := 3, 5
    lim := 1000000
    sum := sumMults(a,b,lim)
    elapsed := time.Since(start)
    fmt.Printf("The sum of all of the multiples of %v and %v below %v is %v", a, b, lim, sum)
    fmt.Println()
    fmt.Println(elapsed)
}
