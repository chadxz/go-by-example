package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func RemoveFile(path string) {
	err := os.Remove(path)
	if err != nil {
		panic(err)
	}
}

func RemoveAll(dir string) {
	err := os.RemoveAll(dir)
	if err != nil {
		panic(err)
	}
}

func main() {
	tmpFile, err := os.CreateTemp("", "sample")
	check(err)
	defer RemoveFile(tmpFile.Name())

	fmt.Println("Temp file name:", tmpFile.Name())

	_, err = tmpFile.Write([]byte{1, 2, 3, 4})
	check(err)

	tmpDir, err := os.MkdirTemp("", "sampledir")
	check(err)
	defer RemoveAll(tmpDir)

	fmt.Println("Temp dir name:", tmpDir)

	fileName := filepath.Join(tmpDir, "file1")
	check(os.WriteFile(fileName, []byte{1, 2}, 0666))
}
