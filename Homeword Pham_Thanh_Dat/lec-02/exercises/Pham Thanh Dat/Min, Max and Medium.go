package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
	/*Cấu trúc lại file*/
)
func MinMax(arr [] string) (min int, max int){
	min, _ = strconv.Atoi(arr[0])
	max, _ = strconv.Atoi(arr[0])
	for i := 0; i < len(arr); i++ {
		intVar, _ := strconv.Atoi(arr[i])
		if(intVar < min) {
			min = intVar
		}
		if(intVar > max) {
			max = intVar
		}
	}
	return
}
func medium(arr [] string) (medium int) {
	sum := 0
	for i := 0; i < len(arr); i++ {
		intVar, _ := strconv.Atoi(arr[i])
		sum = sum + intVar
	}
	medium = sum/len(arr)
	return
}
/*func BubbleSort(arr [] string) (arr [] int) {
	i := len(arr)
}*/
func main() {
	content, err := ioutil.ReadFile("a.txt")
	if err != nil {
		log.Fatal(err)
		fmt.Println(string(content))
	}
	arr := strings.Fields(string(content))
	fmt.Println(MinMax(arr))
	fmt.Println(medium(arr))
/*	fmt.Println(BubbleSort(arr))*/


}