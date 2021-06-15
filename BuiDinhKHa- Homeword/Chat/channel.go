package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
  In this application, we will use go-routine to demonstrate the process of crawling webpages concurrently
*/
func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {
	ch := make(chan int32)
	quit := make(chan bool)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go generateNumberPair(wg, ch, quit)
	go sum(wg, ch, quit)
	time.Sleep(5 * time.Second)
	quit <- true
	quit <- true
	wg.Wait()
	fmt.Println("Exit")
}
func generateNumberPair(wg *sync.WaitGroup, ch chan int32, quit chan bool) {
	defer wg.Done()
	for {
		select {
		case <-quit:
			fmt.Println("Quit generator")
			return
		default:
			n := rand.Int31()
			ch <- n
			// how can it be better?
			time.Sleep(time.Second)
		}
	}
}
func sum(wg *sync.WaitGroup, ch chan int32, quit chan bool) {
	defer wg.Done()
	sum := int64(0)
	for {
		select {
		case <-quit:
			fmt.Println("Quit sum")
			return
		case n := <-ch:
			sum += int64(n)
			fmt.Printf("Current sum is %v, n is: %v\n", sum, n)
		}
	}
}
