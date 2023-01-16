package dbService

import (
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func ReadFromFile(filename string, offset int64, size int) []byte {
	f, err := os.Open(filename)
	check(err)
	defer f.Close()
	bytes := make([]byte, size+1)
	_, err2 := f.Seek(offset, 0)
	check(err2)
	_, err3 := f.Read(bytes)
	check(err3)
	return bytes[1:]
}

func WriteToFile(filename string, data []byte) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModeAppend)
	check(err)
	defer f.Close()
	size := []byte(string(len(data)))
	_, err2 := f.Write(size)
	check(err2)
	_, err3 := f.Write(data)
	check(err3)
}

func Delete(filename string, data string) {

}
