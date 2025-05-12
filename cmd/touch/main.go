package main

import (
	"fmt"
	"os"
)

func main(){
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "touch: missing operand")
		os.Exit(1)
	}

	for _, path := range args {
		err := Touch(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "touch: %v\n", err)
			os.Exit(1)
		}
	}
}