package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strangerDB/dbIndexing"
	"strangerDB/dbService"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	data, err2 := in.ReadBytes('\n')
	data = data[:len(data)-2]
	check(err2)
	dbIndexing.AddIndex("chunks/data", string(data))
	dbService.WriteToFile("chunks/data", data)
	p := dbIndexing.Index["we"]
	fmt.Println(p.Offset, p.Size)
	fmt.Println(string(dbService.ReadFromFile("chunks/data", p.Offset, p.Size)))
}
