package main

import (
	"fmt"
	"log"
	"strangerDB/dbIndexing"
	"strangerDB/dbService"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	kv := dbService.KeyValue{Key: "Freeman", Value: "Lox"}
	dbIndexing.AddIndex("chunks/data", kv)
	dbService.WriteRecord("chunks/data", kv)
	p := dbIndexing.Index["Freeman"]
	fmt.Println(string(dbService.ReadRecord("chunks/data", p.Offset, p.Size)))
}
