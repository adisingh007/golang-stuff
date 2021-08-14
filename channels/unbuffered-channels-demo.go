package main

import (
	"fmt"
	"time"
)

func listen(ch chan string) {
	fmt.Println(<- ch)
	ch <- "No, I'm straight!"
	time.Sleep(time.Second)
}

func main() {
	ch := make(chan string)
	go listen(ch)
	ch <- "Do you speak french?"
	time.Sleep(time.Second)
	fmt.Println(<- ch)
}