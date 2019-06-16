package internal

import (
	"log"
	"runtime"
	"sync"
	"unsafe"
)

type DB struct {
	root       *ByteUnit
	count      uint64
	writeMutex *sync.Mutex
}

// Root is always 0.
// Uses its map to hold actual value
func InitDB() *DB {
	db := DB{}
	db.root = InitByteUnit(byte(0))
	db.writeMutex = &sync.Mutex{}
	db.count = 0

	return &db
}

func (db *DB) Put(bs []byte) {
	// Lock the db
	db.writeMutex.Lock()

	// Put an array of bytes into the root
	root := db.root

	for _, b := range bs {
		root = root.Put(b)
	}

	if !root.exists {
		root.Set()
		db.count++
	}

	db.writeMutex.Unlock()
}

func (db *DB) PutString(s string) {
	db.Put([]byte(s))
}

func (db DB) Find(bs []byte) bool {
	root := db.root
	for _, b := range bs {
		root = root.Find(b)
		if root == nil {
			return false
		}
	}

	return root.exists
}

func (db DB) FindString(s string) bool {
	return db.Find([]byte(s))
}

func (db DB) Size() uintptr {
	sizeofDb := unsafe.Sizeof(db.root)
	return sizeofDb
}

func (db DB) Items() uint64 {
	return db.count
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
