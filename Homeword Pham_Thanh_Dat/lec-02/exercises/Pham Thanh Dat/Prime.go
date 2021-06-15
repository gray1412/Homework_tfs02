package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
	"math"
)

func FindPrime(number int) bool{
	for i:=2; i<= int(math.Sqrt(float64(number))); i++ {
		if(number % i == 0){
			return false
			break
		}
	}
	return true
}
func Prime(arr [] string) bool{
	number := 0
	for i:=0; i<len(arr); i++ {
		number, _ = strconv.Atoi(arr[i])
		if(number == 2){
			return true
			break
		}
		if( number % 2 != 0) {
			if(FindPrime(number) == true){
				return true
				break
			}
		}
	}
	return false
}

func main() {
	content, err := ioutil.ReadFile("a.txt")
	if err != nil {
		log.Fatal(err)
		fmt.Println(string(content))
	}
	arr := strings.Fields(string(content))
	fmt.Println(Prime(arr))


}