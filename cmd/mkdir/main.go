package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	parents := flag.Bool("p", false, "create parent directories")
	flag.Parse()

	args:= flag.Args()
	if len(args)  == 0 {
		fmt.Fprintln(os.Stderr, "mkdir: missing operand")
		os.Exit(1)
	}

	for _, path := range args {
		err := Mkdir(path, *parents)
		if err != nil {
			fmt.Fprintf(os.Stderr, "mkdir: %v\n", err)
			os.Exit(1)
		}
	}	
	
}
