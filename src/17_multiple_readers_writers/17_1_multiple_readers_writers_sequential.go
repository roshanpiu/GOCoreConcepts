package main 

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
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

func main() {
    urls := []string{"http://www.google.com","http://www.yahoo.com","http://www.facebook.com","http://www.twitter.com"}
    
    for _, url := range urls{
        
        pageLength, err := getPage(url)
    
        if err != nil{
            os.Exit(1)
        }
        
        fmt.Printf("%s is length %d\n",url,pageLength)
        
    }
    
    
}