package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestByteUnit_Find(t *testing.T) {
	initialBu := InitByteUnit(byte('a'))
	nextBu := InitByteUnit(byte('b'))

	initialBu.next[byte('b')] = nextBu

	tests := []struct{
		Input byte
		Expected *ByteUnit
	} {
		{byte('b'), nextBu},
		{byte('a'), nil},
	}

	for _, test := range tests {
		actual := initialBu.Find(test.Input)
		assert.Equal(t, test.Expected, actual)
	}
}

func TestInitByteUnit(t *testing.T) {
	tests := []struct{
		Byte byte
	}{
		{byte('a')},
		{byte(10)},
		{byte(0xA)},
	}

	for _, test := range tests {
		byteUnit := InitByteUnit(test.Byte)
		assert.Equal(t, test.Byte, byteUnit.b)
	}
}

func TestInitAndSetByteUnit(t *testing.T) {
	tests := []struct{
		Byte byte
		Expected bool
	}{
		{byte('a'), true},
		{byte(10), true},
		{byte(0xA), true},
	}

	for _, test := range tests {
		byteUnit := InitAndSetByteUnit(test.Byte)
		assert.Equal(t, test.Expected, byteUnit.exists)
	}
}