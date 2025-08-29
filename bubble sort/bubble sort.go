package main

import (
	"fmt"
)

func bubblesort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	array := []int{1, 2, 53, 64, 234, 53, 234, 789}
	result := bubblesort(array)
	fmt.Println(result) // Or use fmt.Printf("%v\n", result)
}
