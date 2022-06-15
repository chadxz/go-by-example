package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func RemoveAll(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		panic(err)
	}
}

func main() {
	tmpDir := "out/subdir"
	err := os.MkdirAll(tmpDir, 0755)
	check(err)
	defer RemoveAll(tmpDir)

	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile(fmt.Sprintf("%s/file1", tmpDir))

	err = os.MkdirAll(fmt.Sprintf("%s/parent/child", tmpDir), 0755)
	check(err)

	createEmptyFile(fmt.Sprintf("%s/parent/file2", tmpDir))
	createEmptyFile(fmt.Sprintf("%s/parent/file3", tmpDir))
	createEmptyFile(fmt.Sprintf("%s/parent/child/file4", tmpDir))

	directoryContents, err := os.ReadDir(fmt.Sprintf("%s/parent", tmpDir))
	check(err)

	fmt.Printf("Listing %s/parent\n", tmpDir)
	for _, entry := range directoryContents {
		fmt.Println(" ", entry.Name(), "is dir?", entry.IsDir())
	}

	err = os.Chdir(fmt.Sprintf("%s/parent/child", tmpDir))
	check(err)

	directoryContents, err = os.ReadDir(".")
	check(err)

	fmt.Printf(fmt.Sprintf("Listing %s/parent/child\n", tmpDir))
	for _, entry := range directoryContents {
		fmt.Println(" ", entry.Name(), "is dir?", entry.IsDir())
	}

	err = os.Chdir("../../../..") // move back so defer can cleanup
	check(err)

	fmt.Println(fmt.Sprintf("Visiting %s", tmpDir))
	err = filepath.Walk(tmpDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(" ", path, "is dir?", info.IsDir())
		return nil
	})
}
