package main

import "fmt"

type Person struct {
	name string
	age int
}

type Engineer struct {
	Person
	engineerLanguage string
}

func main() {
	engineer := Engineer{
		Person: Person {
			age: 26,
			name: "Aditya Singh",
		},
		engineerLanguage: "Golang",
	}
	fmt.Println(engineer)
	fmt.Println(engineer.age)
	fmt.Println(engineer.name)
	fmt.Println(engineer.engineerLanguage)
}
