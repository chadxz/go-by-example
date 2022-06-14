package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	path := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("path:", path)

	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	fmt.Println("Dir(path):", filepath.Dir(path))
	fmt.Println("Base(path):", filepath.Base(path))
	d, b := filepath.Split(path)
	fmt.Println("Split:", d, b)

	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"
	extension := filepath.Ext(filename)
	fmt.Println("extension:", extension)
	fmt.Println(strings.TrimSuffix(filename, extension))

	relative, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(relative)

	relative, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(relative)
}
