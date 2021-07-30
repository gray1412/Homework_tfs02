package dataHandling

import (
	// "fmt"
	"encoding/json"
	"strings"
)

func SplitLineToMap(lineChannel chan string, mapChannel chan string) {
	for {
		line := <-lineChannel
		s, _ := json.Marshal(splitLineToMap(line))
		mapChannel <- string(s)
	}
}

func splitLineToMap(line string) map[string]int {
	m := make(map[string]int)
	arr := strings.Split(line, " ")
	for _, value := range arr {
		_, ok := m[value]
		if ok {
			m[value]++
		} else {
			m[value] = 1
		}
	}
	return m
}

func UpdateMapResult(result *map[string]int, mapReceiveChannel chan string) {
	for {
		newMapString := <-mapReceiveChannel
		var newMap = make(map[string]int)
		json.Unmarshal([]byte(newMapString), &newMap)
		updateMap(result, newMap)
	}
}

func updateMap(old *map[string]int, newInput map[string]int) {
	oldMap := *old
	for key, value := range newInput {
		_, ok := oldMap[key]
		if ok {
			oldMap[key] += value
		} else {
			oldMap[key] = value
		}
	}
}
