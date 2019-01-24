package main

import (
	"reflect"
	"testing"
)

func TestValidPawnMoves(t *testing.T) {
	got := ValidPawnMoves(Bitboard(0x8000))
	want := Bitboard(0x800000)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got '%b', want '%b'", got, want)
	}
}
