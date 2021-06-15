//Tìm min, max, avg của file
package main

import (
	f "CallPkg/pkgs/handleFile"
	"fmt"
)

func handleArr(a []int64) (max, min int64, avg float64) {
	max = a[0]
	min = a[0]

	var sum int64 = 0

	for _, value := range a {
		sum += value
		if value > max {
			max = value
		} else {
			min = value
		}
	}
	avg = float64(sum) / float64(len(a))
	return
}
func checkFile(filename string) {
	arr, success := f.ReadFile(filename)
	if success == false {
		fmt.Println("Can't read file !")
	} else {
		max, min, avg := handleArr(arr)
		fmt.Printf("\nmax = %v", max)
		fmt.Printf("\nmin = %v", min)
		fmt.Printf("\navg = %v", avg)
	}

}

func main() {
	checkFile("file/test.txt")
}
