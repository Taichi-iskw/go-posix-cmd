package main

import "os"

func Rm(path string, recursive bool, force bool) error {
	if recursive {
		return os.RemoveAll(path)
	}

	err := os.Remove(path)
	if force && os.IsNotExist(err) {
		return nil
	}

	return nil
}