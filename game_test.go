package main

import (
	"testing"
)

func TestGetWhitePawnMovesUnmoved(t *testing.T) {
	board := NewChessboard()
	index := CoordsToSquareIndex(0, 1)

	board.squares[index] = white_pawn
	game := Game{board, true, []Move{}}

	moves := game.GetPawnMoves(true)

	if len(moves) != 2 {
		t.Errorf("Extra moves generated")
	}
	if moves[0].to != CoordsToSquareIndex(0, 2) {
		t.Errorf("Pawn cannot be moved up one y position")
	}
	if moves[1].to != CoordsToSquareIndex(0, 3) {
		t.Errorf("Pawn cannot be moved up two y positions")
	}
}

func TestGetWhitePawnMovesMoved(t *testing.T) {
	board := NewChessboard()
	index := CoordsToSquareIndex(0, 2)

	board.squares[index] = white_pawn
	game := Game{board, true, []Move{}}

	moves := game.GetPawnMoves(true)

	if len(moves) != 1 {
		t.Errorf("Pawn can be moved twice after leaving starting square")
	}
}

func TestGetWhitePawnPromote(t *testing.T) {
	board := NewChessboard()
	index := CoordsToSquareIndex(0, 6)

	board.squares[index] = white_pawn
	game := Game{board, true, []Move{}}

	moves := game.GetPawnMoves(true)

	if len(moves) != 4 {
		t.Errorf("Pawn cannot promote!")
	}
}

func TestGetWhitePawnMovesUnmovedBlockedClose(t *testing.T) {
	board := NewChessboard()
	index := CoordsToSquareIndex(0, 1)

	board.squares[index] = white_pawn
	board.squares[index+8] = black_pawn

	game := Game{board, true, []Move{}}

	moves := game.GetPawnMoves(true)

	if len(moves) != 0 {
		t.Errorf("Pawn can be moved when blocked")
	}
}

func TestGetWhitePawnMovesUnmovedBlockedFar(t *testing.T) {
	board := NewChessboard()
	index := CoordsToSquareIndex(0, 1)

	board.squares[index] = white_pawn
	board.squares[index+16] = black_pawn

	game := Game{board, true, []Move{}}

	moves := game.GetPawnMoves(true)

	if len(moves) != 1 {
		t.Errorf("Pawn can be moved when blocked")
	}
}

func TestGetWhitePawnMovesCaptures(t *testing.T) {
	board := NewChessboard()
	index1 := CoordsToSquareIndex(1, 1)
	index2 := index1 + 8 + 1
	index3 := index1 + 8 - 1

	board.squares[index1] = white_pawn
	board.squares[index2] = black_pawn
	board.squares[index3] = black_pawn

	game := Game{board, true, []Move{}}

	moves := game.GetPawnMoves(true)

	if len(moves) != 4 {
		t.Errorf("Pawn cannot be captured")
	}
}

func TestGetBlackPawnMovesUnmoved(t *testing.T) {
	board := NewChessboard()
	index := CoordsToSquareIndex(0, 6)

	board.squares[index] = black_pawn
	game := Game{board, false, []Move{}}

	moves := game.GetPawnMoves(false)

	if moves[0].to != CoordsToSquareIndex(0, 5) {
		t.Errorf("Pawn cannot be moved up one y position")
	}
	if moves[1].to != CoordsToSquareIndex(0, 4) {
		t.Errorf("Pawn cannot be moved up two y positions")
	}
}

func TestGetBlackPawnMovesMoved(t *testing.T) {
	board := NewChessboard()
	index := CoordsToSquareIndex(0, 5)

	board.squares[index] = black_pawn
	game := Game{board, false, []Move{}}

	moves := game.GetPawnMoves(false)

	if len(moves) != 1 {
		t.Errorf("Pawn can be moved twice after leaving starting square")
	}
}

func TestGetBlackPawnPromote(t *testing.T) {
	board := NewChessboard()
	index := CoordsToSquareIndex(0, 1)

	board.squares[index] = black_pawn
	game := Game{board, false, []Move{}}

	moves := game.GetPawnMoves(false)

	if len(moves) != 4 {
		t.Errorf("Pawn cannot promote!")
	}
}

func TestGetBlackPawnMovesUnmovedBlockedClose(t *testing.T) {
	board := NewChessboard()
	index := CoordsToSquareIndex(0, 7)

	board.squares[index] = black_pawn
	board.squares[index-8] = white_pawn

	game := Game{board, false, []Move{}}

	moves := game.GetPawnMoves(false)

	if len(moves) != 0 {
		t.Errorf("Pawn can be moved when blocked")
	}
}

func TestGetBlackPawnMovesUnmovedBlockedFar(t *testing.T) {
	board := NewChessboard()
	index := CoordsToSquareIndex(0, 7)

	board.squares[index] = black_pawn
	board.squares[index-16] = white_pawn

	game := Game{board, false, []Move{}}

	moves := game.GetPawnMoves(false)

	if len(moves) != 1 {
		t.Errorf("Pawn can be moved when blocked")
	}
}
