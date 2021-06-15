package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

func main() {
	//array := []int{3, 4, 5, 2, 1, 7, 8, -1, -3}
	// fmt.Println(bubbleSort(array))
	// fmt.Println(mergeSort(array))
	// fmt.Println(quickSort(array))

	// getHandler := http.HandlerFunc(getEquation)
	// http.Handle("/hello", getHandler)
	// http.ListenAndServe(":8090", nil)

	// readFile()
	// isExist("data.txt", array)
	// check.IsExist("data.txt", array)
	writeToFile("data1.txt", "hello123")
}

// bubble sort
func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// merge sort
func mergeSort(arr []int) []int {
	var length = len(arr)

	if length == 1 {
		return arr
	}

	middle := int(length / 2)
	var (
		left  = make([]int, middle)
		right = make([]int, length-middle)
	)
	for i := 0; i < length; i++ {
		if i < middle {
			left[i] = arr[i]
		} else {
			right[i-middle] = arr[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}

// quick sort
func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1

	pivot := rand.Int() % len(arr)

	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i, _ := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	quickSort(arr[:left])
	quickSort(arr[left+1:])

	return arr
}

// API
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

type Input struct {
	OP     string `json:"OP"`
	A      int32  `json:"a"`
	B      int32  `json:"b"`
	Result int32  `json:"ans"`
	Fault  string `json:"fault"`
}

func getEquation(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	s := query.Get("s")
	a := query.Get("a")
	b := query.Get("b")
	w.WriteHeader(200)

	new_a, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println(err)
	}
	new_b, err := strconv.Atoi(b)
	if err != nil {
		fmt.Println(err)
	}
	var ans int32
	var fault string
	if s == "sum" {
		ans = int32(new_a) + int32(new_b)
		fault = "No Fault"
	} else if s == "sub" {
		ans = int32(new_a) - int32(new_b)
		fault = "No Fault"
	} else if s == "add" {
		ans = int32(new_a) * int32(new_b)
		fault = "No Fault"
	} else if s == "div" {
		if new_b > 0 {
			ans = int32(new_a) / int32(new_b)
			fault = "No Fault"
		} else {
			ans = 0
			fault = "Wrong b"
		}
	} else {
		ans = 0
		fault = "Wrong equation"
	}
	i := Input{s, int32(new_a), int32(new_b), ans, fault}

	bb, err := json.Marshal(&i)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, "ha")
		return
	}
	fmt.Fprintln(w, string(bb))
}

//read file to find min, max, avarage
func readFile() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var max = -9223372036854775808
	var min = 9223372036854775807
	var avarage = 0
	var count = 0

	scanner := bufio.NewScanner(file)
	var primes []int

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
		}
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
		avarage += num
		count++
		if checkIsPrime(num) {
			primes = append(primes, num)
		}
	}
	fmt.Println(max)
	fmt.Println(min)
	fmt.Println(avarage / count)
	fmt.Println(primes)
}

//check is Prime
func checkIsPrime(num int) bool {
	if num < 2 {
		return false
	}
	if num == 2 {
		return true
	}

	if num%2 == 0 {
		return false
	}

	for i := 3; i < num/2; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}

//check is exist
func isExist(path string, arr []int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	m := make(map[int]int)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
		}
		if m[num] != 0 {
			m[num]++
		} else {
			m[num] = 1
		}

	}
	fmt.Println("Exist Number:")
	for i := 0; i < len(arr); i++ {
		if m[arr[i]] != 0 {
			fmt.Println(arr[i])
		}
	}
}

//Write to file
func writeToFile(filename string, newcontent string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)

	if err == nil {
		fmt.Println("\nFile was found and content was written to file !")
		_, err = fmt.Fprintln(file, "\n"+newcontent)
		file.Close()
	} else {
		file.Close()
		file2, err2 := os.Create(filename)
		if err2 != nil {
			fmt.Println("\nFile not found and Can't create newfile !")
			return
		}
		_, err3 := file2.WriteString(newcontent)
		if err3 != nil {
			fmt.Println("\nCreated newfile but can't write to newfile !")
			file2.Close()
			return
		}
	}
}
