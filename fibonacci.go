package main

import "fmt"

func getFibonacci(N int) []int {
	if N == 0 {
		return []int {}
	}

	if N == 1 {
		return []int {1}
	}

	fibonacciSequence := make([]int, N)
	fibonacciSequence[0] = 1
	fibonacciSequence[1] = 1

	for i := 2; i < N; i++ {
		fibonacciSequence[i] = fibonacciSequence[i-1]+fibonacciSequence[i-2]
	}

	return fibonacciSequence
}

func main() {
	var N int
	fmt.Printf("Enter length of desired fibonacci sequence: ")
	fmt.Scanf("%d", &N)
	fmt.Println(getFibonacci(N))
}
