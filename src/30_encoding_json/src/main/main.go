package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"poetry"
)

var c config

type poemWithTitle struct {
	Title  string
	Body   poetry.Poem
	title2 string //this will not get encoded to json because its a private member
}

type config struct {
	Route       string
	BindAddress string   `json:"addr"` //this is a hint to json decoder that says in json this BindAddress will be called addr
	ValidPoems  []string `json:"Valid"`
}

func poemHandler(w http.ResponseWriter, r *http.Request) {
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
	} else {
		pwt := poemWithTitle{poemName, p, poemName}
		enc := json.NewEncoder(w)
		//enc.Encode(p) //converts a poem to json
		enc.Encode(pwt) //converts a struct to json (only members of the struct begins with uppercase character will get output to json)
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

	http.HandleFunc(c.Route, poemHandler)
	http.ListenAndServe(c.BindAddress, nil)
}
