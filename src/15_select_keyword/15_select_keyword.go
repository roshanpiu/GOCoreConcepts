//the select keyword can be used to coordinate between channels

package main 

import (
    "fmt"
    "time"
)

func emit(wordChannel chan string, done chan bool)  {
    defer close(wordChannel) // before exiting this function the wordChannel channel will be closed
    
    words := []string{"The", "quick", "brown", "fox"}
    
    i := 0
    t := time.NewTimer(3* time.Second)
    
    //this will loop around forever doing 2 things
    //if someone is listening to the channel they will get the next word from words
    //if some one recive a message and done it will stop the transmission
    for {
        select {
            case wordChannel <- words[i]:
                i++
                if i == len(words){
                    i = 0
                }  
            case <-done:
                fmt.Printf("got done!\n")
                done <- true //sends true from the done channel
                close(done)
                return
            
            case <- t.C:
                return
            
        }
    }
}

func main() {
    wordCh := make(chan string)
    doneCh := make(chan bool)
    
    go emit(wordCh, doneCh)
    
    for i := 0; i<10; i++ {
        fmt.Printf("%s ", <-wordCh)
    }
    
    for word := range wordCh{
        fmt.Printf("%s ", word)
    }
    
    doneCh <- true
    <- doneCh //waits for true to come. this program will not terminate untill it recive something from the done channel
}
