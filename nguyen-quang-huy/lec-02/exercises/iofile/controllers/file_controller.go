package controllers

import (
	"bufio"
	"io"
	"math"
	"strconv"
)

func ExistNumber(number float64, data map[float64]int64) bool {
	_, ok := data[number]
	return ok
}
func IsPrimeNumber(number float64) bool {
	if number <= 0 || number != float64(int64(number)) {
		return false
	}
	if number == 1 || number == 2 {
		return true
	}
	for i := 2.0; i <= math.Sqrt(number); i++ {
		if math.Mod(number, i) == 0 {
			return false
		}
	}
	return true
}
func Average(data map[float64]int64) float64 {
	sumData := 0.0
	var sumIndex int64 = 0
	for key, value := range data {
		sumData += key * float64(value)
		sumIndex += value
	}
	return sumData / float64(sumIndex)

}
func FindMin(data map[float64]int64) float64 {
	var min float64
	for key, _ := range data {
		min = key
		break
	}
	for key, _ := range data {
		if key < min {
			min = key
		}
	}
	return min
}
func FindMax(data map[float64]int64) float64 {
	var max float64
	for key, _ := range data {
		max = key
		break
	}
	for key, _ := range data {
		if key > max {
			max = key
		}
	}
	return max
}
func ReadNumber(r io.Reader) (map[float64]int64, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	result := make(map[float64]int64)
	for scanner.Scan() {
		index, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		result[index]++
	}
	return result, scanner.Err()
}
