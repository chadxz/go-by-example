package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	if _, err := fmt.Fprintf(w, "hello\n"); err != nil {
		panic(err)
	}
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			if _, err := fmt.Fprintf(w, "%v: %v\n", name, h); err != nil {
				panic(err)
			}
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	if err := http.ListenAndServe(":8090", nil); err != nil {
		panic(err)
	}
}
