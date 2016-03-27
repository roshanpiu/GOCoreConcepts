package main

import (
    "fmt"
    "os"
)

func main() {
    f, err := os.Open("test.txt")
    if err != nil {
        fmt.Printf("%s\n",err) //error type will get turned into a string
        os.Exit(1)
    }
    defer f.Close()
    
    b := make([]byte, 100)
    
    //here when we use := go will look if we had declared the variable if not it will declare them
    //here is we assign _ to n like '_, err := f.Read(b)' this we have to use = not := like '_, err = f.Read(b)' because no new variables are declared
    n, err := f.Read(b)
    
    fmt.Printf("%d: % x\n",n,b)
    
    //convert byte slice to a string
    stringVersion := string(b)
    fmt.Printf("%d: %s\n",n,stringVersion)
}