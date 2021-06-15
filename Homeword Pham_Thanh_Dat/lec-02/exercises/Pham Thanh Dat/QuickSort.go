package main

import "fmt"
func Sort(arr []int, low int, high int)(int) {
	i := low - 1
	for j := low; j<= high; j++{
		if(arr[j] <= arr[high]) {
			i++
			arr[i],arr[j] = arr[j], arr[i]
		}
	}
	/*fmt.Println(arr)*/
	return i
}
func QuickSort(arr []int, low int, high int) {
	if low < high {
		pivot := Sort(arr, low, high)
		/*fmt.Println(pivot, pivot-1, pivot+1)*/
		QuickSort(arr, low, pivot-1)
		QuickSort(arr, pivot+1, high)
	}

}
func main() {
	arr := []int{1,9,5,4,7,6,3,5,2,4,8,7,6,9,2,1,5,4,6,8,9,6,5,2,1,4,2,5}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}