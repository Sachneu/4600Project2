package builtins_test

import (
	"os"
	"testing"

	"github.com/Sachneu/4600Project/Project2/builtins"
)

func TestTouch(t *testing.T) {
	testFile := "testfile.txt"

	err := builtins.Touch(testFile)
	if err != nil {
		t.Fatalf("Error creating file: %v", err)
	}

	defer func() {
		err := os.Remove(testFile)
		if err != nil {
			t.Fatalf("Error cleaning up file: %v", err)
		}
	}()

	_, err = os.Stat(testFile)
	if err != nil {
		t.Errorf("File %s not created", testFile)
	}
}
