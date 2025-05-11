package main

import (
	"fmt"
	"io"
	"strings"
)

func Echo(w io.Writer, args []string, noNewline bool) error {
	out := strings.Join(args, " ")
	if noNewline {
		_, err := fmt.Fprint(w, out)
		return err
	}
	_, err := fmt.Fprintln(w, out)
	return err
}
