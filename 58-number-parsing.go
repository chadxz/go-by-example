package main

import (
	"fmt"
	"strconv"
)

func main() {
	myfloat64, _ := strconv.ParseFloat("1.234", 64)
	fmt.Printf("%T %f\n", myfloat64, myfloat64)

	myfloat64two, _ := strconv.ParseFloat("1.234", 32)
	myfloat32 := float32(myfloat64two)
	fmt.Printf("%T %f\n", myfloat32, myfloat32)

	myint, _ := strconv.ParseInt("123", 0, 64)
	fmt.Printf("%T %d\n", myint, myint)

	myinttwo, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Printf("%T %d\n", myinttwo, myinttwo)

	myuint, _ := strconv.ParseUint("789", 0, 64)
	fmt.Printf("%T %d\n", myuint, myuint)

	simpleint, _ := strconv.Atoi("135")
	fmt.Printf("%T %d\n", simpleint, simpleint)

	_, e := strconv.Atoi("notanumber")
	fmt.Printf("%T %s\n", e, e)
}
