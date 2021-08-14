package main

import (
	"fmt"
	"time"
)

func listen(ch chan string) {
	fmt.Println(<- ch)
	ch <- "No, I'm straight!"
}

func main() {
	ch := make(chan string)
	go listen(ch)
	ch <- "Do you speak french?"
	fmt.Println(<- ch)
	time.Sleep(time.Second)
}