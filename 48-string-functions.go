package main

import (
	"fmt"
	str "strings"
)

func main() {
	fmt.Println("Contains: ", str.Contains("test", "es"))
	fmt.Println("Count:    ", str.Count("test", "t"))
	fmt.Println("HasPrefix:", str.HasPrefix("test", "te"))
	fmt.Println("HasSuffix:", str.HasSuffix("test", "st"))
	fmt.Println("Index:    ", str.Index("test", "e"))
	fmt.Println("Join:     ", str.Join([]string{"a", "b"}, "-"))
	fmt.Println("Repeat:   ", str.Repeat("a", 5))
	fmt.Println("Replace:  ", str.Replace("foo", "o", "0", -1))
	fmt.Println("Replace-: ", str.Replace("foo", "o", "0", 1))
	fmt.Println("Split:    ", str.Split("a-b-c-d-e", "-"))
	fmt.Println("ToLower:  ", str.ToLower("TEST"))
	fmt.Println("ToUpper:  ", str.ToUpper("test"))
}
