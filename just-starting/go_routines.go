package main

import "fmt"

func display(str string) {
	for w := 1; w <= 100; w++ {
		fmt.Printf("%s: %d\n", str, w)
	}
}


func main() {
	go display("Go routine")
	go display(" Another Go routine")

	display("Main function")
	fmt.Println("Scroll up and notice that 'Go routine' may have without even printing 100 times.")
	fmt.Println("If it has printed 100 times, re-run. Re-run 10 times to notice.")
	fmt.Println("This shows that if 'main' routine ends, all the created sub-routines end.")
	fmt.Println("-- End --")
}
