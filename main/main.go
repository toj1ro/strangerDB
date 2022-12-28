package main

import (
	"bufio"
	"log"
	"os"
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
	check(err2)
	dbService.WriteToFile("db/backet", data)
	dbService.ReadFromFile("db/backet", 0, 3)
}
