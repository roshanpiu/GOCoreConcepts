package poetry

import (
    "testing"
)
//t.ErrorF() will print the error and keep running the test suite
//t.Fatalf() will print the error and terminate the test suite immediately

//in go testing is firstclass it is part of the language it self

func TestNumStanzas(t *testing.T)  {
    
    p := Poem{}
    
    if p.NumStanzas() != 0 {
        t.Fatalf("Empty poem is not empty")
    }
    
    p = Poem{{"Hello","Helloooo","Hellooooooooo","Hello World","Hello From Golang"}}
    
    if p.NumStanzas() != 1 {
        t.Fatalf("Unexpected stanza count %d", p.NumStanzas())
    } 
    
}

func TestNumLines(t *testing.T)  {
    
    p := Poem{}
    
    if p.NumStanzas() != 0 {
        t.Fatalf("Empty poem is not empty")
    }
    
    p = Poem{{"Hello","Helloooo","Hellooooooooo","Hello World","Hello From Golang"}}
    
    if p.NumLines() != 5 {
        t.Fatalf("Unexpected line count %d", p.NumLines())
    } 
    
}

func TestStats(t *testing.T)  {
    p := Poem{}
    
    v, c, puncs := p.Stats()
    if v != 0 || c != 0 || puncs != 0 {
        t.Fatalf("Bad number of vowels or consonants or puncs")
    }
    
    p = Poem{{"Hello, World!"}}
    v, c, puncs = p.Stats()
    if v != 3 || c != 7 || puncs != 3 {
        t.Fatalf("Bad number of vowels or consonants (%d %d %d)", v, c, puncs)
    }
}