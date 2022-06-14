package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~"
	fmt.Println("original string:", data)
	encoded := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("std encoded:", encoded)
	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	fmt.Println("std decoded:", string(decoded))

	fmt.Println()

	urlEncoded := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println("url encoded:", urlEncoded)
	urlDecoded, _ := base64.URLEncoding.DecodeString(urlEncoded)
	fmt.Println("url decoded:", string(urlDecoded))

}
