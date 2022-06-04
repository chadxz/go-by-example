package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("empty:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("slice", s)
	fmt.Println("s[0]", s[0])
	fmt.Println("length", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("slice appended", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copied slice", c)

	fmt.Println("slice of slice", s[2:5])
	fmt.Println("slice of slice", s[:5])
	fmt.Println("slice of slice", s[2:])

	inline := []string{"x", "y", "z"}
	fmt.Println(inline)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD", twoD)
}
