package source

import "strconv"

type Piece int

const (
	None Piece = 0

	White_pawn   Piece = iota
	White_rook   Piece = iota
	White_bishop Piece = iota
	White_knight Piece = iota
	White_queen  Piece = iota
	White_king   Piece = iota

	Black_pawn   Piece = iota
	Black_rook   Piece = iota
	Black_bishop Piece = iota
	Black_knight Piece = iota
	Black_queen  Piece = iota
	Black_king   Piece = iota
)

var piece_symbols map[Piece]string = map[Piece]string{
	White_pawn:   "p",
	White_rook:   "r",
	White_bishop: "b",
	White_knight: "n",
	White_queen:  "q",
	White_king:   "k",

	Black_pawn:   "P",
	Black_rook:   "R",
	Black_bishop: "B",
	Black_knight: "N",
	Black_queen:  "Q",
	Black_king:   "K",

	None: "_",
}

type Chessboard struct {
	Squares [64]Piece
}

func NewChessboard() Chessboard {
	return Chessboard{[64]Piece{}}
}

// https://en.wikipedia.org/wiki/Forsyth%E2%80%93Edwards_Notation
func ParseFen(fen string) Chessboard {
	board := NewChessboard()

	symbols_to_pieces := make(map[string]Piece)
	for key, value := range piece_symbols {
		symbols_to_pieces[value] = key
	}

	i := 0
	for _, char := range fen {
		if val, ok := symbols_to_pieces[string(char)]; ok {
			board.Squares[i] = val
			i++
		}

		if num, err := strconv.Atoi(string(char)); err == nil {
			i += num
		}
	}

	return board
}

func (board Chessboard) GetPieceLocations(piece_type Piece) (bb Bitboard) {
	bb = Bitboard{}

	for i, piece := range board.Squares {
		if piece == piece_type {
			bb.Board[i] = true
		}
	}

	return
}

func (board Chessboard) String() string {
	output := ""
	for i, piece := range board.Squares {
		output += " " + piece_symbols[piece]

		if (i+1)%8 == 0 && i != 0 {
			output += "\n"
		}
	}
	return output
}
