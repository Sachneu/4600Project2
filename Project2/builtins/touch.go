package builtins

import "os"

func Touch(args ...string) error {
	for _, file := range args {
		f, err := os.Create(file)
		if err != nil {
			return err
		}
		f.Close()
	}
	return nil
}
