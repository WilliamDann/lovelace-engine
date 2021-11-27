package main

import "testing"

func TestRookMoves(t *testing.T) {
	board := NewChessboard()
	expected := []Move{
		{CoordsToSquareIndex(0, 0), CoordsToSquareIndex(0, 1), none, none},
		{CoordsToSquareIndex(0, 0), CoordsToSquareIndex(0, 2), none, none},
		{CoordsToSquareIndex(0, 0), CoordsToSquareIndex(0, 3), none, none},
		{CoordsToSquareIndex(0, 0), CoordsToSquareIndex(0, 4), none, none},
		{CoordsToSquareIndex(0, 0), CoordsToSquareIndex(0, 5), none, none},
		{CoordsToSquareIndex(0, 0), CoordsToSquareIndex(0, 6), none, none},
		{CoordsToSquareIndex(0, 0), CoordsToSquareIndex(0, 7), none, none},

		{CoordsToSquareIndex(0, 0), CoordsToSquareIndex(1, 0), none, none},
		{CoordsToSquareIndex(0, 0), CoordsToSquareIndex(2, 0), none, none},
		{CoordsToSquareIndex(0, 0), CoordsToSquareIndex(3, 0), none, none},
		{CoordsToSquareIndex(0, 0), CoordsToSquareIndex(4, 0), none, none},
		{CoordsToSquareIndex(0, 0), CoordsToSquareIndex(5, 0), none, none},
		{CoordsToSquareIndex(0, 0), CoordsToSquareIndex(6, 0), none, none},
		{CoordsToSquareIndex(0, 0), CoordsToSquareIndex(7, 0), none, none},
	}

	board.squares[CoordsToSquareIndex(0, 0)] = black_rook

	game := Game{board, false, []Move{}}
	moves := game.GetRookMoves()

	if len(expected) != len(moves) {
		t.Errorf("Incorrect amount of moves generated")
	}

	for i, move := range expected {
		AsertMoveEquals(t, moves[i], move)
	}
}
