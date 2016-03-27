package main 

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

type webPage struct {
    url string
    body []byte
    err error
}

//(w *webPage) is the reciever
//this function will get attached to the webPage type
func (w *webPage) get()  {
    res, err := http.Get(w.url)
    if err != nil {
        w.err = err
        return
    }
    defer res.Body.Close()
    
    resBody, err := ioutil.ReadAll(res.Body)
    w.body = resBody
    if err != nil {
        w.err = err
    }
}

//this function will get attached to the webPage type
func (w *webPage) isOK() bool {
    return w.err == nil 
}

func main() {
    //this creates a pointer
    //these are the 3 ways of declaring a value from a struct
    w := &webPage{url: "http://www.google.com"}
    
    w1 := &webPage{} //this will initialize a empty struct every value inside will get initialize by it's 0 value
    w1.url = "http://www.yahoo.com"
    
    w2 := new(webPage)
    w2.url = "http://www.gmail.com"
    
    w.get()
    if(w.isOK()){
        fmt.Printf("URL: %s Error: %s Length: %d ",w.url ,w.err, len(w.body))
    }else {
        fmt.Printf("Something went wrong")
    }
    
}