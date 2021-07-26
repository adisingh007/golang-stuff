package main

import "fmt"

func main() {
	var num int
	for true {
		fmt.Scanf("%d", &num)
		if num == 42 {
			break
		}
		fmt.Printf("%d\n", num)
	}
}
