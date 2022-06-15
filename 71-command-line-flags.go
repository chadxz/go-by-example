package main

import (
	"flag"
	"fmt"
)

func main() {
	wordFlag := flag.String("word", "foo", "a string")
	numbFlag := flag.Int("numb", 42, "an int")
	forkFlag := flag.Bool("fork", false, "a bool")
	var svarFlag string
	flag.StringVar(&svarFlag, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *wordFlag)
	fmt.Println("numb:", *numbFlag)
	fmt.Println("fork:", *forkFlag)
	fmt.Println("stringVarFlag:", svarFlag)
	fmt.Println("tail:", flag.Args())
}
