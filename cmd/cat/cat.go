package main

import "io"

func Cat(r io.Reader, w io.Writer) error {
	_, err := io.Copy(w, r)
	return err
}
