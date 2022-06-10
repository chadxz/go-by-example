package main

import "os"

func main() {
	//panic("something bad happened")

	_, err := os.Create("/file")
	if err != nil {
		panic(err)
	}
}
