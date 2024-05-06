package builtins

import (
	"fmt"
	"io"
)

func Echo(w io.Writer, args ...string) error {
	var printArgs []interface{}
	for _, arg := range args {
		printArgs = append(printArgs, arg)
	}
	_, err := fmt.Fprintln(w, printArgs...)
	return err
}
