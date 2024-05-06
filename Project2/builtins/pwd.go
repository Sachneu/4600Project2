package builtins

import (
	"fmt"
	"io"
	"os"
)

func PrintWorkingDirectory(w io.Writer) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(w, wd)
	return err
}
