package main 

import (
    "flag"
    "fmt"
)

func main() {
    name := flag.String("name", "World", "Name to print")
    
    var name2 string
    flag.StringVar(&name2, "name2", "World", "Name to print")
    
    flag.Parse()
    
    fmt.Println("Hello ", *name)
    fmt.Println("Hello ", name2)
    
    //go run 34_flags.go -name=roshan to pass command line args
    //go run 34_flags.go -name2=roshan to pass command line args
}