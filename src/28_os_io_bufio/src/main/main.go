package main

import (
    "fmt"
    "poetry"
)

func main() {
    //p := poetry.NewPoem() //can initialize like this
    //p := poetry.Poem{{"Hello","Helloooo","Hellooooooooo","Hello World","Hello From Golang"}}
    
    p, err := poetry.LoadPoem("theprelude") 
    
    if err != nil {
        fmt.Printf("Error loading poem %s\n",err)
    }
    
    fmt.Printf("%s\n",p)
    
    
    v, c, puncs := p.Stats()
    fmt.Printf("Vovels:%d Consonants:%d Puncs:%d \n",v,c, puncs)
    fmt.Printf("Stanzas:%d Lines:%d \n",p.NumStanzas(),p.NumLines())
}