package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("numbers.txt")
	if err != nil {
		log.Fatal(err)
		fmt.Println(string(data))
	}
	a := strings.Fields(string(data))
	fmt.Println(findMinMax(a))
	fmt.Println(finAverage(a))

}
func findMinMax(a []string) (min int, max int) {
	min, _ = strconv.Atoi(a[0])
	max, _ = strconv.Atoi(a[0])
	for i := 0; i < len(a); i++ {
		value, _ := strconv.Atoi(a[i])
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return
}
func finAverage(a []string) (medium float32) {
	sum := 0
	for i := 0; i < len(a); i++ {
		value, _ := strconv.Atoi(a[i])
		sum = sum + value
	}
	medium = float32(sum) / float32(len(a))
	return
}
