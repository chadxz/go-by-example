package main

import (
	"fmt"
	"time"
)

func f(from string, pause bool) {
	for i := 0; i < 5; i++ {
		if i == 3 && pause {
			time.Sleep(time.Second)
		}
		fmt.Println(from, ":", i)
	}
}

func main() {
	go f("goroutine", true)

	f("direct", true)

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second * 2)
	fmt.Println("done")
}
