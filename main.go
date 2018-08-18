package main

import (
	"fmt"
)

func main() {
	bytes := []byte{0x82, 0xA0}
	s := string(bytes)
	fmt.Printf("%v: %x\n", s, s)
	const sample = "\xbd"
	fmt.Println(sample)
	fmt.Println("%x", sample)

}
