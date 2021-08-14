package main

import (
	"fmt"
	"time"
)

func writeToCh(ch chan int, num int) {
	for i := 1; i <= num; i++ {
		ch <- i
		fmt.Println("Successfully wrote", i, "to ch")
	}
	close(ch)
}


func main() {
	ch := make(chan int, 2)
	go writeToCh(ch, 10)
	time.Sleep(time.Second)
	for v := range ch {
		fmt.Println("read value ", v, "from ch")
		time.Sleep(time.Second)
	}
}