package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	noNewline := flag.Bool("n", false, "do not print the trailing newline")
	flag.Parse()

	args := flag.Args()

	err := Echo(os.Stdout, args, *noNewline)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
