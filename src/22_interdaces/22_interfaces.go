//interfaces define behavour rather than a specific type

package main 

import (
    "fmt"
    "math/rand"
)

type shuffler interface {
    Len() int
    Swap(i, j int)
}

//anything which satisfies the shuffler interface can be passed to this function
//in this case intSlice and stringSlice satisfies the shuffler interface
func shuffle(s shuffler)  {
    for i:=0; i<s.Len(); i++{
        j := rand.Intn(s.Len()-i)
        s.Swap(i, j)
    }
}

///////////////////////////////////////
type intSlice []int

func (is intSlice) Len() int{
    return len(is)
}

func (is intSlice) Swap(i, j int)  {
    is[i], is[j] = is[j], is[i]
}
/////////////////////////////////////
type stringSlice []string

func (ss stringSlice) Len() int{
    return len(ss)
}

func (ss stringSlice) Swap(i, j int)  {
    ss[i], ss[j] = ss[j], ss[i]
}
/////////////////////////////////////


func main() {
    is := intSlice{1,2,3,4,5}
    shuffle(is)
    fmt.Println(is)
    
    ss := stringSlice{"a","b","c","d","e"}
    shuffle(ss)
    fmt.Println(ss)
}