//Các thuật toán sắp xếp
package main

import (
	s "CallPkg/pkgs/sort"
	"fmt"
)

func main() {

	var a = []int64{4, 3, 1, 6, 2, 7, 9, 4, 0, 6, 5}

	s.BubbleSort1(a)
	// s.BubbleSort2(a)
	// s.QuickSort(a, 0, len(a)-1)
	// s.MergeSort(a, 0, len(a)-1)

	fmt.Println(a)
}
