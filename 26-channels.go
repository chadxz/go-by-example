package main

import (
	"encoding/json"
	"fmt"
)

type Wave struct {
	Hand    string `json:"hand,omitempty"`
	Fingers int    `json:"fingers,omitempty"`
}

func main() {
	messages := make(chan Wave)

	go func() { messages <- Wave{Hand: "right", Fingers: 5} }()

	msg := <-messages

	if j, err := json.Marshal(msg); err == nil {
		fmt.Println(string(j))
	}
}
