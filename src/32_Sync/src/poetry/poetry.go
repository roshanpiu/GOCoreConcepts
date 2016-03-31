package poetry

import (
    "bufio"
    "os"
    "strings"
)

type Line string
type Stanza []Line
type Poem []Stanza

func NewPoem() Poem {
    return Poem{}
}

//new code for this chapter LoadPoem()
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

func (s Stanza) Len() int {
    return len(s)
}

func (s Stanza) Swap(i, j int)  {
    s[i], s[j] = s[j], s[i]
}

func (s Stanza) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
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

func (p Poem) NumWords() int {
    count := 0 
    for _, s := range p {
        for _, l := range s {
            sl := string(l)
            parts := strings.Split(sl, " ")
            count += len(parts)
        }
    }
    return count
}

func (p Poem) NumThe() int {
    count := 0 
    for _, s := range p {
        for _, l := range s {
            sl := string(l)
            if strings.Contains(sl, "The"){
                count++
            }
        }
    }
    return count
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