package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("%d => %s\n", i, fizzBuzz(i))
	}
}

func fizzBuzz(Num int) string {
	var result string = ""
	if Num % 3 == 0 {
		result += "Fizz"
	}
	if Num % 5 == 0 {
		result += "Buzz"
	}
	return result
}
