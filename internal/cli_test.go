package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCLI_Parse(t *testing.T) {
	cli := InitCLI()

	tests := []struct{
		cmd string
		expected [2]string
	}{
		{"put 123", [2]string{"put", "123"}},
		{"find 123", [2]string{"find", "123"}},
		{"put hello world", [2]string{"put", "hello world"}},
		{"find hello world", [2]string{"find", "hello world"}},
	}

	for _, test := range tests {
		actual, _ := cli.parse(test.cmd)

		assert.Equal(t, test.expected[0], actual[0])
		assert.Equal(t, test.expected[1], actual[1])
	}
}

func TestCLI_Parse_Error(t *testing.T) {
	cli := InitCLI()

	tests := []struct{
		cmd string
		expected string
	} {
		{"puts 123", "Invalid command"},
	}

	for _, test := range tests {
		_, err := cli.parse(test.cmd)

		assert.EqualError(t, err, test.expected)
	}
}