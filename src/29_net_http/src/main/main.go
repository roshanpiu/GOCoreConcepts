package main

import (
   "fmt"
   "poetry"
   "net/http"
)

func poemHandler(w http.ResponseWriter, r *http.Request)  {
    r.ParseForm()
    poemName := r.Form["name"][0]
    
    p, err := poetry.LoadPoem(poemName)
    
    if err != nil {
        http.Error(w, "File Not Found", http.StatusNotFound)
    }else {
        fmt.Fprintf(w, "%s\n",p)
    }
    
   
}

func main() {
    
    http.HandleFunc("/poem",poemHandler)
    http.ListenAndServe(":8088",nil)
}