package main 

import (
    "fmt"
    "os"
)

func main() {
    //handling the errors
    //the numberOfBytes and theError variables are scoped only for the if statement
    //the variables numberOfBytes and theError are scoped also for the else statement
    //if _, theError := fmt.Printf("Hello World\n"); -> ignores the num of bytes
    if numberOfBytes, theError := fmt.Printf("Hello World\n"); theError != nil {
        os.Exit(1)
    }else {
        fmt.Printf("Printed %d characters\n", numberOfBytes)
    }
    
    //the alternative way 
    //declaring the numberBytes and err outside the scope of the if
    var numberBytes int
    var err error
    
    if numberBytes, err = fmt.Printf("Hello World\n"); err != nil {
        os.Exit(1)
    }
    
    fmt.Printf("Printed %d characters\n", numberBytes)
    
    
}