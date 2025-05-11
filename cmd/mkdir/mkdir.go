package main

import (
	"os"
)

func Mkdir(path string, parents bool) error {
	if parents {
		return os.MkdirAll(path, 0755)
	}
	return os.Mkdir(path, 0755)
}
