//Kiểm tra xem các giá trị ở file đầu vào có là số nguyên tố ko
package main

import (
	f "CallPkg/pkgs/handleFile"
	"fmt"
)

func checkPrime(n int64) bool {
	if n <= 1 {
		return false
	} else if n%2 == 0 {
		return true
	} else {
		var i int64 = 3
		for i < n/2 {
			if n%i == 0 {
				return false
			}
			i += 2
		}
	}
	return true
}

func checkPrimeFile(filename string) {
	isPrime := true
	arr, success := f.ReadFile(filename)
	if success == true {
		for _, value := range arr {
			if isPrime == true {
				if checkPrime(value) == false {
					isPrime = false
				}
			} else {
				break
			}

		}
		if isPrime {
			fmt.Println("Tất cả các số trong file đều là số nguyên tố")
		} else {
			fmt.Println("File có chứa số không phải là số nguyên tố")
		}
	}

}

func main() {
	checkPrimeFile("file/test.txt")
}
