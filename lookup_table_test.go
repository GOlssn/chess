package main

import "testing"

func TestClearRank(t *testing.T) {
	testCases := []struct {
		bitboard Bitboard
		rank     int
		want     Bitboard
	}{
		{Bitboard(0xffffffffffffffff), 1, Bitboard(0xffffffffffffff00)},
		{Bitboard(0xffffffffffffffff), 2, Bitboard(0xffffffffffff00ff)},
		{Bitboard(0xffffffffffffffff), 3, Bitboard(0xffffffffff00ffff)},
		{Bitboard(0xffffffffffffffff), 4, Bitboard(0xffffffff00ffffff)},
		{Bitboard(0xffffffffffffffff), 5, Bitboard(0xffffff00ffffffff)},
		{Bitboard(0xffffffffffffffff), 6, Bitboard(0xffff00ffffffffff)},
		{Bitboard(0xffffffffffffffff), 7, Bitboard(0xff00ffffffffffff)},
		{Bitboard(0xffffffffffffffff), 8, Bitboard(0x00ffffffffffffff)},
	}

	for _, tc := range testCases {
		got, err := ClearRank(tc.bitboard, tc.rank)

		if err != nil {
			t.Fatalf("Provided rank does not exist on a chess board")
		}

		if got != tc.want {
			t.Errorf("Got '%b', want '%b'", got, tc.want)
		}
	}
}

func TestMaskRank(t *testing.T) {
	testCases := []struct {
		bitboard Bitboard
		rank     int
		want     Bitboard
	}{
		{Bitboard(0xff000000000000ff), 1, Bitboard(0x00000000000000ff)},
		{Bitboard(0xff0000000000ff00), 2, Bitboard(0x000000000000ff00)},
		{Bitboard(0xff00000000ff0000), 3, Bitboard(0x0000000000ff0000)},
		{Bitboard(0xff000000ff000000), 4, Bitboard(0x00000000ff000000)},
		{Bitboard(0xff0000ff00000000), 5, Bitboard(0x000000ff00000000)},
		{Bitboard(0xff00ff0000000000), 6, Bitboard(0x0000ff0000000000)},
		{Bitboard(0xffff000000000000), 7, Bitboard(0x00ff000000000000)},
		{Bitboard(0xff000000000000ff), 8, Bitboard(0xff00000000000000)},
	}

	for _, tc := range testCases {
		got, err := MaskRank(tc.bitboard, tc.rank)

		if err != nil {
			t.Fatalf("Provided rank does not exist on a chess board")
		}

		if got != tc.want {
			t.Errorf("Got '%b', want '%b'", got, tc.want)
		}
	}
}

func TestClearFile(t *testing.T) {
	testCases := []struct {
		bitboard Bitboard
		file     int
		want     Bitboard
	}{
		{Bitboard(0xffffffffffffffff), 1, Bitboard(0xfefefefefefefefe)},
		{Bitboard(0xffffffffffffffff), 2, Bitboard(0xfdfdfdfdfdfdfdfd)},
		{Bitboard(0xffffffffffffffff), 3, Bitboard(0xfbfbfbfbfbfbfbfb)},
		{Bitboard(0xffffffffffffffff), 4, Bitboard(0xf7f7f7f7f7f7f7f7)},
		{Bitboard(0xffffffffffffffff), 5, Bitboard(0xefefefefefefefef)},
		{Bitboard(0xffffffffffffffff), 6, Bitboard(0xdfdfdfdfdfdfdfdf)},
		{Bitboard(0xffffffffffffffff), 7, Bitboard(0xbfbfbfbfbfbfbfbf)},
		{Bitboard(0xffffffffffffffff), 8, Bitboard(0x7f7f7f7f7f7f7f7f)},
	}

	for _, tc := range testCases {
		got, err := ClearFile(tc.bitboard, tc.file)

		if err != nil {
			t.Fatalf("Provided file does not exist on a chess board")
		}

		if got != tc.want {
			t.Errorf("Got '%b', want '%b'", got, tc.want)
		}
	}
}

func TestMaskFile(t *testing.T) {
	testCases := []struct {
		bitboard Bitboard
		file     int
		want     Bitboard
	}{
		{Bitboard(0xffffffffffffffff), 1, Bitboard(0x0101010101010101)},
		{Bitboard(0xffffffffffffffff), 2, Bitboard(0x0202020202020202)},
		{Bitboard(0xffffffffffffffff), 3, Bitboard(0x0404040404040404)},
		{Bitboard(0xffffffffffffffff), 4, Bitboard(0x0808080808080808)},
		{Bitboard(0xffffffffffffffff), 5, Bitboard(0x1010101010101010)},
		{Bitboard(0xffffffffffffffff), 6, Bitboard(0x2020202020202020)},
		{Bitboard(0xffffffffffffffff), 7, Bitboard(0x4040404040404040)},
		{Bitboard(0xffffffffffffffff), 8, Bitboard(0x8080808080808080)},
	}

	for _, tc := range testCases {
		got, err := MaskFile(tc.bitboard, tc.file)

		if err != nil {
			t.Fatalf("Provided file does not exist on a chess board")
		}

		if got != tc.want {
			t.Errorf("Got '%b', want '%b'", got, tc.want)
		}
	}
}
