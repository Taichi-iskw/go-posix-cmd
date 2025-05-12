package main

import (
	"os"
	"time"
)

func Touch(path string) error {
	now := time.Now()
	
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f.Close()
	}

	return os.Chtimes(path, now, now)
}