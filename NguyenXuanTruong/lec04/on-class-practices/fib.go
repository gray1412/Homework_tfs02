package main

import (
	"fmt"
)

func fibo(n int, c chan int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		c <- b
		a, b = b, a+b
	}
	close(c) // Khi close(c) sẽ thực hiện dừng vòng lặp for .. range
}

func main() {
	c := make(chan int)
	go fibo(7, c)
	for i := range c {
		fmt.Println(i)
	}
}