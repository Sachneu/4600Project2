package builtins

import "os"

func MakeDirectory(args ...string) error {
	for _, dir := range args {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
