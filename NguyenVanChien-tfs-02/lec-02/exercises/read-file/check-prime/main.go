package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("testPrime.txt")
	if err != nil {
		log.Fatal(err)
		fmt.Println(string(data))
	}
	a := strings.Fields(string(data))
	fmt.Println(Prime(a))
}

func FindPrime(number int) bool {
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
func Prime(a []string) bool {
	number := 0
	for i := 0; i < len(a); i++ {
		number, _ = strconv.Atoi(a[i])
		if number == 2 {
			return true
		}
		if number%2 != 0 {
			if FindPrime(number) == true {
				return true
			}
		}
	}
	return false
}
