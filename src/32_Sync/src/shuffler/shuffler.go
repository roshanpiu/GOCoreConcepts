package shuffler

import (
    "math/rand"
)

//global variable like this will be setup when package is loaded 
var counter int 

//when initializing the packege any work needs to be done can be done in init function
func init()  {
    counter = 0
}

func GetCount() int {
    return counter
}

type Shuffleable interface {
    Len() int
    Swap(i, j int)
}

type WeightShuffleable interface{
    Shuffleable
    Weight(i int) int
}

func Shuffle(s Shuffleable)  {
    //shuffles
    for i:=0; i<s.Len(); i++{
        j:= rand.Intn(s.Len()-i)
        s.Swap(i,j)
        counter++
    }
}

func WeightedShuffle(w WeightShuffleable)  {
    total := 0
    for i:=0; i<w.Len(); i++{
        total += w.Weight(i)
    }
    
    for i:=0; i<w.Len(); i++{
        pos := rand.Intn(total)
        cum := 0
        for j:=i; j<w.Len(); i++{
            cum += w.Weight(j)
            if pos >= cum {
                total -= w.Weight(j)
                w.Swap(i, j)
                break
            }
        }
    }
}