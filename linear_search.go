package main

import (
	"fmt"
)

func main() {
	helo()
}

func helo() {
	numbrs := []int{1, 65, 7, 5, 34, 454, 64, 34}
	target := 64

	indx := search(numbrs, target)

	if indx != -1 {
		fmt.Printf("Element %d found at index %d\n", target, indx)
	} else {
		fmt.Printf("Element %d not found\n", target)
	}
}

func search(array []int, key int) int {
	for k, value := range array {
		if value == key {
			return k // Found, return index
		}
	}
	return -1 // Not found
}
