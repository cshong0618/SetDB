package main

import (
	"github.com/google/uuid"
	"log"
	"setdb/internals"
)

func main() {
	db := internals.Init()
	loops := 1000000
	for i := 0; i < loops; i++ {
		log.Println(i)
		uuids := uuid.New()
		db.PutString(uuids.String())
	}

	internals.MemUsage()
}
