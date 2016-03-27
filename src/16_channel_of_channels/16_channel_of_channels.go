//the select keyword can be used to coordinate between channels

package main 

import (
    "fmt"
    "time"
)

func emit(chanChannel chan chan string, done chan bool)  {
    wordChannel := make(chan string)
    chanChannel <- wordChannel
    defer close(wordChannel)
    
    words := []string{"The", "quick", "brown", "fox"}
    
    i := 0
    t := time.NewTimer(3* time.Second)

    for {
        select {
            case wordChannel <- words[i]:
                i++
                if i == len(words){
                    i = 0
                }  
            case <-done:
                fmt.Printf("got done!\n")
                done <- true 
                close(done)
                return
            
            case <- t.C:
                return
            
        }
    }
}

func main() {
    channelCh := make(chan chan string)
    doneCh := make(chan bool)
    
    go emit(channelCh, doneCh) 
    wordCh := <-channelCh
    
    for word := range wordCh{
        fmt.Printf("%s ", word)
    }

}
