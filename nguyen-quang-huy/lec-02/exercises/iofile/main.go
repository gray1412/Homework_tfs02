package main

import (
	"fmt"
	"strings"
	"tfs-02/lec-02/exercises/iofile/controllers"
	"tfs-02/lec-02/exercises/iofile/helpers"
)

func main() {
	pathRead := "lec-02\\exercises\\iofile\\data\\number.txt"

	// Đọc File
	rawData, err := helpers.ReadFile(pathRead)
	if err != nil {
		fmt.Println("Error ", err)
		return
	}
	fmt.Println()
	data, err := controllers.ReadNumber(strings.NewReader(rawData))

	fmt.Println(data, err)
	// tim so lon nhat
	max := controllers.FindMax(data)
	fmt.Println("max: ",max)

	// tim so nho nhat
	min := controllers.FindMin(data)
	fmt.Println("min: ", min)

	// trung binh
	avg := controllers.Average(data)
	fmt.Printf("avg: %0.2f \n", avg)

	// ktra snt
	for key, _ := range data {
		if controllers.IsPrimeNumber(key) {
			fmt.Println(key, "la snt")
		} else {
			fmt.Println(key, "khong la snt")
		}
	}

	// so ton tai
	arr := []float64{1,8}
	for i := 0; i < len(arr); i++ {
		if controllers.ExistNumber(arr[i], data) {
			fmt.Println(arr[i], "ton tai trong file")
		} else {
			fmt.Println(arr[i], "ko ton tai trong file")
		}
	}

	// Ghi file
	dataWrite := "Nguyen Quang Huy"
	pathWrite := "lec-02\\exercises\\iofile\\data\\write.txt"

	err = helpers.WriteFile(pathWrite,dataWrite)
	if err == nil {
		fmt.Println("Ghi thanh cong")
	} else{
		fmt.Println("Ghi ko thanh cong")
	}
}