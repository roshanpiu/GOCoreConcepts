package main 

import (
    "fmt"
    "shuffler"
)
/////////////////////////////////////////////
type intSlice []int 

func (is intSlice) Len() int  {
    return len(is)
}

func (is intSlice) Swap(i, j int) {
    is[i], is[j] = is[j], is[i]
}

/////////////////////////////////////////////

type weightedString struct{
    weight int
    s string
}

type stringSlice []weightedString

func (ss stringSlice) Len() int  {
    return len(ss)
}

func (ss stringSlice) Swap(i, j int) {
    ss[i], ss[j] = ss[j], ss[i]
}

func (ss stringSlice) Weight(i int) int{
    return ss[i].weight
}

func main() {
    fmt.Printf("Hello World!\n")
    i := intSlice{1,2,3,4}
    fmt.Println(i)
    shuffler.Shuffle(i)
    fmt.Println(i)
    fmt.Println("Loop Count: ",shuffler.GetCount())
    shuffler.Shuffle(i)
    shuffler.Shuffle(i)
    fmt.Println(i)
    fmt.Println("Loop Count: ",shuffler.GetCount())
    
    // s:= stringSlice{weightedString{100, "Hello"},weightedString{200,"World"},weightedString{300,"GoLang"},weightedString{10,"Good bye!"}}
    // fmt.Println(s)
    // shuffler.WeightedShuffle(s)
    // fmt.Println(s)
}