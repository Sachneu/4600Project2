package builtins_test

import (
	"os"
	"testing"

	"github.com/Sachneu/4600Project2/Project2/builtins"
)

func TestPwd(t *testing.T) {
	expected, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error getting current directory: %v", err)
	}

	var bufPwd string
	err = builtins.PrintWorkingDirectory(&bufPwd)
	if err != nil {
		t.Fatalf("Error executing pwd command: %v", err)
	}

	if bufPwd != expected+"\n" {
		t.Errorf("PrintWorkingDirectory() returned %s, expected %s", bufPwd, expected)
	}
}
