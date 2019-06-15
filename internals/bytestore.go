package internals

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
		root = root.Put(b)
	}
	root.Set()
}

func (db *DB) PutString(s string) {
	db.Put([]byte(s))
}

func (db DB) Find(bs []byte) bool {
	root := db.root
	for _, b := range bs {
		next := root.Find(b)

		if next == nil {
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

	log.Printf("\n")
	log.Printf("\tHeap alloc = %v MiB\n", bToMb(m.HeapAlloc))
	log.Printf("\tHeap inuse = %v MiB\n", bToMb(m.HeapInuse))
	log.Printf("\tHeap objects = %v\n", m.HeapObjects)

}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}