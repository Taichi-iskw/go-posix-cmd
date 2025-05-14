package main

import (
	"fmt"
	"os"
)

func main(){
	if len(os.Args) < 2{
		fmt.Fprintln(os.Stderr, "Usage: alias <add/list> ")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 4{
			fmt.Fprintln(os.Stderr, "Usage: alias add <alias> <command>")
			os.Exit(1)
		}
		alias := os.Args[2]
		command := os.Args[3]
		if err := AddAlias(alias, command); err != nil{
			fmt.Fprintln(os.Stderr, "Error adding alias:", err)
			os.Exit(1)
		}
		fmt.Println("Alias added successfully")
	case "list":
		err := ListAliases()
		if err != nil{
			fmt.Fprintln(os.Stderr, "Error listing aliases:", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Usage: alias <command>")
		os.Exit(1)
	}

}