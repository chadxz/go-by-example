package main

import "fmt"

func willPanic() {
	panic("boom")
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from error:\n", err)
		}
	}()

	willPanic()

	fmt.Println("I have survived!")
}
