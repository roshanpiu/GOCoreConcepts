//empty interfaces can be useful it gives you a way to have an unknown type which determines at runtime

//avoid empty interfaces if you can
//if you use them do a type assertion and make it safe

package main 

import (
    "fmt"
)

func whatIsThis(i interface{})  {
    switch v := i.(type) {
        case string:
            fmt.Printf("it is a string %s\n", v)
        case uint32:
            fmt.Printf("it is a uint32 %d\n", v)
        default:
            fmt.Printf("I don't know this type %v\n", v)
    }
    //fmt.Printf("%T\n",i)
}

func whatIsThisTypeAssersion(i interface{})  {
    switch i.(type) {
        case string:
            fmt.Printf("it is a string %s\n", i.(string))
        case uint32:
            fmt.Printf("it is a uint32 %d\n", i.(uint32))
            //fmt.Printf("it is a uint32 %d\n", i.(string)) //this will create a runtime panic
        default:
            fmt.Printf("I don't know this type\n")
    }
}

func main() {
    whatIsThis(42)
    whatIsThis("42")
    whatIsThis(uint32(42))
    whatIsThis([...]string{"A","B","C"})
    
    whatIsThisTypeAssersion("Hello")
}