package main

import (
	"fmt"
)

func fib(n int, memo map[int]int) int {
	if val, found := memo[n]; found {
		return val
	}

	if n <= 1 {
		return n
	}

	memo[n] = fib(n-1, memo) + fib(n-2, memo)
	return memo[n]
}

func main() {
	memo := make(map[int]int)
	n := 7
	result := fib(n, memo)
	fmt.Println("Fibonacci number at position", n, "is:", result)
}
