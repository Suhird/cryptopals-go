package main

import (
	b64 "encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	input := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	hexString := decodeHex(input)
	result := b64.StdEncoding.EncodeToString(hexString)
	fmt.Println(result)
}

func decodeHex(s []byte) []byte {
	dst := make([]byte, hex.DecodedLen(len(s)))
	n, err := hex.Decode(dst, s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", dst[:n])
	return dst
}
