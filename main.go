package main

import (
	"fmt"
)

func main() {
	bytes := []byte{0x82, 0xA0}
	s := string(bytes)
	fmt.Printf("%v: %x\n", s, s)
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Println(sample)
	fmt.Printf("% x\n", sample)

}
