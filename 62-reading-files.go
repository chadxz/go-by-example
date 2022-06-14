package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Panicf("error: %v\n", err)
	}
}

func main() {
	data, err := os.ReadFile("62-reading-files.go")
	check(err)
	fmt.Print(string(data))

	file, err := os.Open("62-reading-files.go")
	defer closeFile(file)
	check(err)

	bytes1 := make([]byte, 5)
	bytesRead1, err := file.Read(bytes1)
	check(err)
	fmt.Printf("%d bytes read: %s\n", bytesRead1, string(bytes1[:bytesRead1]))

	offset2, err := file.Seek(5, 0)
	check(err)
	bytes2 := make([]byte, 2)
	bytesRead2, err := file.Read(bytes2)
	check(err)
	fmt.Printf("%d bytes @ %d: %v\n", bytesRead2, offset2, string(bytes2[:bytesRead2]))

	offset3, err := file.Seek(5, 0)
	check(err)
	bytes3 := make([]byte, 2)
	bytesRead3, err := io.ReadAtLeast(file, bytes3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", bytesRead3, offset3, string(bytes3))

	_, err = file.Seek(0, 0)
	check(err)

	reader := bufio.NewReader(file)
	bytes4, err := reader.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(bytes4))
}
