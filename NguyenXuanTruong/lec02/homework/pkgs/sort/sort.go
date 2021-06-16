package sort

func partition(arr []int, low int, high int) int {
	// Chọn pivot là phần tử cuối
	// Các phần tử nhỏ hơn pivot: đưa về bên trái pivot
	// Các phần tử lớn hơn pivot: đưa về bên phải pivot
	pivot := arr[high]
	i := low-1 //Chỉ số thể hiện vị trí của phần tử nhỏ hơn pivot
	for j := low; j <= high- 1; j++ {
		// Nếu phần tử hiện tại nhỏ hơn pivot
		if arr[j] < pivot {
			i++ // Cập nhật vị trí, đây là vị trí mà phần tử này sẽ chuyển tới
			arr[i], arr[j] = arr[j], arr[i] 
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1] // Đưa pivot về vị trí giữa 2 phần: phần nhỏ hơn pivot và phần lớn hơn pivot
	return i+1
}

// low: vị trí bắt đầu, high: vị trí kết thúc của phần mảng thực hiện Quick Sort
func QuickSort(arr []int, low int, high int) {
	if low < high {
		pivot := partition(arr, low, high) // pivot là vị trí phân chia mảng thành 2 phần, nó cũng chính là vị trí của phần tử arr[pivot] sau khi sắp xếp xong
		// Thực hiện đệ quy QuickSort với 2 phần mảng bị chia ra
		QuickSort(arr, low, pivot-1)
		QuickSort(arr, pivot+1, high)
	}
}

func BubbleSort(arr []int) []int {
    for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if (arr[j] > arr[j+1]){
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
    return arr   
}

//func MergeSort(arr)

