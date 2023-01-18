package dbIndexing

import (
	"log"
	"os"
	"strangerDB/dbService"
)

var Index = make(map[string]Position)

type Position struct {
	Offset int64
	Size   int
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func AddIndex(filename string, data dbService.KeyValue) {
	fi, err := os.Stat(filename)
	check(err)
	offset := fi.Size()
	key := data.Key
	size := len([]byte(data.Value+data.Key)) + 1
	p := Position{Offset: offset, Size: size}
	Index[key] = p
}
