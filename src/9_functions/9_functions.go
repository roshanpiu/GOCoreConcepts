package main 

import (
    "fmt"
    "os"
)

func printer1(msg1 string, msg2 string)  {
    fmt.Printf("%s %s\n", msg1, msg2)
}

//specify type for both paramaters

func printer2(msg1, msg2 string, repeat int) {
    for i:=1; i<= repeat; i++{
        fmt.Printf("%s %s\n", msg1, msg2)
    } 
}

//this function returns an error type
func printer3(msg string) error {
    _,err := fmt.Printf("%s \n", msg)
    return err
}

//returning multiple values
func printer4(msg string) (string, error) {
    msg += "\n"
    _,err := fmt.Printf(msg)
    return msg, err
}

//using defer
//defer can be used to stackup the things to do when we exit a function
func printer5(msg string) error {
    
    defer fmt.Printf("\n") //when you exit this function print a line feed
    defer fmt.Printf(" the end\n") //the end prints first and newline prints second
    _,err := fmt.Printf("%s", msg)
    return err
}

//using named return values
func printer6(msg string) (err error) {
    f, err := os.Create("hello.txt")
    if err != nil {
        return err
    }
    defer f.Close() //on exit this function the file handle will be closed
    
    //f.Write(byte(msg))
    
    //this return will return err variable of the os.Create
    //this happens because of the named return values
    return
}

//passing unknows number of arguments to a function
//we cannot have multiple ... for a function only one is allowed
func printer7(format string, msgs...string)  {
    for _, msg := range msgs{
        fmt.Printf(format,msg)
    }
}

func main() {
    printer1("Hello","World")
    printer2("Hello","World" , 3)
    printer3("Hello From Printer 3")
    
    appendedMsg, err := printer4("Hello World")
    if(err == nil){
        fmt.Printf("%q\n",appendedMsg) // %q prints friendly strings
        fmt.Printf("%x\n",appendedMsg) // %x prints hex version of strings
    }
    
    printer5("Hello World")
    printer7("%s\n","H","E","L","L","O")
    printer7("%x\n","H","E","L","L","O")
}