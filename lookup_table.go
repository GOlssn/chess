package main

import "fmt"

// ClearRank returns a Bitboard where all bits on the provided rank have been set to 0
func ClearRank(bitboard Bitboard, rank int) (Bitboard, error) {
	switch rank {
	case 1:
		bitboard = bitboard & Bitboard(0xffffffffffffff00)
	case 2:
		bitboard = bitboard & Bitboard(0xffffffffffff00ff)
	case 3:
		bitboard = bitboard & Bitboard(0xffffffffff00ffff)
	case 4:
		bitboard = bitboard & Bitboard(0xffffffff00ffffff)
	case 5:
		bitboard = bitboard & Bitboard(0xffffff00ffffffff)
	case 6:
		bitboard = bitboard & Bitboard(0xffff00ffffffffff)
	case 7:
		bitboard = bitboard & Bitboard(0xff00ffffffffffff)
	case 8:
		bitboard = bitboard & Bitboard(0x00ffffffffffffff)
	default:
		return Bitboard(0), fmt.Errorf("Must provide a rank between 1 and 8, got %d", rank)
	}

	return bitboard, nil
}

// MaskRank returns a Bitboard where all set bits on the provided rank have been kept
func MaskRank(bitboard Bitboard, rank int) (Bitboard, error) {
	switch rank {
	case 1:
		bitboard = bitboard & Bitboard(0x00000000000000ff)
	case 2:
		bitboard = bitboard & Bitboard(0x000000000000ff00)
	case 3:
		bitboard = bitboard & Bitboard(0x0000000000ff0000)
	case 4:
		bitboard = bitboard & Bitboard(0x00000000ff000000)
	case 5:
		bitboard = bitboard & Bitboard(0x000000ff00000000)
	case 6:
		bitboard = bitboard & Bitboard(0x0000ff0000000000)
	case 7:
		bitboard = bitboard & Bitboard(0x00ff000000000000)
	case 8:
		bitboard = bitboard & Bitboard(0xff00000000000000)
	default:
		return Bitboard(0), fmt.Errorf("Must provide a rank between 1 and 8, got %d", rank)
	}

	return bitboard, nil
}

// ClearFile returns a Bitboard where all bits on the provided file have been set to 0
func ClearFile(bitboard Bitboard, file int) (Bitboard, error) {
	switch file {
	case 1:
		bitboard = bitboard & Bitboard(0xfefefefefefefefe)
	case 2:
		bitboard = bitboard & Bitboard(0xfdfdfdfdfdfdfdfd)
	case 3:
		bitboard = bitboard & Bitboard(0xfbfbfbfbfbfbfbfb)
	case 4:
		bitboard = bitboard & Bitboard(0xf7f7f7f7f7f7f7f7)
	case 5:
		bitboard = bitboard & Bitboard(0xefefefefefefefef)
	case 6:
		bitboard = bitboard & Bitboard(0xdfdfdfdfdfdfdfdf)
	case 7:
		bitboard = bitboard & Bitboard(0xbfbfbfbfbfbfbfbf)
	case 8:
		bitboard = bitboard & Bitboard(0x7f7f7f7f7f7f7f7f)
	default:
		return Bitboard(0), fmt.Errorf("Must provide a file between 1 and 8, got %d", file)
	}

	return bitboard, nil
}

// MaskFile returns a Bitboard where all set bits on the provided file have been kept
func MaskFile(bitboard Bitboard, file int) (Bitboard, error) {
	switch file {
	case 1:
		bitboard = bitboard & Bitboard(0x0101010101010101)
	case 2:
		bitboard = bitboard & Bitboard(0x0202020202020202)
	case 3:
		bitboard = bitboard & Bitboard(0x0404040404040404)
	case 4:
		bitboard = bitboard & Bitboard(0x0808080808080808)
	case 5:
		bitboard = bitboard & Bitboard(0x1010101010101010)
	case 6:
		bitboard = bitboard & Bitboard(0x2020202020202020)
	case 7:
		bitboard = bitboard & Bitboard(0x4040404040404040)
	case 8:
		bitboard = bitboard & Bitboard(0x8080808080808080)
	default:
		return Bitboard(0), fmt.Errorf("Must provide a file between 1 and 8, got %d", file)
	}

	return bitboard, nil
}
