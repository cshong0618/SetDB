package setdb

import (
	"log"
	"runtime"
	"unsafe"
)

type DB struct {
	root *ByteUnit
	count int
}

// Root is always 0.
// Uses its map to hold actual value
func Init() *DB {
	db := DB{}
	db.root = InitByteUnit(byte(0))

	return &db
}

func (db *DB) Put(bs []byte) {
	// Put an array of bytes into the root
	root := db.root

	for _, b := range bs {
		if _, ok := root.next[b]; !ok {
			root.next[b] = InitByteUnit(b)
		}
		root = root.next[b]
	}
	root.Set()
}

func (db *DB) PutString(s string) {
	db.Put([]byte(s))
}

func (db DB) Find(bs []byte) bool {
	root := db.root
	for _, b := range bs {
		next, ok := root.next[b]
		if !ok {
			return false
		}
		root = next
	}

	return root.exists
}

func (db DB) Size() uintptr {
	sizeofDb := unsafe.Sizeof(db.root)
	return sizeofDb
}

func MemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	log.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	log.Printf("\tSys = %v MiB", bToMb(m.Sys))
	log.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}