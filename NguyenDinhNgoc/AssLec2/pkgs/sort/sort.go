//Các thuật toán sắp xếp
package sort

//BubbleSort
func BubbleSort1(a []int64) {
	for i := 0; i < len(a); i++ {
		for j := len(a) - 1; j > i; j-- {
			if a[j] < a[i] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
func BubbleSort2(a []int64) {

	var done bool = false

	for !done {
		done = true
		for i := len(a) - 1; i > 0; i-- {
			if a[i] < a[i-1] {
				a[i], a[i-1] = a[i-1], a[i]
				done = false
			}
		}
	}
}

//QuickSort
func quickSort_partition(arr []int64, low, high int) int {
	pivot := arr[low]
	i := low + 1
	j := i

	for ; j <= high; j++ {
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[low], arr[i-1] = arr[i-1], arr[low]

	return i - 1
}
func QuickSort(sli []int64, low, high int) {
	if low < high {
		pivot := quickSort_partition(sli, low, high)
		QuickSort(sli, low, pivot-1)
		QuickSort(sli, pivot+1, high)
	}
}

func merge(arr []int64, start1, end1, start2, end2 int) {
	var tmp []int64
	i, j := start1, start2
	for m := 0; m < len(arr); m++ {
		if i > end1 {
			tmp = append(tmp, arr[j])
			j++
		} else if j > end2 {
			tmp = append(tmp, arr[i])
			i++
		} else if arr[i] <= arr[j] {
			tmp = append(tmp, arr[i])
			i++
		} else {
			tmp = append(tmp, arr[j])
			j++
		}
	}
	copy(arr, tmp)
}

// func MergeSort(arr []int64, low, high int) {
// 	if low < high {
// 		mid := (high-low)/2 + low
// 		MergeSort(arr, low, mid)                   //Error
// 		MergeSort(arr, mid+1, high)
// 		merge(arr, low, mid, mid+1, high)
// 	}
// }
