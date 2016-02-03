package main

import (
	"fmt"
)

func Flatten(multiArr [][]int) []int {
	flattened := []int{}
	for _, arr := range multiArr {
		for _, item := range arr {
			flattened = append(flattened, item)
		}
	}
	return flattened
}

func main() {
	irregularMatrix := [][]int{{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11},
		{12, 13, 14, 15},
		{16, 17, 18, 19, 20}}
	slice := Flatten(irregularMatrix)
	fmt.Printf("1x%d: %v\n", len(slice), slice)
}
