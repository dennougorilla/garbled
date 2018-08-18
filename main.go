package main

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"strings"
)

func main() {
	strRaw := "\x8c\x8e\x93"
	bytesRaw := []byte(strRaw)
	s := "a"
	fmt.Printf("%v: %x\n", s, s)
}
