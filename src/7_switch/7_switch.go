package main 

import (
    "fmt"
    "os"
)

func main() {
    
    n, err := fmt.Printf("Hello World\n")
    
    //there is no fallthrough behavior in go switch statements
    
    switch {
        case err != nil:
            os.Exit(1)
        case n == 0:
            fmt.Printf("No bytes output")
        case n != 12:
            fmt.Printf("Wrong number of characters : %d", n)
        default:
            fmt.Printf("OK!")
        
    }
    
    fmt.Printf("\n")
    
    atoz := "the quick brown fox jumps over the lazy dog"
    
    vowels := 0
    consonants := 0
    zeds := 0
    
    for _, r := range atoz {
        switch r {
            case 'a','e','i','o','u':
                vowels++
            case 'z':
                zeds++
                fallthrough //enforcing the fallthrough behaviour
            default:
                consonants++
        }
    }
    
    fmt.Printf("Vowels: %d; Consonants: %d; Zeds: %d; Length: %d\n", vowels, consonants, zeds, len(atoz))
}