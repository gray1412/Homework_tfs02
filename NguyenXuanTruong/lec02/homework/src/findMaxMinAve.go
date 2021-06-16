package main

import (
	"fmt"
	"strconv"
	"strings"
	"io/ioutil"
)

func findMax(arr []int) int{
	// if len(arr) = 0 {
	// 	fmt.Println("Mảng rỗng")
	// 	return -1
	// } 
	max := arr[0]
	for _, i := range arr {
		if i > max {
			max = i
		}
	}
	return max
}

func findMin(arr []int) int{
	// if len(arr) = 0 {
	// 	fmt.Println("Mảng rỗng")
	// 	return -1
	// } 
	min := arr[0]
	for _, i := range arr {
		if i < min {
			min = i
		}
	}
	return min
}

func averaged(arr []int) int{
	// if len(arr) = 0 {
	// 	fmt.Println("Mảng rỗng")
	// 	return -1
	// } 
	ave := 0
	for _, i := range arr {
		ave = ave + i
	}
	return ave
}

func handleFile(filename string) {
	arr := []int{}

	// Đọc file input.txt
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	//Duyệt các phần tử file input
	for _, i := range strings.Fields(string(data)) {
		num, _ := strconv.Atoi(i)
		arr = append(arr, num)
	}

	fmt.Println(findMax(arr))
	fmt.Println(findMin(arr))
	fmt.Println(averaged(arr))
}

func main() {
	handleFile("file/input.txt")
}