package internals

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	db := Init()

	assert.Equal(t, byte(0), db.root.b)
	assert.Equal(t, false, db.root.exists)
}

func TestDB_Find(t *testing.T) {
	setdb := Init()

	tests := []struct{
		input []byte
		output []byte
		expected bool
	} {
		{[]byte("1234567"),[]byte("1234567"), true},
		{[]byte("1234567"),[]byte("1234568"), false},
	}

	for _, test := range tests {
		setdb.Put(test.input)
		assert.Equal(t, test.expected, setdb.Find(test.output))
	}
}

func BenchmarkDB_PutString_Memory(b *testing.B) {
	setdb := Init()

	loops := 1000000

	b.StartTimer()
	for i := 0; i < loops; i++ {
		uuids := uuid.New()
		setdb.PutString(uuids.String())
	}
	b.StopTimer()
	MemUsage()
}

func BenchmarkDB_PutString(b *testing.B) {
	setdb := Init()

	uuids := uuid.New()

	b.StartTimer()
	setdb.PutString(uuids.String())
	b.StopTimer()
}
