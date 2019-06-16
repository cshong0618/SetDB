package setdb

import (
	"github.com/google/uuid"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
	"log"
	"runtime"
	"testing"
)

func TestInit(t *testing.T) {
	db := Init()

	assert.Equal(t, byte(0), db.root.b)
	assert.Equal(t, false, db.root.exists)
}

func TestDB_Find(t *testing.T) {
	setdb := Init()

	tests := []struct {
		input    []byte
		output   []byte
		expected bool
	}{
		{[]byte("1234567"), []byte("1234567"), true},
		{[]byte("1234567"), []byte("1234568"), false},
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
		ids := xid.New()
		setdb.PutString(ids.String())
	}
	b.StopTimer()
	MemUsage()
	log.Println("Items in db: ", setdb.Items())
}

func BenchmarkDB_PutString(b *testing.B) {
	setdb := Init()

	uuids := uuid.New()

	b.StartTimer()
	setdb.PutString(uuids.String())
	b.StopTimer()
	MemUsage()
}

func TestDB_Items(t *testing.T) {

	tests := []struct {
		loops    int
		expected uint64
	}{
		{100, 100},
		{1000, 1000},
		{10000, 10000},
		{100000, 100000},
		{1000000, 1000000},
		{10000000, 10000000},
	}

	for _, test := range tests {
		setdb := Init()

		for i := 0; i < test.loops; i++ {
			ids := xid.New()
			setdb.PutString(ids.String())
		}

		assert.Equal(t, test.expected, setdb.Items())

		runtime.GC()
	}
}
