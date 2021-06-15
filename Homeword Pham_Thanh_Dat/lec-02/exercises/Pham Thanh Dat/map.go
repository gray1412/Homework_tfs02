package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
)


func main() {
	content, err := ioutil.ReadFile("a.txt")
	if err != nil {
		log.Fatal(err)
		fmt.Println(string(content))
	}
	arr := strings.Fields(string(content))
	number := 0
	var s []int
	for i:= 0; i< len(arr); i++ {
		number, _ = strconv.Atoi(arr[i])
		s = append(s, number)
	}
	m := make(map[int]int)
	for i:= 0; i< len(s); i++ {
		m[s[i]] = 1
	}

	for {
		fmt.Println("Enter the number you need to find")
		var enterNumber int
		fmt.Scanln(&enterNumber)
		fmt.Println(m[enterNumber])
		if(m[enterNumber] == 1){
			fmt.Println("Yes")
		}
		if(m[enterNumber] == 0){
			fmt.Println("No")
		}
		fmt.Println("Do you want to search further, enter any key if yes and no if not")
		var stop string
		fmt.Scanln(&stop)
		if(stop == "no") {
			break
		}
	}
	
}