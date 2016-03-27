//buffered channels allows you to send data when nobody is ready to recieve up to some capacity
//using unbuffered channels makes reasoning easy
package main 

import (
    "fmt"
    "math/rand"
    "time"
    "sync/atomic"
)

var running int64 = 0

func work()  {
    atomic.AddInt64(&running, 1)
    fmt.Printf("[%d", running)
    time.Sleep(time.Duration(rand.Intn(6)) * time.Second)
    atomic.AddInt64(&running, -1)
    fmt.Printf("]")
}

func worker(sema chan bool)  {
    <-sema
    work()
    sema <- true
}

func main() {
    //declaring buffered channels
    //this channel can hold upto 100 ints
    //once this channel gets full if anyone try to send on this channel they will have to wait to get space
    //when recieving anyone can recieve even when someone is not transmitting at the time
    //this can be used to control the concurrency 
    sema := make(chan bool, 10)
    
    for i:=0; i<1000; i++ {
        go worker(sema)
    }
    
    for i:=0; i<cap(sema); i++ {
        sema <- true
    }
    
    time.Sleep(20 * time.Second)
}