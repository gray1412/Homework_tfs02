package findNum

import (
	"fmt"
	"strings"
	"io/ioutil"
	"strconv"
)

func Checking(namefile string, element int) bool{
	data, err := ioutil.ReadFile(namefile)
	if err != nil {
		fmt.Println(err)
	}

	isHaving := false
	myMap := make(map[string]int)
	estr := strconv.Itoa(element)

  	for _, j := range strings.Fields(string(data)) {
  		myMap[j]++
  	}

	if myMap[estr] > 0 {
		return isHaving
	}
	return isHaving
}
