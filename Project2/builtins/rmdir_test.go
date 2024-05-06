package builtins_test

import (
	"os"
	"testing"

	"github.com/Sachneu/4600Project2/Project2/builtins"
)

func TestRmdir(t *testing.T) {
	testDir := "testdir"
	err := os.Mkdir(testDir, 0755)
	if err != nil {
		t.Fatalf("Error creating test directory: %v", err)
	}

	err = builtins.RemoveDirectory(testDir)
	if err != nil {
		t.Fatalf("Error removing directory: %v", err)
	}

	_, err = os.Stat(testDir)
	if !os.IsNotExist(err) {
		t.Errorf("Directory %s still exists", testDir)
	}
}
