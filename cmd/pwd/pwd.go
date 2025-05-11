package main

import "os"

func GetWorkingDir() (string, error) {
	return os.Getwd()
}
