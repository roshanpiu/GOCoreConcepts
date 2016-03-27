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

//building aliases for built in types and adding additional methods to them
type SummableSlice  []int

func (s SummableSlice) sum() int {
    sum := 0
    for _,i := range s{
        sum += i
    }
    return sum
}

//(w *webPage) is the reciever
//this function will get attached to the webPage type
//when we dont use pass by reference to pass the values to a function the changes done to the values are not saved 
//because a copy of the value is passed to the function
//if a slice is passed then we dont have to use pointers because slices are passed by reference
func (w webPage) get()  {
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
func (w webPage) isOK() bool {
    return w.err == nil 
}

func main() {
    w := webPage{url: "http://www.google.com"}

    w.get()
    if(w.isOK()){
        fmt.Printf("URL: %s Error: %s Length: %d ",w.url ,w.err, len(w.body))
    }else {
        fmt.Printf("Something went wrong")
    }
    
    s := SummableSlice{1,2,3,4,5}
    
    fmt.Printf("Sum :%d\n",s.sum())
    
}