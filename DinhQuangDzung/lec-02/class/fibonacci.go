package main

import (
	"fmt"
)

func fib(n int) int {
	if n <= 2 {
		return 1
	}
	return fib(n-2) + fib(n-1)
}

func fibSequence(n int) []int {
	var seq []int
	for i := 0; i < 10; i++ {
		num := fib(i)
		seq = append(seq, num)
	}
	return seq
}

func main() {
	sequence := fibSequence(10)
	fmt.Println("Fib sequence:", sequence)
}
