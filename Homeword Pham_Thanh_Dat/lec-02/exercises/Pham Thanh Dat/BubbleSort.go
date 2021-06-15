package main

import "fmt"
func BubbleSort(arr [6]int)([6] int) {
	for x:= len(arr); x > 1; x-- {
		for i := 1; i < x; i++ {
			if(arr[i] < arr[i-1]) {
				arr[i], arr[i-1] = arr[i-1], arr[i]
			}

		}
	}
	return arr
}
func main() {
	arr := [6]int{1,9,5,3,7,6}
	fmt.Println(BubbleSort(arr))
}
