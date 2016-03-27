package main 

import "fmt"

//declaring constants
const (
    answer = 42
    
    //automatically created values
    //answer1's value is 1 
    //answer2's value is 2
    // if answer1 = iota*2 then values will be 2 and 4
    answer1 = iota
    answer2
)

var number = 42

func main() {
    //the long way of declaring and initializing variables
    var message1 string
    message1 = "Hello World"
    
    //declaring and initializing happens at the same time when we follow this short hand notation
    message2 := "Hello"
    
    fmt.Println(message1)
    fmt.Println(message2)
    
    //we can do this but can not assign a different data type to the number (ex - number = "abs")
    number += 10
    
    fmt.Println(message2, answer, number)
    fmt.Println(answer1)
    fmt.Println(answer2)
    
    pi := float64(3.14)
    fmt.Println("Value of pi is ",pi)
    
    nine := 9
    unsignedInt := uint(9)
    unsignedInt8bits := uint8(9)
    fmt.Println(nine, unsignedInt, unsignedInt8bits)
    
    isTrue := true
    var isFalse bool // this is initialized as a 0 in go 0 is equal to false
    //when we declare a variable without initializing it will be initialized with the 0 value
    fmt.Println(isTrue, isFalse)
    
    b := byte(65)
    fmt.Printf("hex value of 65 is %x\n", b) //%x gives the hex value
    
}