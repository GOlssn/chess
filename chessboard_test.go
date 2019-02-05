package main

import (
	"reflect"
	"testing"
)

func TestValidPawnMoves(t *testing.T) {
	testCases := []struct {
		board Chessboard
		side  Color
		want  Bitboard
	}{
		{Chessboard{WhitePawns: 0x8000}, White, Bitboard(0x80800000)},
		{Chessboard{WhitePawns: 0x800000, BlackPieces: 0x40000000}, White, Bitboard(0xc0000000)},
		{Chessboard{WhitePawns: 0x20000400, BlackPieces: 0x1000000000}, White, Bitboard(0x3004040000)},
	}

	for _, tc := range testCases {
		got := tc.board.ValidPawnMoves(White)

		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Got '%b', want '%b'", got, tc.want)
		}
	}

}
