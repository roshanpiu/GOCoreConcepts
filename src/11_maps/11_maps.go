package main 

import (
    "fmt"
)

func main() {
    dayMonths := make(map[string]int)
    dayMonths["Jan"] = 31
    dayMonths["Feb"] = 28
    dayMonths["Mar"] = 31
    
    fmt.Printf("Days in February: %d\n",dayMonths["Feb"])
    //if we access non existing element in a map we dont get an error (in this case gets a 0 because value type is 0)
    
    //check if the values exist or not
    days, ok := dayMonths["January"]
    if !ok{
        fmt.Println("Can't get the days for January")
        
    }else {
        fmt.Println(days)
    }
    
    //looping over a map
    for month, days := range dayMonths {
        fmt.Printf("%s has %d days\n",month,days)
    }
    
    //deleting items in a map
    fmt.Println(dayMonths)
    delete(dayMonths, "Feb")
    fmt.Println(dayMonths)
    
    //declaring maps
    worldCups := map[string]int{
        "Sri lanka": 1,
        "Pakisten": 1,
        "India": 2,
        "Bangladesh": 0,
    }
    
    fmt.Println(worldCups)
}