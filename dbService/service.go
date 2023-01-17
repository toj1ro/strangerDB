package dbService

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func ReadRecord(filename string, offset int64, size int) []byte {
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

func WriteRecord(filename string, data []byte) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModeAppend)
	check(err)
	defer f.Close()
	sep := []byte("\n")
	_, err2 := f.Write(sep)
	check(err2)
	_, err3 := f.Write([]byte(data))
	check(err3)
}

func DeleteRecord(filename string, data []byte) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModeAppend)
	check(err)
	defer f.Close()
	size := []byte(string(0))
	_, err2 := f.Write(size)
	check(err2)
	_, err3 := f.Write(data)
	check(err3)
}

func reverseText(text []string) []string {
	if len(text) == 0 {
		return text
	}
	return append(reverseText(text[1:]), text[0])
}

func Read(filename string, key string) (string, error) {
	f, err := os.Open(filename)
	check(err)
	defer f.Close()
	info, err2 := f.Stat()
	check(err2)
	bytes := make([]byte, info.Size())
	_, err3 := f.Read(bytes)
	check(err3)
	text := string(bytes)
	textArray := strings.Split(text, "\n")
	reverseTextArray := reverseText(textArray[1:])
	for _, i := range reverseTextArray {
		key_value := strings.Split(i, ":")
		if key_value[0] == key {
			return i, nil
		}
	}
	return "", fmt.Errorf("не нашло")
}
