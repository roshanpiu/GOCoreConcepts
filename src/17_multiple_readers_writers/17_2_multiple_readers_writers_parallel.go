package main 

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func getPage(url string) (int, error) {
     res, err := http.Get(url)
     if(err != nil){
         return 0, err
     }
     
     defer res.Body.Close() //this will execute when the function exits
     body, err := ioutil.ReadAll(res.Body)
     if err != nil {
         return 0, err
     }
     
     return len(body), nil
}

func getter(url string, size chan string)  {
    length, err := getPage(url)
    
    //tansmits the length through the size channel
    if err == nil {
        size <- fmt.Sprintf("%s has length %d", url, length) //Sprintf returns a formatted string rather than printing it
    }else{
        size <- fmt.Sprintf("Error getting %s: %s", url, err)
    }
}

func main() {
    urls := []string{"http://www.google.com","http://www.yahoo.com","http://www.facebook.com","http://www.twitter.com"}
    
    //every goroutine uses this same channel to transmit data
    size := make(chan string)
    
    //starts a goroutine for each url in urls
    
    for _, url := range urls{
        go getter(url, size)
    }
    
    for i:=0; i<len(urls); i++ {
        fmt.Printf("%s\n", <-size)
    }
    
    
}