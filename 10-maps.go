package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["foo"] = 3
	m["bar"] = 5

	fmt.Println("map:", m)

	v1 := m["foo"]
	fmt.Println("value!", v1)
	fmt.Println("length!", len(m))

	delete(m, "bar")
	fmt.Println("map:", m)

	if _, ok := m["missing"]; !ok {
		fmt.Println("key not found :(")
	}

	m2 := map[string]int{"foo": 3, "bar": 5}
	fmt.Println("map2", m2)
}
