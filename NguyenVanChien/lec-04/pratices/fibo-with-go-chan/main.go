package main

import (
	"fmt"
	"time"
)

/*
 In mathematics, the Fibonacci numbers, commonly denoted Fn, form a sequence,
 called the Fibonacci sequence, such that each number is the sum of the two preceding ones,
 starting from 0 and 1.
 Fibonacci Formula: Fn = Fn-1 + Fn-2
 Eg: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, ……
 With F(0) = 0 and F(1) = 1 Or F(0) = 1 and F(1) = 1
*/

func main() {
	ch := make(chan int64)
	go CalculateWithoutRecursiveMethod(10, ch)
	go print(ch)
	time.Sleep(time.Minute)
}

// CalculateWithRecursionMethod calculates fibonacci nth number by using recursion method
func CalculateWithRecursionMethod(n int64, ch chan int64) int64 {
	// If n <= 1, just returns 1
	switch n {
	case 0:
		ch <- 1
		return 1
	case 1:
		ch <- 1
		return 1
	}
	// recursively returns the same fibonacci function with n-1 and n-2 based on Fibonacci Formula
	ch <- CalculateWithRecursionMethod(n-1, ch) + CalculateWithRecursionMethod(n-2, ch)
	v := <-ch
	return v
}

// CalculateWithoutRecursiveMethod calculates fibonacci nth number without using recursion
// We must use a loop inside the function replace recursion method
func CalculateWithoutRecursiveMethod(n int64, ch chan int64) {
	// If n <= 1, just returns immediately for fast calculation
	switch n {
	case 0:
		ch <- 0
	case 1:
		ch <- 1
	}

	// calculate nth number from beginning of fibonacci number string by using Formula:
	// Fn = Fn-1 + Fn-2
	nth := int64(0)
	first, second := int64(0), int64(1)
	ch <- 0
	ch <- 1
	for i := int64(2); i <= n; i++ {
		nth = first + second
		ch <- nth
		first, second = second, nth
	}
}

func print(ch chan int64) {
	i := 0
	for n := range ch {
		fmt.Printf("Element %v is %v \n", i, n)
		i++
		time.Sleep(4 * time.Second)
	}
}
