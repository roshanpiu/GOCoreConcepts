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

func worker(urlCh chan string, size chan string, id int)  {
    for {
        url := <- urlCh
        length, err := getPage(url)
        //tansmits the length and error through the size channel
        if err == nil {
            size <- fmt.Sprintf("%s has length %d (%d)", url, length, id)
        }else {
            size <- fmt.Sprintf("Error getting %s: %s", url, err)
        }
    }
    
    
}

func generator(url string, urlCh chan string)  {
    urlCh <- url
}

func main() {
    
    urls := []string{"http://www.google.com","http://www.yahoo.com","http://www.facebook.com","http://www.twitter.com"}
    
    urlCh := make(chan string)
    sizeCh := make(chan string)
    
    for i := 0; i < 10; i++ {
        go worker(urlCh, sizeCh, i)
    }
    
    // 4 goroutines will be created to transmit the url to channel
    for _, url := range urls {
        go generator(url, urlCh)
    }
    
    for i:=0; i<len(urls); i++ {
        fmt.Printf("%s\n", <-sizeCh)
    }
    
}