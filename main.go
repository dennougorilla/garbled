package main

import (
	"fmt"
)

func main() {
	strRaw := "\x8c\x8e\x93"
	bytesRaw := []byte(strRaw)
	s := "a"
	fmt.Printf("%v: %x\n", s, s)
}
