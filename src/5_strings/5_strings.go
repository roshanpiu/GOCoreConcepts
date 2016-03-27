package main 

import (
    "fmt"   
)

func main() {
    atoz := "the quick brown fox jumps over the lazy dog\n"
    
    //in backtickStrings the special characters such as \n are ignored
    backtickStrings := `the "quick" brown fox jumps over the lazy dog\n`
    fmt.Printf(backtickStrings)
    fmt.Printf("\n")
    
    //prints the length of the string
    fmt.Printf("the length is %d\n", len(atoz))
    
    
    //using substrings
    fmt.Printf("%s\n", atoz[0:9])   // output -> 'the quick'
    fmt.Printf("%s\n", atoz[:9])    // output -> 'the quick'
    
    fmt.Printf("%s\n", atoz[15:19]) // output -> ' fox'
    fmt.Printf("%s\n", atoz[15:])   // output -> ' fox jumps over the dog'
    
    for i, r := range atoz {
        //i is index
        //r is the character
        fmt.Printf("%d %c\n", i, r)
    }
    
    for _, r := range atoz {
        //index is assigned to underscore because we dont use it
        //r is the character
        fmt.Printf("%c\n", r)
    }
    
    for i := range atoz {
        //this for loop use only the index
        fmt.Printf("%d\n", i)
    }
}