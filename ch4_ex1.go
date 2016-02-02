package main

import (
	"fmt"
)

func UniqueInts(ints []int) []int {
	uniques := []int{ints[0]}
	for _, integer := range ints {
		seen := false
		for _, unique := range uniques {
			if integer == unique {
				seen = true
			}
		}
		if seen == false {
			uniques = append(uniques, integer)
		}
	}
	return uniques
}

func main() {
	vals := []int{9, 1, 9, 5, 4, 4, 2, 1, 5, 4, 8, 8, 4, 3, 6, 9, 5, 7, 5}
	uniques := UniqueInts(vals)
	fmt.Println(uniques)
}
