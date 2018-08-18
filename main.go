package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	bytes := []byte{0xbd, 0xb2, 0x3d, 0xbc, 0x20, 0xe2, 0x8c, 0x98}
	s := string(bytes)
	fmt.Printf("% x\n", s)
	fmt.Println(s)
	const sample = "\x9e\x40"
	fmt.Println(sample)
	fmt.Printf("% x\n", sample)

	rand.Seed(time.Now().UnixNano())
	src := int64(rand.Intn(65535))
	result := make([]byte, binary.MaxVarintLen64)
	binary.PutVarint(result, src)
	fmt.Printf("% x\n", result[0:2])
	fmt.Println(string(result[0:2]))
}

//print(int(random.randrange(0,65535)).to_bytes(2,byteorder='little').decode('shift-jis','replace'), end='')
