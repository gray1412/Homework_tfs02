package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("test.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Write to file: ")
	r := bufio.NewReader(os.Stdin)
	number, _ := r.ReadString('\n')
	text, err := file.WriteString(number)
	if err != nil {
		fmt.Print(err)
		return
	}
	file.Close()
	fmt.Println(text, "have writen to file")

}
