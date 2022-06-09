package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)

	go func(done chan bool) {
		fmt.Print("working...")
		time.Sleep(time.Second * 2)
		fmt.Println("done")

		done <- true
	}(done)

	<-done // SYNCHRONIZE!
}
