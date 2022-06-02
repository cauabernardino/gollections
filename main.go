package main

import (
	"fmt"
	"os"

	"github.com/cauabernardino/gollections/algorithms"
	"github.com/cauabernardino/gollections/structures"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("wrong usage")
		os.Exit(1)
	}

	arg := os.Args[1]
	if arg == "algo" {
		algorithms.Run()
	} else if arg == "ds" {
		structures.Run()
	}
}
