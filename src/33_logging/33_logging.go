package main 

import (
    "log"
)

func main() {
    log.SetFlags(log.Lmicroseconds)
    log.Fatalf("Failed to run the program")
}