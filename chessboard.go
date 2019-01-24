package main

type Bitboard uint64

type Chessboard struct {
	WhitePawns   Bitboard
	WhiteKnights Bitboard
	WhiteBishops Bitboard
	WhiteRooks   Bitboard
	WhiteQueens  Bitboard
	WhiteKing    Bitboard

	BlackPawns   Bitboard
	BlackKnights Bitboard
	BlackBishops Bitboard
	BlackRooks   Bitboard
	BlackQueens  Bitboard
	BlackKing    Bitboard

	WhitePieces Bitboard
	BlackPieces Bitboard
	AllPieces   Bitboard
}

// New sets up the initial position of all pieces and returns a pointer to a Chessboard object
func NewChessboard() *Chessboard {
	chessBoard := Chessboard{
		WhiteRooks:   Bitboard(0x0000000000000081),
		WhiteKnights: Bitboard(0x0000000000000042),
		WhiteBishops: Bitboard(0x0000000000000024),
		WhiteQueens:  Bitboard(0x0000000000000008),
		WhiteKing:    Bitboard(0x0000000000000010),
		WhitePawns:   Bitboard(0x000000000000ff00),
		BlackRooks:   Bitboard(0x8100000000000000),
		BlackKnights: Bitboard(0x4200000000000000),
		BlackBishops: Bitboard(0x2400000000000000),
		BlackQueens:  Bitboard(0x0800000000000000),
		BlackKing:    Bitboard(0x1000000000000000),
		BlackPawns:   Bitboard(0x00ff000000000000),
	}

	return &chessBoard
}

func ValidPawnMoves(b Bitboard) Bitboard {
	return b << 8
}

type Square int8

const (
	A1 Square = iota
	A2
	A3
	A4
	A5
	A6
	A7
	A8
	B1
	B2
	B3
	B4
	B5
	B6
	B7
	B8
	C1
	C2
	C3
	C4
	C5
	C6
	C7
	C8
	D1
	D2
	D3
	D4
	D5
	D6
	D7
	D8
	E1
	E2
	E3
	E4
	E5
	E6
	E7
	E8
	F1
	F2
	F3
	F4
	F5
	F6
	F7
	F8
	G1
	G2
	G3
	G4
	G5
	G6
	G7
	G8
	H1
	H2
	H3
	H4
	H5
	H6
	H7
	H8
)
