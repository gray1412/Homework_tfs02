package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	numbers := ReadFile()
	fmt.Print("Max, min, avg: ")
	fmt.Println(min_max_avg(numbers))
	fmt.Print("sort: ")
	fmt.Println(bubbleSort(numbers, true))
	fmt.Print("prime in numbers: ")
	fmt.Println(findPrime(numbers))

	fmt.Print("search value in file: ")
	fmt.Println(searchValue(numbers, 7))

	writeFile("newFile.txt", "Hello World 123")
}

func ReadFile() []int {
	//Read content in file number.txt
	data, err := ioutil.ReadFile("number.txt")
	if err != nil {
		fmt.Println(err)
	}

	// get number form data of file
	number := strings.Split(string(data), " ")

	var listNum = []int{}

	// change value string to array
	for _, value := range number {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			println(err)
		}
		listNum = append(listNum, intValue)
	}
	return listNum
}

func writeFile(fileName string, content string) {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(content)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
}

func min_max_avg(numbers []int) (int, int, int) {
	// find max, min, avg
	max := numbers[0]
	min := numbers[0]
	s := 0
	for _, value := range numbers {
		s += value
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}

	avg := s / len(numbers)
	return max, min, avg
}

// func inout about array, orinet, orient true sort up, orient true sort dow
func bubbleSort(numbers []int, orient bool) []int {
	sort := 1
	if !orient {
		sort = -1
	}
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers)-1; j++ {
			if numbers[i]*sort > numbers[j]*sort {
				intermediate := numbers[i]
				numbers[i] = numbers[j]
				numbers[j] = intermediate
			}
		}
	}
	return numbers
}

func prime(n int) bool {
	sqr := math.Sqrt(float64(n))
	// index := 0
	if n < 2 {
		return false
	}
	for i := 2; i <= int(sqr); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func findPrime(numbers []int) bool {
	for _, value := range numbers {
		if prime(value) {
			return true
		}
	}
	return false
}

func searchValue(numbers []int, n int) bool {
	for _, value := range numbers {
		if value == n {
			return true
		}
	}
	return false
}
