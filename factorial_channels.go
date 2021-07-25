package main

import "fmt"

func factorial(n int, c chan int) {
	x := 1
	for i := 2; i <= n; i++ {
		c <- x
		x = x*i
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go factorial(10, c)
	for i := range c {
		fmt.Println(i)
	}
}
