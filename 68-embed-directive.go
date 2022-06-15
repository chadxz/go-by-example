package main

import (
	"embed"
	_ "embed"
)

//go:embed 66-directories.go
var fileString string

//go:embed 68-embed-directive.go
var fileByte []byte

//go:embed 67-temporary-files-and-dirs.go
//go:embed 5*.go
var miscFiles embed.FS

func main() {
	print(fileString)
	print(string(fileByte))

	content1, _ := miscFiles.ReadFile("59-url-parsing.go")
	print(string(content1))

	content2, _ := miscFiles.ReadFile("55-epoch.go")
	print(string(content2))
}
