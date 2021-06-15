package checkExist

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func IsExist(path string, arr []int) {
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
