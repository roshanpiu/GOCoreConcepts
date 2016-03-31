package main 

import (
    "time"
    "fmt"
)

func main() {
    time1 := time.Now()
    fmt.Println(time1)
    time.Sleep(3 * time.Second)
    time2 := time.Now()
    elapsed := time.Since(time1) //gets the time difference between time1 and now
    elapsed1 := time.Now().Sub(time1)
    fmt.Println("Difference ", elapsed)
    fmt.Println("Difference ", elapsed1)
    fmt.Println(time2)
    fmt.Println("Ended: ",time2.Format(time.Kitchen)) // formats time
}