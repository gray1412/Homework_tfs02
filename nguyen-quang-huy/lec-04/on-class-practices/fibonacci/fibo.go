package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		c <- a
		a, b = b, a+b
	}
	close(c)
}

func main(){
	n :=10
	c := make(chan int)
	go fibonacci(n, c)
	for i := range c {
		fmt.Println(i)
	} 
}


