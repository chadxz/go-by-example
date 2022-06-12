package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	// does the regex match this?
	fmt.Println(r.MatchString("peach"))
	// find regex match
	fmt.Println(r.FindString("peach punch"))
	// start and end indices
	fmt.Println("idx:", r.FindStringIndex("peach punch"))
	// whole pattern and capture group(s)
	fmt.Println(r.FindStringSubmatch("peach punch"))
	// start/end indices of whole match and capture group(s)
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	fmt.Println("all:", r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	fmt.Println("only two:", r.FindAllString("peach punch pinch", 2))
	fmt.Println(regexp.Match("p([a-z]+)ch", []byte("peach")))
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
