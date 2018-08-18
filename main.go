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
	const sample = "\xe6\x9b\x84"
	fmt.Println(sample)
	fmt.Printf("% x\n", sample)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100000000; i++ {
		src := int64(i)
		result := make([]byte, binary.MaxVarintLen64)
		binary.PutVarint(result, src)
		fmt.Printf("% x\n", result[0:3])
		fmt.Println(string(result[0:3]))
		if string(result[0:3]) == "" {
			break
		}
	}
	s = "曄"
	fmt.Printf("%v: % x\n", s, s)
}

//print(int(random.randrange(0,65535)).to_bytes(2,byteorder='little').decode('shift-jis','replace'), end='')
