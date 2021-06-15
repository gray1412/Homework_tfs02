package main

import "fmt"

func fibonacci(c chan int, quit chan bool) {
	prev2, prev1 := 0, 1
	for {
		select {
		case c <- prev2:
			prev2, prev1 = prev1, prev2+prev1
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan bool)

	go func() {
		for i := 0; i <= 10; i++ {
			fmt.Printf("%v ", <-c)
		}
		quit <- true
	}()

	fibonacci(c, quit)
}
