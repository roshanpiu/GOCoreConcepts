//go is not built around exceptions it is built around handling errors returned by functions

package main 

import (
    "errors"
    "fmt"
    "os"
    
)

var(
    errorEmptyString = errors.New("Cannot print an empty string")
)

func printer1(msg string) error  {
    
    if msg == ""{
        return fmt.Errorf("Cannot print an empty string")
    }
    
    _, err := fmt.Printf("%s\n",msg)
    return err
}

func printer2(msg string) error  {
    
    if msg == ""{
        return errorEmptyString
        //panic(errorEmptyString) use this panic machanism when its absolutely neccessary 
        //always try to error checking 
    }
    
    _, err := fmt.Printf("%s\n",msg)
    return err
}

func main() {
    if err := printer1("Hello World"); err != nil {
        fmt.Printf("Printer failed %s\n",err)
        os.Exit(1) //when program exits the default return value is 0
    }
    
    //this code will throw an error
    // if err := printer1(""); err != nil {
    //     fmt.Printf("Printer failed %s\n",err)
    //     //os.Exit(1)
    // }
    
    //this code uses the printer2 function
    if err := printer2(""); err != nil {
        
        //we can compare errors if we declare errors like errorEmptyString
        if err == errorEmptyString {
            fmt.Printf("You tried to print an empty string")
        }else {
            fmt.Printf("Printer failed %s\n",err)
        }
       
        os.Exit(1)
    }
}