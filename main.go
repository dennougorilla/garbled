package main

import (
	"fmt"
)

func main() {
	bytes := []byte{0x82, 0xA0}
	s := string(bytes)
	fmt.Printf(s)
	fmt.Printf("%v: %x\n", s, s)
}
