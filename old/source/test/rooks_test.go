package tests

import (
	. "lovelace/source"
	"testing"
)

func TestRookMoves(t *testing.T) {
	board := NewChessboard()
	expected := []Move{
		{From: CoordsToSquareIndex(0, 0), To: CoordsToSquareIndex(0, 1), Capture: None, Promote: None},
		{From: CoordsToSquareIndex(0, 0), To: CoordsToSquareIndex(0, 2), Capture: None, Promote: None},
		{From: CoordsToSquareIndex(0, 0), To: CoordsToSquareIndex(0, 3), Capture: None, Promote: None},
		{From: CoordsToSquareIndex(0, 0), To: CoordsToSquareIndex(0, 4), Capture: None, Promote: None},
		{From: CoordsToSquareIndex(0, 0), To: CoordsToSquareIndex(0, 5), Capture: None, Promote: None},
		{From: CoordsToSquareIndex(0, 0), To: CoordsToSquareIndex(0, 6), Capture: None, Promote: None},
		{From: CoordsToSquareIndex(0, 0), To: CoordsToSquareIndex(0, 7), Capture: None, Promote: None},

		{From: CoordsToSquareIndex(0, 0), To: CoordsToSquareIndex(1, 0), Capture: None, Promote: None},
		{From: CoordsToSquareIndex(0, 0), To: CoordsToSquareIndex(2, 0), Capture: None, Promote: None},
		{From: CoordsToSquareIndex(0, 0), To: CoordsToSquareIndex(3, 0), Capture: None, Promote: None},
		{From: CoordsToSquareIndex(0, 0), To: CoordsToSquareIndex(4, 0), Capture: None, Promote: None},
		{From: CoordsToSquareIndex(0, 0), To: CoordsToSquareIndex(5, 0), Capture: None, Promote: None},
		{From: CoordsToSquareIndex(0, 0), To: CoordsToSquareIndex(6, 0), Capture: None, Promote: None},
		{From: CoordsToSquareIndex(0, 0), To: CoordsToSquareIndex(7, 0), Capture: None, Promote: None},
	}

	board.Squares[CoordsToSquareIndex(0, 0)] = Black_rook

	game := Game{Board: board, WhiteToPlay: false, Moves: []Move{}}
	moves := game.GetRookMoves()

	if len(expected) != len(moves) {
		t.Errorf("Incorrect amount of moves generated")
	}

	for i, move := range expected {
		AsertMoveEquals(t, moves[i], move)
	}
}

func TestRookCaptures(t *testing.T) {
	board := NewChessboard()
	expected := []Move{
		{From: CoordsToSquareIndex(0, 0), To: CoordsToSquareIndex(0, 1), Capture: None, Promote: None},
		{From: CoordsToSquareIndex(0, 0), To: CoordsToSquareIndex(0, 2), Capture: White_bishop, Promote: None},

		{From: CoordsToSquareIndex(0, 0), To: CoordsToSquareIndex(1, 0), Capture: None, Promote: None},
		{From: CoordsToSquareIndex(0, 0), To: CoordsToSquareIndex(2, 0), Capture: White_bishop, Promote: None},
	}

	board.Squares[CoordsToSquareIndex(0, 0)] = Black_rook
	board.Squares[CoordsToSquareIndex(0, 2)] = White_bishop
	board.Squares[CoordsToSquareIndex(2, 0)] = White_bishop

	game := Game{Board: board, WhiteToPlay: false, Moves: []Move{}}
	moves := game.GetRookMoves()

	if len(expected) != len(moves) {
		t.Errorf("Incorrect amount of moves generated")
	}

	for i, move := range expected {
		AsertMoveEquals(t, moves[i], move)
	}
}
