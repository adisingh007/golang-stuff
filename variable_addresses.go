package main

import "fmt"

func main() {
	i, j := 24, 2701
	fmt.Println("i = ", i, "j = ", j)
	fmt.Println("&i = ", &i, "&j = ", &j)

	p := &i
	fmt.Println("p = ", p)
	fmt.Println("*p = ", *p)

	fmt.Println(willUseMemoryHeap()) // Haha WTF is this? Why would this print the value of the object and not the real address?
	fmt.Println(*willUseMemoryHeap())
}

type Person struct {
	name string
	age int
}


func willUseMemoryHeap() *Person {
	somePersonInMemoryHeap := Person{"Some anonymous joe", 56}
	return &somePersonInMemoryHeap
}
