package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		log.Panicf("error: %v\n", err)
	}
}

func main() {
	data1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/data1", data1, 0644)
	check(err)

	file, err := os.Create("/tmp/data2")
	check(err)

	defer closeFile(file)

	data2 := []byte{115, 111, 109, 101, 10}
	bytesWritten2, err := file.Write(data2)
	check(err)
	fmt.Printf("wrote %d bytes\n", bytesWritten2)

	bytesWritten3, err := file.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", bytesWritten3)
	err = file.Sync()
	check(err)

	writer := bufio.NewWriter(file)
	bytesWritten4, err := writer.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", bytesWritten4)
	err = writer.Flush()
	check(err)

	fmt.Println()
	out, err := exec.Command("cat", "/tmp/data2").Output()
	check(err)
	fmt.Printf("file contents --v\n~~~~~~~~~~~~~~~~~\n%s", string(out))
}
