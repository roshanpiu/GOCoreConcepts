package poetry

import (
    "bufio"
    "os"
)

type Line string
type Stanza []Line
type Poem []Stanza

func NewPoem() Poem {
    return Poem{}
}

//new code for this chaper LoadPoem()
//use ioutil to read large chunks of data using high level abatractions 

func LoadPoem(name string) (Poem, error)  {
    f, err := os.Open(name)
    if err != nil {
        return Poem{}, err
    }
    defer f.Close()
    
    p := Poem{}
    
    var s Stanza
    
    scan := bufio.NewScanner(f)
    for scan.Scan() {
        l := scan.Text()
        if l == "" {
            p = append(p, s)
            s = Stanza{}
            continue
        }
        
        s = append(s, Line(l))
    }
    p = append(p, s)
    
    if scan.Err() != nil {
        return nil, scan.Err()
    }
    
    return p, nil
}

func (p Poem) NumStanzas() int {
    return  len(p)
}

func (s Stanza) NumLines() int {
    return len(s)
}

func (p Poem) NumLines() (count int) {
    for _, s := range p {
        count += s.NumLines()
    }
    return
}

func (p Poem) Stats() (numVowels, numConsonants, numPuncs int) {
    for _, s := range p {
        for _, l := range s {
            for _, r := range l {
                switch r {
                    case 'a','e','i','o','u':
                        numVowels++
                    case ',','!',' ':
                        numPuncs++
                    default: 
                        numConsonants++
                        
                }
            }
            
        }
    }
    
    return
}