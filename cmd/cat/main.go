package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		if err := Cat(os.Stdin, os.Stdout); err != nil {
			fmt.Fprintf(os.Stderr, "cat: %v\n", err)
			os.Exit(1)
		}
		return
	}

	for _, fname:= range args{
		f, err := os.Open(fname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cat: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		
		if err := Cat(f, os.Stdout); err != nil {
			fmt.Fprintf(os.Stderr, "cat: %v\n", err)
			os.Exit(1)
		}
	}
}