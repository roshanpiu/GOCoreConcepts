package main

import (
   "poetry"
   "net/http"
   "encoding/json"
   "os"
   "fmt"
   "strconv"
   "sort"
   "shuffler"
   "sync"
)

// any variable created by this type will get the methods of sync.Mutex
type protectedCache struct {
    sync.Mutex
    c map[string]poetry.Poem
}

// var cache protectedCache
// cache.Lock()
// cache.Unlock()

var cacheMutex sync.Mutex
var c config

var cache map[string]poetry.Poem

type poemWithTitle struct {
    Title string
    Body poetry.Poem
    title2 string //this will not get encoded to json because its a private member
    WordCount string 
    TheCount int
}

type config struct {
    Route string
    BindAddress string `json:"addr"` //this is a hint to json decoder that says in json this BindAddress will be called addr
    ValidPoems []string  `json:"Valid"`
}

func poemHandler(w http.ResponseWriter, r *http.Request)  {
    r.ParseForm()
    
    poemName := r.Form["name"][0]
    
    //lets users to get only valid poems
    p, ok := cache[poemName]
    
    
    if !ok {
         http.Error(w, "Not Found (invalid)", http.StatusNotFound)
         return
    }
    
    pwt := poemWithTitle{poemName, p, poemName, strconv.FormatInt(int64(p.NumWords()), 16), p.NumThe()}
    enc := json.NewEncoder(w)
    enc.Encode(pwt) 
    
}

func main() {
    
    //decoding json into structs
    f, err := os.Open("config")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    
    dec := json.NewDecoder(f)
    err = dec.Decode(&c) //passes the reference of c so the Decode could modify it
    f.Close()
    if err != nil {
        os.Exit(1)
    }
    
    cache = make(map[string]poetry.Poem) //we have to initialize a map before using it because maps does not have a 0 value
    
    var wg sync.WaitGroup
    
    for _, name := range c.ValidPoems{
        wg.Add(1)
        go func (n string)  {
            cacheMutex.Lock()
            defer cacheMutex.Unlock()
            
            cache[name], err = poetry.LoadPoem(name)
            if err != nil {
                os.Exit(1)
            }
            wg.Done()
        }(name)
    }
    
    wg.Wait() // this will wait untill the number of the wait groups are 0 and all the goroutines are done 
    
    http.HandleFunc(c.Route,poemHandler)
    http.ListenAndServe(c.BindAddress,nil)
}