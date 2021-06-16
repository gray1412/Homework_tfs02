package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"io/ioutil"
)

func checkPrime(n int) bool{
	var isPrime bool = true
	// Số nhỏ hơn 2 không phải số nguyên tố
	if n < 2 {
		isPrime = false
		return isPrime
	}

	// Số có nhiều hơn 2 ước 1 và chính nó không phải số nguyên tố
	// Chỉ kiểm tra các số nhỏ hơn sqrt(n) có phải ước hay không
	// Kiểm tra số đã cho có là bội của 2 không, nếu không loại bỏ kiểm tra có phải bội của các số chẵn khác không
	if n % 2 == 0 {
		isPrime = false
		return isPrime
	}

	can := int(math.Sqrt(float64(n)))
	for i:=3; i < can; { //chỉ kiểm tra đến n/2
		if can % i == 0 {
			isPrime = false
			return isPrime
		}
		i = i+2 // Không kiểm tra lại các số chẵn, bởi vậy bước nhảy là 2
	}
	return isPrime
}

func checkPrimeInFile(filename string) {
	// Đọc file input.txt
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	//Duyệt các phần tử file input
	for _, i := range strings.Fields(string(data)) {
		num, _ := strconv.Atoi(i)
		if checkPrime(num) == true {
			fmt.Println("%v là số nguyên tố", num)
		} else {
			fmt.Println("%v không phải là số nguyên tố", num)
		}
	}
}

func main() {
	checkPrimeInFile("file/input.txt")
}