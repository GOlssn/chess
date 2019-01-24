package main

import (
	"fmt"
)

func main() {
	bitboard := Bitboard(0x404040404040404)
	fmt.Printf("%b\n", uint64(bitboard))
}
