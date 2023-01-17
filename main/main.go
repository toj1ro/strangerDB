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
	dbIndexing.AddIndex("chunks/data", data)
	dbService.WriteRecord("chunks/data", data)
	p := dbIndexing.Index["we"]
	fmt.Println(string(dbService.ReadRecord("chunks/data", p.Offset, p.Size)))
	text, _ := dbService.Read("chunks/data", "FOXICK")
	fmt.Println(text)

}

//Stividoh
//[10 208 187 208 190 209 133]
