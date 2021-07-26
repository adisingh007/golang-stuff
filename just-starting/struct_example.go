package main

import "fmt"

type animal struct {
	name string
	age int
}

func (a animal) Age() int {
	return a.age
}

func main() {
	myDog := animal{"Suresh",35}
	fmt.Println(myDog.Age())
}
