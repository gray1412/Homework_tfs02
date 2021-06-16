package main

import (
	"fmt"
	"sync"
	"time"
)

//su dung go rountine tinh va in toan bo tu 1 den n
//cac buoc:
//b1: code fibo khong co go rountine
//b2: bien code thanh go routine
//b3: them 1 go rountine de in ra
//b4: su dung channel de giao tiep

func CalculateWithoutRecursiveMethod(n int) int64 {
	// If n <= 1, just returns immediately for fast calculation
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	}
	// calculate nth number from beginning of fibonacci number string by using Formula:
	// Fn = Fn-1 + Fn-2
	nth := int64(0)
	first, second := int64(0), int64(1)
	for i := 2; i <= n; i++ {
		nth = first + second
		first, second = second, nth
	}
	return nth
}

//them vao channel ch
func fibo(ch chan int64, n int, wg *sync.WaitGroup, quit chan bool) {
	for i := 0; i < n; i++ {
		ch <- int64(CalculateWithoutRecursiveMethod(i))
	}
}

func print(ch chan int64) {
	for v := range ch {
		fmt.Println("gia tri thu la: ", v)
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan int64)
	wg := &sync.WaitGroup{}
	quit := make(chan bool)
	go fibo(ch, 9, wg, quit)
	go print(ch)
}
