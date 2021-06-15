package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := strings.Fields(string(content))
	fmt.Println(string(content))
	MinMax(data)
	CheckValueInFiled(data)

}
func MinMax(data []string) {
	max := data[0]
	min := data[0]
	sum := 0
	for i := 0; i < len(data); i++ {
		if data[i] > max {
			max = data[i]
		}
		if data[i] < min {
			min = data[i]
		}
		number, _ := strconv.Atoi(data[i])
		sum = sum + int(number)
		fmt.Println(number)
		fmt.Println(CheckPrime(number))
	}
	fmt.Println("Max is: ", max)
	fmt.Println("Min is: ", min)
	fmt.Println("Sum is: ", sum)
	fmt.Println("The average value: ", sum/len(data))
}

func CheckPrime(number int) bool {
	if number < 2 {
		return false
	}
	if number == 2 {
		return true
	}
	if number > 2 {
		for j := 2; j < number; j++ {
			if number%j == 0 {
				return false
				break
			}
		}
	}
	return true
}
func CheckValueInFiled(data []string) {
	fmt.Print(data)
	for {
		var chose int
		fmt.Print("Enter your chose, 1- check ||0 - stop: ")
		fmt.Scanln(&chose)
		if chose == 0 {
			break
		}
		if chose == 1 {
			var Number int
			fmt.Print("Enter your Number: ")
			fmt.Scanln(&Number)
			check := false
			for i := 0; i < len(data); i++ {
				value, _ := strconv.Atoi(data[i])
				if value == Number {
					fmt.Printf(" Number %v had in filed: \n", Number)
					check = true
					break
				}
			}
			if check == false {
				fmt.Printf("Number %v not had in filed: \n", Number)
			}

		}

	}
}
