package findNumber

import (
	f "CallPkg/pkgs/handleFile"
)

func FindNumbyWay1(n int64, filename string) bool {
	arr, success := f.ReadFile(filename)
	if !success {
		return false
	}
	ret := false
	for _, value := range arr {
		if n == value {
			ret = true
			break
		}
	}
	return ret
}

//tìm kiếm sử dụng hàm băm
func hash(n int64) (a, b, c int64) {
	a = (n * 79 / 23 * 2) % 1000
	b = (n * 81 / 25 * 3) % 1000
	c = (n * 85 / 20 * 3) % 1000
	return
}
func hashArr(a []int64) (ret [1000]int64) {
	for _, value := range a {
		a, b, c := hash(value)
		ret[a] = 1
		ret[b] = 1
		ret[c] = 1
	}
	return
}
func FindNumByHashing(n int64, filename string) bool {

	arr, success := f.ReadFile(filename)

	if !success {
		return false
	}

	hashArr := hashArr(arr)
	a, b, c := hash(n)

	if hashArr[a] == 1 && hashArr[b] == 1 && hashArr[c] == 1 {
		return true
	}

	return false
}
