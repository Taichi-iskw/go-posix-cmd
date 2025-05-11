package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := GetWorkingDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(dir)
}