package builtins_test

import (
	"bytes"
	"errors"
	"github.com/Sachneu/4600Project2/builtins"
	"testing"
)

func TestEcho(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "echo single string",
			args:     []string{"Hello, World!"},
			expected: "Hello, World!\n",
		},
		{
			name:     "echo multiple strings",
			args:     []string{"The", "quick", "brown", "fox"},
			expected: "The quick brown fox\n",
		},
		{
			name:     "echo empty string",
			args:     []string{""},
			expected: "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := builtins.Echo(&buf, tt.args...)
			if err != nil {
				t.Errorf("Echo() error = %v, want nil", err)
				return
			}
			if buf.String() != tt.expected {
				t.Errorf("Echo() = %v, want %v", buf.String(), tt.expected)
			}
		})
	}
}
