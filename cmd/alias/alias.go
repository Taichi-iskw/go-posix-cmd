package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

var aliasFile string = filepath.Join(userHomeDir(), ".go_aliases.json")

func userHomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return "."
	}
	return home
}

func AddAlias(name, command string) error {
	aliases := make(map[string]string)

	// read alias file
	data, err := os.ReadFile(aliasFile)
	if err == nil {
		_ = json.Unmarshal(data, &aliases)
	}

	// set alias
	aliases[name] = command

	// write alias file
	newData, err := json.MarshalIndent(aliases, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(aliasFile, newData, 0644)
}

func ListAliases() error {
	data, err := os.ReadFile(aliasFile)
	if os.IsNotExist(err){
		fmt.Println("No aliases found")	
		return nil
	}
	if err != nil{
		return err
	}
	var aliases map[string]string
	if err := json.Unmarshal(data, &aliases); err != nil{
		return err
	}
	for alias, command := range aliases{
		fmt.Printf("%s -> %s\n", alias, command)
	}
	return nil
}