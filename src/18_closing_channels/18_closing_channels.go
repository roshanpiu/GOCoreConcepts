package main 

import "fmt"
import "time"

func printer1(msg string, goCh chan bool)  {
    <-goCh // goCh waits to receive something.this go routine will start as soon as it receive something on the goCh
    fmt.Printf("%s\n", msg)
}

func printer2(msg string, stopCh chan bool)  {
   for {
       select {
            case <- stopCh: //this go routine will stop as soon as it receive something on the stopCh
                return
            default:
                fmt.Printf("%s\n",msg)
       }
   }
}

//use close for coordinating work 
//use close when multiple things needs to start at the same time and multiple things needs to stop at the same time
//close() can be used as a broadcasting machanism

func main() {
    
    //goCh := make(chan bool)

    // for i:=0; i<10; i++ { 
    //     go printer1(fmt.Sprintf("printer: %d", i), goCh)
    // }
    
    // time.Sleep(5 * time.Second)
    // close(goCh)
    // time.Sleep(2 * time.Second)
    
    stopCh := make(chan bool)
    
    for i:=0; i<10; i++ { 
        go printer2(fmt.Sprintf("printer: %d", i), stopCh)
    }
    
    //all the goroutines will stop after we close the stop channel
    time.Sleep(5 * time.Second)
    close(stopCh)
    time.Sleep(2 * time.Second)
}