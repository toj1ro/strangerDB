package dbService

import (
	"fmt"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func ReadFromFile(filename string, start int, end int) []byte {
	f, err := os.Open(filename)
	check(err)
	defer f.Close()
	bytes := make([]byte, 5)
	_, err2 := f.Read(bytes)
	check(err2)
	return bytes[start:end]

}

func WriteToFile(filename string, data []byte) {
	f, err := os.OpenFile(filename, os.O_WRONLY, os.ModeAppend)
	check(err)
	defer f.Close()

	_, err2 := f.Write(data)
	check(err2)
	fmt.Print("Done\n")

}

func Delete(filename string, data string) {

}
