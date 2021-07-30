package fileHandling

import (
	"bufio"
	"fmt"
	"os"
)

//pathExample: "./sample.txt"
func ReadFileLineByLine(path string, output chan string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Cannot read file from path: %v", path)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		output <- scanner.Text()
		// fmt.Println(scanner.Text()+ "\n")
	}
}
