package main

import (
	"flag"
	"fmt"
	"os"
)

func main(){
	recursive:= flag.Bool("r", false, "remove directories and their contents recursively")
	force := flag.Bool("f", false, "ignore nonexistent files and arguments, never prompt")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "rm: missing operand")
		os.Exit(1)
	}

	for _, path := range args {
		err := Rm(path, *recursive, *force)
		if err != nil{
			fmt.Fprintf(os.Stderr, "rm: %v\n", err)
			if !*force {
				os.Exit(1)
			}
		}
	}
}