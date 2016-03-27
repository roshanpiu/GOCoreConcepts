package main 

import (
    "fmt"
)

//w [4]string is an array
//w []string is a slice
func arrayPrinter(w [4]string)  {
    for _, word := range w{
        fmt.Printf("%s ", word)
    }
    fmt.Printf("\n")
    w[2] = "blue"
    fmt.Println(w)
}

func slicePrinter(w []string)  {
    //the original slice's value will be changed because they are passed by reference
    w[2] = "blue"
    fmt.Println(w)
}

func main() {
    
    //array declaration of unknows size
    //arrays are passed by value by default (we can pass by reference when we specify it explicitly)
    words1 := [...]string{"the","quick","brown","fox"}
    words2 := [4]string{"the","quick","brown","fox"} //when size is known
    fmt.Printf("%s\n",words1[2])
    fmt.Printf("%s\n",words2[2])
    
    arrayPrinter(words2)
    fmt.Println(words2)
    
    //sclices
    //sclices are passed by reference
    //passing slices around functions is cheap
    words3 := []string{"the","quick","brown","fox","jumps","over","the","lazy","dog"}
    fmt.Println("before passing:", words3)
    slicePrinter(words3)
    fmt.Println("after passing:", words3)
    fmt.Println("number of words: ",len(words3))
    
    //slicing slices
    words4 := words3[5:]
    words5 := words3[:5]
    fmt.Println("after slicing words3:", words4)
    fmt.Println("after slicing words3:", words5)
    
    //using make() to create slices 0 is length 4 is capacity
    //capacity can dynamically grow when we append to slices
    //when appending if the capacity is not enough go will double it
    words6 := make([]string, 0, 4)
    words6 = append(words6, "The ")
    words6 = append(words6, "quick ")
    words6 = append(words6, "brown ")
    words6 = append(words6, "fox ")
    fmt.Println(words6, "length: ", len(words6), "capacity: ", cap(words6))
    words6 = append(words6, "jumps ")
    //words6[5] = "jumps" we cannot do this when the capacity limit is reached we should use append
    fmt.Println(words6, "length: ", len(words6), "capacity: ", cap(words6))
    
    //copying slices
    sliceCopy := make([]string, 4)
    copy(sliceCopy, words6)
    fmt.Println(sliceCopy)
    fmt.Println()
    fmt.Println()
    
    //when we change the value of sliceCopy1 the value of words6 changes
    //because both the slices points to the same underlying array
    sliceCopy1 := words6[0:4]
    fmt.Println(sliceCopy1)
    fmt.Println(words6)
    fmt.Println()
    sliceCopy1[2] = "blue"
    fmt.Println(sliceCopy1)
    fmt.Println(words6)
}