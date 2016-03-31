package main

import (
	"fmt"
)

func main() {

	//take a look at the formatting options in godoc

	fmt.Printf("%v\n", "Hello World!")
	fmt.Printf("%#v\n", "Hello World!") //go syntax version
	fmt.Printf("%#v\n", 1.23)           //go syntax version
	fmt.Printf("%T\n", 1.23)            //go type

	fmt.Printf("% x\n", "Hello World!") //hex version
}
