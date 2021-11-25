package main

import "testing"

// combining these tests is good enough for me
// if you want to make it better be my guest

func TestPlaceAndGetPieceAt0(t *testing.T) {
	board := NewChessboard()
	board.place(white_pawn, 0)

	if board.get_piece_at(0) != white_pawn {
		t.Errorf("Piece not at placed square")
	}
}

func TestPlaceAndGetPieceAt1(t *testing.T) {
	board := NewChessboard()
	board.place(white_pawn, 1)

	if board.get_piece_at(1) != white_pawn {
		t.Errorf("Piece not at placed square")
	}
}

func TestPlaceAndGetPieceAt63(t *testing.T) {
	board := NewChessboard()
	board.place(white_pawn, 63)

	if board.get_piece_at(63) != white_pawn {
		t.Errorf("Piece not at placed square")
	}
}

func TestClear0(t *testing.T) {
	board := NewChessboard()
	board.place(white_pawn, 0)
	board.clear(0)

	if board.get_piece_at(0) != none {
		t.Errorf("Square not cleared")
	}
}

func TestClear1(t *testing.T) {
	board := NewChessboard()
	board.place(white_pawn, 1)
	board.clear(1)

	if board.get_piece_at(1) != none {
		t.Errorf("Square not cleared")
	}
}

func TestClear63(t *testing.T) {
	board := NewChessboard()
	board.place(white_pawn, 63)
	board.clear(63)

	if board.get_piece_at(63) != none {
		t.Errorf("Square not cleared")
	}
}
