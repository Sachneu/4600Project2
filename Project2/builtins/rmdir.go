package builtins

import "os"

func RemoveDirectory(args ...string) error {
	for _, dir := range args {
		err := os.Remove(dir)
		if err != nil {
			return err
		}
	}
	return nil
}
