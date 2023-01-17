package dbIndexing

import (
	"log"
	"os"
	"strings"
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

func AddIndex(filename string, data []byte) {
	fi, err := os.Stat(filename)
	check(err)
	offset := fi.Size()
	str := strings.Split(string(data), ":")
	key := str[0]
	size := len(data)
	p := Position{Offset: offset, Size: size}
	Index[key] = p
}

//func GetIndex(key string) Position {
//
//}
