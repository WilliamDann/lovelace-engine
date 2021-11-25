package main

type Piece int

const (
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

	none Piece = iota

	// TODO maybe not have a None piece type?
	//   This is for when get_piece_at does not find a piece
	//   Maybe the behavior of that function should be different instead
	//   notice that none is not in piece_types for this reason
	//
	//   this might just mean that a square centered board representation is better, rather than a
	//   piece centered one like this
	//       adding none to piece_symbols as '_' would seem to make sense
	//
	//   Because we use bitboards you cannot simply place a none piece to clear a square so this is misleading
	//   or we could call clear_square if that is attempted in place_piece
	//
	//   This is a decision that still needs to be made either way:
	//        - should a Piece of type none exist to represent an empty square or should
	//          something like nil be used
)

var piece_types []Piece = []Piece{white_pawn, white_rook, white_bishop, white_knight, white_queen, white_king, black_pawn, black_rook, black_bishop, black_knight, black_queen, black_king}
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
}

type Chessboard struct {
	pieces map[Piece]Bitboard
}

func NewChessboard() Chessboard {
	return Chessboard{make(map[Piece]Bitboard)}
}

func (board Chessboard) place(piece Piece, square int) Chessboard {
	copy := board.pieces[piece].board
	copy[square] = true

	board.pieces[piece] = Bitboard{copy}
	return board
}

func (board Chessboard) get_piece_at(square int) Piece {
	for _, piece_type := range piece_types {
		if board.pieces[piece_type].board[square] {
			return piece_type
		}
	}

	return none
}

func (board Chessboard) clear(square int) Chessboard {
	piece_type := board.get_piece_at(square)
	copy := board.pieces[piece_type].board

	copy[square] = false
	board.pieces[piece_type] = Bitboard{copy}

	return board
}

func (board Chessboard) String() string {
	output := ""

	for i := 0; i < 64; i++ {
		found := false
		for _, piece_type := range piece_types {
			if board.pieces[piece_type].board[i] {
				output += " " + piece_symbols[piece_type]
				found = true
				break
			}
		}

		if !found {
			output += " _"
		}

		if (i+1)%8 == 0 && i != 0 {
			output += "\n"
		}
	}

	return output
}
