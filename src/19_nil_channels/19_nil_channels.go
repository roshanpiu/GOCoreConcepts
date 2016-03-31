package main 

import (
    "fmt"
    "time"
    "math/rand"
)

//nil channels can be used to turn on and off channels 

func reader(ch chan int)  {
    
    t := time.NewTimer(10*time.Second)
    
    for {
        select {
            case i := <- ch: //if something comes from the the channel this case will print it otherwise carry on
                fmt.Printf("%d\n", i)
                
            case <-t.C  :
                ch = nil //when this case heppens after 3 seconds the goroutine will stop and wait
        }
    }
}

func writer(ch chan int)  {
    
    stopper := time.NewTimer(2* time.Second)
    restarter := time.NewTimer(5* time.Second)
    
    savedCh := ch //save the channel before it gets set to nil
    
    for {
        select {
            case ch <- rand.Intn(42):
            case <- stopper.C: //after 2 seconds the writer will stop and wait 
                ch = nil    
            case <- restarter.C: //after 5 seconds the writer start and transmit random numbers  
                ch = savedCh
        }
        
    }
}

func main() {
    
    ch := make(chan int)
    
    go reader(ch)
    go writer(ch)
    
    time.Sleep(20 * time.Second)
    
}