package poetry

type Line string
type Stanza []Line
type Poem []Stanza

func NewPoem() Poem {
    return Poem{}
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