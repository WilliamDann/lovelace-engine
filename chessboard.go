package main

type Piece int

const (
	none Piece = 0

	white_pawn   Piece = iota
	white_rook   Piece = iota
	white_bishop Piece = iota
	white_knight Piece = iota
	white_queen  Piece = iota
	white_king   Piece = iota

	black_pawn   Piece = iota
	black_rook   Piece = iota
	black_bishop Piece = iota
	black_knight Piece = iota
	black_queen  Piece = iota
	black_king   Piece = iota
)

var piece_symbols map[Piece]string = map[Piece]string{
	white_pawn:   "P",
	white_rook:   "R",
	white_bishop: "B",
	white_knight: "N",
	white_queen:  "Q",
	white_king:   "K",

	black_pawn:   "p",
	black_rook:   "r",
	black_bishop: "b",
	black_knight: "n",
	black_queen:  "q",
	black_king:   "k",

	none: "_",
}

type Chessboard struct {
	squares [64]Piece
}

func NewChessboard() Chessboard {
	return Chessboard{[64]Piece{}}
}

func (board Chessboard) String() string {
	output := ""
	for i, piece := range board.squares {
		output += " " + piece_symbols[piece]

		if (i+1)%8 == 0 && i != 0 {
			output += "\n"
		}
	}
	return output
}
