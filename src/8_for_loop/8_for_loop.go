package main

import "fmt"

func main() {

	//this is an infinite loop without the break
	for {
		fmt.Println("Hello World")
		break
	}
	fmt.Printf("\n\n")

	for i := 1; i <= 3; i++ {
		fmt.Printf("Hello World %d\n", i)

	}
	fmt.Printf("\n\n")

	for i, j := 1, 2; i <= 3; i, j = i+1, j*2 {
		fmt.Printf("Hello World %d %d\n", i, j)

	}
	fmt.Printf("\n\n")

	var stop bool //stop initialized with false
	i := 0

	for !stop {
		fmt.Printf("Hello World %d\n", i)
		i++
		if i == 5 {
			stop = !stop
		}
	}

}
