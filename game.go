package main

type Color int

type Move struct {
	Origin      Square
	Destination Square
}

const (
	Black Color = iota
	White
)

type Game struct {
	Turn  Color
	Board *Chessboard
}
