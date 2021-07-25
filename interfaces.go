package main

import "fmt"

type Noisy interface {
	makeNoise() string
}

type Dog struct {
	name string
}

type Cat struct {
	name string
	age int
	toy string
}

func (d Dog) makeNoise() string {
	return fmt.Sprintf("%s barks", d.name)
}

func (c Cat) makeNoise() string {
	return fmt.Sprintf("%d year old %s meows with his %s", c.age, c.name, c.toy)
} 

func main() {
	myDog := Dog{"Matthieu"}
	myCat := Cat{"Mottet", 40, "PHd"}
	noisies := []Noisy{myDog, myCat}
	for _,p := range noisies {
		fmt.Println(p.makeNoise())
	}
}
