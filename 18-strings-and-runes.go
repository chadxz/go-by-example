package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const h = "hello"
	const s = "สวัสดี"
	fmt.Println("len h:", len(h))
	fmt.Println("len s:", len(s))
	fmt.Print("h bytes: ")
	for i := 0; i < len(h); i++ {
		fmt.Printf("0x%X ", h[i])
	}
	fmt.Println()
	fmt.Print("s bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("0x%X ", s[i])
	}
	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for i, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, i)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
