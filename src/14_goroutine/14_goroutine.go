//goroutine is a seperately parallely running function from the main program 
//channels are go's fundamental communication machanism
//passing messages between goroutines are done using channels
package main 

import (
    "fmt"
    "math/rand"
)

func emit(c chan string)  {
    //c channel passes strings
    words := []string{"the", "quick", "brown", "fox"}
    for _, word := range words {
        //transmits word through c
        c <- word
    }
    close(c) //closes channel (this is important because if we dont close channels all the channels will get dead locked)
}

func makeRandoms(c chan int)  {
    //c channel passes ints
    for {
        c <- rand.Intn(1000)
    }
    
}

func makeId(idChan chan int)  {
    var id int
    id = 0
    
    for {
        idChan <- id
        id += 1
    }
}

func main() {
    wordChannel := make(chan string) //mekes a channel capable of passing strings
    randoms := make(chan int)
    idChannel := make(chan int)
    
    //the go keyword starts a goroutine which runs concurrently with main function
    go emit(wordChannel)
    go makeRandoms(randoms)
    go makeId(idChannel)
    
    for word := range wordChannel {
        fmt.Printf("%s ", word)
    }
fmt.Printf("\n")
    
    //this will be a infinite loop
    // for n := range randoms {
    //     fmt.Printf("%d ", n)
    // }
    
    //gets a single word from the channel
    word := <- wordChannel
    fmt.Printf("%s \n", word)
    
    //gets a single word from the channel
    randomNumber := <- randoms
    fmt.Printf("Random Number: %d \n", randomNumber)
    
    fmt.Printf("Your new ID: %d \n", <- idChannel)
    fmt.Printf("Your new ID: %d \n", <- idChannel)
    fmt.Printf("Your new ID: %d \n", <- idChannel)
    fmt.Printf("Your new ID: %d \n", <- idChannel)
    
    fmt.Printf("\n")
}