package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	input1 := []byte("1c0111001f010100061a024b53535009181c")
	input2 := []byte("686974207468652062756c6c277320657965")
	hexString1 := decodeHex(input1)
	hexString2 := decodeHex(input2)
	r := fixedXor(hexString1, hexString2)
	result := make([]byte, hex.EncodedLen(len(r)))
	hex.Encode(result, r)
	fmt.Printf("%s\n", result)
}

func decodeHex(s []byte) []byte {
	dst := make([]byte, hex.DecodedLen(len(s)))
	_, err := hex.Decode(dst, s)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s\n", dst[:n])
	return dst
}

func fixedXor(a []byte, b []byte) []byte {
	c := make([]byte, len(a))
	for i := 0; i < len(a); i++ {
		c[i] = a[i] ^ b[i]

	}
	return c
}
