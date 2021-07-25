package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("%d => %s\n", i, fizzBuzz(i))
	}
}

func fizzBuzz(Num int) string {
	if Num % 3 == 0 && Num % 5 == 0 {
		return "FizzBuzz"
	}
	if Num % 3 == 0 {
		return "Fizz"
	}
	if Num % 5 == 0 {
		return "Buzz"
	}
	return ""
}
