package main

import "fmt"

func getSomething(stringChannel chan string, channelNumber int) {
	fmt.Printf("In the getSomething() function of channelNumber %d\n", channelNumber)
	stringChannel <- fmt.Sprintf("Hello, I am channel #%d", channelNumber)
	fmt.Printf("Msg sent to channel #%d\n", channelNumber)
}

func main() {
	stringChannel := make(chan string, 2)

	go getSomething(stringChannel, 1)
	fmt.Println(<- stringChannel)
	go getSomething(stringChannel, 2)
	fmt.Println(<- stringChannel)
}
