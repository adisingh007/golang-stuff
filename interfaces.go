package main

import "fmt"

type Noisy interface {
	makeNoise() string
}

type Dog struct {
	name string
}

func (d Dog) makeNoise() string {
	return d.name + " Barks"
}

func main() {
	myDog := Dog{"Matthieu"}
	fmt.Println(myDog.makeNoise())
}
