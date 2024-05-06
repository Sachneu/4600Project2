package builtins_test

import (
	"os"
	"testing"

	"github.com/Sachneu/4600Project2/Project2/builtins"
)

func TestMkdir(t *testing.T) {
	testDir := "testdir"

	err := builtins.MakeDirectory(testDir)
	if err != nil {
		t.Fatalf("Error creating directory: %v", err)
	}

	defer func() {
		err := os.Remove(testDir)
		if err != nil {
			t.Fatalf("Error cleaning up directory: %v", err)
		}
	}()

	_, err = os.Stat(testDir)
	if err != nil {
		t.Errorf("Directory %s not created", testDir)
	}
}
