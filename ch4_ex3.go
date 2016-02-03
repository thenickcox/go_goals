package main

import "fmt"

func Make2D(slice []int, columnCount int) [][]int {
	container := [][]int{}
	for i := 0; i < len(slice); i += columnCount {
		inner := []int{}
		for j := 0; j < columnCount; j++ {
			item := slice[i] + j
			if item > slice[len(slice)-1] {
				item = 0
			}
			inner = append(inner, item)
		}
		container = append(container, inner)
	}
	return container
}

func main() {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	res := Make2D(ints, 3)
	fmt.Printf("3 %v\n", res)
}
