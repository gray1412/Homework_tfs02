package sort_algorithm

// BubbleSort
func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// QuickSort
func QuickSort(arr []int, low int, high int) {
	if low < high {
		pointer := partition(arr, low, high)
		QuickSort(arr, low, pointer-1)
		QuickSort(arr, pointer+1, high)
	}
}
func partition(arr []int, low int, high int) int {
	pivot, left, right := arr[high], low, high-1
	for {
		for left <= right && arr[left] < pivot {
			left++
		}
		for right >= left && arr[right] > pivot {
			right--
		}
		if left >= right {
			break
		}
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	arr[left], arr[high] = arr[high], arr[left]
	return left
}

// Merge Sort
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	length := len(arr) / 2
	left := MergeSort(arr[:length])
	right := MergeSort(arr[length:])
	return merge(left, right)
}
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(result, right...)
		}
		if len(right) == 0 {
			return append(result, left...)
		}
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	return result
}
