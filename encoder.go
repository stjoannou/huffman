package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

var encoding = []string{
	"00111110",  // 0
	"110",       // 1
	"1110",      // 2
	"11110",     // 3
	"00110",     // 4
	"0011100",   // 5
	"0011110",   // 6
	"0011101",   // 7
	"001111110", // 8
	"001111111", // 9
	"10",        // A
	"00101",     // B
	"01",        // C
	"000",       // D
	"11111",     // E
	"00100",     // F
}

func convert(s string) string {
	var returnStr string
	for i := 0; i < len(s); i++ {
		decoded, err := hex.DecodeString("0" + string(s[i]))
		if err != nil {
			log.Fatal(err)
		}
		var dec = decoded[0] - 1
		// fmt.Printf("%s ", encoding[dec])
		returnStr += encoding[dec]

	}
	// fmt.Printf("\n")
	return returnStr
}

func main() {
	fmt.Printf("Huffman Encoder in Go\n")

	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <string>\n", os.Args[0])
		os.Exit(-1)
	}

	src := []byte(os.Args[1])
	encodedStr := hex.EncodeToString(src)

	fmt.Printf("Hex String [%s]\n", encodedStr)

	fmt.Printf("Encoded: [%s]\n", convert(encodedStr))

}
