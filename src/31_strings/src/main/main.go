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
)

var c config

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
    found := false
    for _, v := range c.ValidPoems {
        if v == poemName {
            found = true
            break
        }
    }
    
    
    if !found {
         http.Error(w, "Not Found (invalid)", http.StatusNotFound)
         //fmt.Println(c.BindAddress)
         return
    }
    
    p, err := poetry.LoadPoem(poemName)
    
    if err != nil {
        http.Error(w, "File Not Found", http.StatusNotFound)
        fmt.Println(err)
    }else {
        
        //sorts the poem (inorder to pass a Stanza to Sort() it has to implement Len, Swap, Less functions)
        sort.Sort(p[0])
        
        //we can pass the stanza to Shuffle because it now implements the Shuffleable interface 
        shuffler.Shuffle(p[0]) 
        
        
        pwt := poemWithTitle{poemName, p, poemName, strconv.FormatInt(int64(p.NumWords()), 16), p.NumThe()}
        enc := json.NewEncoder(w)
        
        //enc.Encode(p) //converts a poem to json
        //converts a struct to json (only members of the struct begins with uppercase character will get output to json)
        enc.Encode(pwt) 
    }
    
    
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
    
    if err != nil {
        os.Exit(1)
    }
    
    http.HandleFunc(c.Route,poemHandler)
    http.ListenAndServe(c.BindAddress,nil)
}