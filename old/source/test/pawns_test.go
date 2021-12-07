package tests

import (
	. "lovelace/source"
	"testing"
)

func TestWhitePawnsOn2(t *testing.T) {
	board := NewChessboard()
	expected := []Move{
		{From: CoordsToSquareIndex(0, 1), To: CoordsToSquareIndex(0, 1) + 8, Capture: None, Promote: None},
		{From: CoordsToSquareIndex(0, 1), To: CoordsToSquareIndex(0, 1) + 16, Capture: None, Promote: None},
	}

	board.Squares[CoordsToSquareIndex(0, 1)] = White_pawn

	game := Game{Board: board, WhiteToPlay: true, Moves: []Move{}}
	moves := game.GetPawnMoves()

	if len(expected) != len(moves) {
		t.Errorf("Incorrect amount of moves generated")
	}

	for i, move := range expected {
		AsertMoveEquals(t, moves[i], move)
	}
}

func TestBlackPawnsOn2(t *testing.T) {
	board := NewChessboard()
	expected := []Move{
		{From: CoordsToSquareIndex(0, 6), To: CoordsToSquareIndex(0, 6) - 8, Capture: None, Promote: None},
		{From: CoordsToSquareIndex(0, 6), To: CoordsToSquareIndex(0, 6) - 16, Capture: None, Promote: None},
	}

	board.Squares[CoordsToSquareIndex(0, 6)] = Black_pawn

	game := Game{Board: board, WhiteToPlay: false, Moves: []Move{}}
	moves := game.GetPawnMoves()

	if len(expected) != len(moves) {
		t.Errorf("Incorrect amount of moves generated")
	}

	for i, move := range expected {
		AsertMoveEquals(t, moves[i], move)
	}
}

func TestWhitePawnsOff2(t *testing.T) {
	board := NewChessboard()
	expected := []Move{
		{From: CoordsToSquareIndex(0, 2), To: CoordsToSquareIndex(0, 2) + 8, Capture: None, Promote: None},
	}

	board.Squares[CoordsToSquareIndex(0, 2)] = White_pawn

	game := Game{Board: board, WhiteToPlay: true, Moves: []Move{}}
	moves := game.GetPawnMoves()

	if len(expected) != len(moves) {
		t.Errorf("Incorrect amount of moves generated")
	}

	for i, move := range expected {
		AsertMoveEquals(t, moves[i], move)
	}
}

func TestBlackPawnsOff2(t *testing.T) {
	board := NewChessboard()
	expected := []Move{
		{From: CoordsToSquareIndex(0, 5), To: CoordsToSquareIndex(0, 5) - 8, Capture: None, Promote: None},
	}

	board.Squares[CoordsToSquareIndex(0, 5)] = Black_pawn

	game := Game{Board: board, WhiteToPlay: false, Moves: []Move{}}
	moves := game.GetPawnMoves()

	if len(expected) != len(moves) {
		t.Errorf("Incorrect amount of moves generated")
	}

	for i, move := range expected {
		AsertMoveEquals(t, moves[i], move)
	}
}

func TestWhitePawnsPromote(t *testing.T) {
	board := NewChessboard()
	expected := []Move{
		{From: CoordsToSquareIndex(0, 6), To: CoordsToSquareIndex(0, 6) + 8, Promote: White_queen, Capture: None},
		{From: CoordsToSquareIndex(0, 6), To: CoordsToSquareIndex(0, 6) + 8, Promote: White_knight, Capture: None},
		{From: CoordsToSquareIndex(0, 6), To: CoordsToSquareIndex(0, 6) + 8, Promote: White_bishop, Capture: None},
		{From: CoordsToSquareIndex(0, 6), To: CoordsToSquareIndex(0, 6) + 8, Promote: White_rook, Capture: None},
	}

	board.Squares[CoordsToSquareIndex(0, 6)] = White_pawn

	game := Game{Board: board, WhiteToPlay: true, Moves: []Move{}}
	moves := game.GetPawnMoves()

	if len(expected) != len(moves) {
		t.Errorf("Incorrect amount of moves generated")
	}

	for i, move := range expected {
		AsertMoveEquals(t, moves[i], move)
	}
}

func TestBlackPawnsPromote(t *testing.T) {
	board := NewChessboard()
	expected := []Move{
		{From: CoordsToSquareIndex(0, 1), To: CoordsToSquareIndex(0, 1) - 8, Promote: Black_queen, Capture: None},
		{From: CoordsToSquareIndex(0, 1), To: CoordsToSquareIndex(0, 1) - 8, Promote: Black_knight, Capture: None},
		{From: CoordsToSquareIndex(0, 1), To: CoordsToSquareIndex(0, 1) - 8, Promote: Black_bishop, Capture: None},
		{From: CoordsToSquareIndex(0, 1), To: CoordsToSquareIndex(0, 1) - 8, Promote: Black_rook, Capture: None},
	}

	board.Squares[CoordsToSquareIndex(0, 1)] = Black_pawn

	game := Game{Board: board, WhiteToPlay: false, Moves: []Move{}}
	moves := game.GetPawnMoves()

	if len(expected) != len(moves) {
		t.Errorf("Incorrect amount of moves generated")
	}

	for i, move := range expected {
		AsertMoveEquals(t, moves[i], move)
	}
}

func TestWhitePawnsCapture(t *testing.T) {
	board := NewChessboard()
	expected := []Move{
		{From: CoordsToSquareIndex(1, 1), To: CoordsToSquareIndex(1, 1) + 8, Promote: None, Capture: None},
		{From: CoordsToSquareIndex(1, 1), To: CoordsToSquareIndex(1, 1) + 16, Promote: None, Capture: None},
		{From: CoordsToSquareIndex(1, 1), To: CoordsToSquareIndex(1, 1) + 8 + 1, Promote: None, Capture: Black_pawn},
		{From: CoordsToSquareIndex(1, 1), To: CoordsToSquareIndex(1, 1) + 8 - 1, Promote: None, Capture: Black_pawn},
	}

	board.Squares[CoordsToSquareIndex(1, 1)] = White_pawn
	board.Squares[CoordsToSquareIndex(1, 1)+8+1] = Black_pawn
	board.Squares[CoordsToSquareIndex(1, 1)+8-1] = Black_pawn

	game := Game{Board: board, WhiteToPlay: true, Moves: []Move{}}
	moves := game.GetPawnMoves()

	if len(expected) != len(moves) {
		t.Errorf("Incorrect amount of moves generated")
	}

	for i, move := range expected {
		AsertMoveEquals(t, moves[i], move)
	}
}

func TestBlackPawnsCapture(t *testing.T) {
	board := NewChessboard()
	expected := []Move{
		{From: CoordsToSquareIndex(1, 6), To: CoordsToSquareIndex(1, 6) - 8, Promote: None, Capture: None},
		{From: CoordsToSquareIndex(1, 6), To: CoordsToSquareIndex(1, 6) - 16, Promote: None, Capture: None},
		{From: CoordsToSquareIndex(1, 6), To: CoordsToSquareIndex(1, 6) - 8 + 1, Promote: None, Capture: White_pawn},
		{From: CoordsToSquareIndex(1, 6), To: CoordsToSquareIndex(1, 6) - 8 - 1, Promote: None, Capture: White_pawn},
	}

	board.Squares[CoordsToSquareIndex(1, 6)] = Black_pawn
	board.Squares[CoordsToSquareIndex(1, 6)-8+1] = White_pawn
	board.Squares[CoordsToSquareIndex(1, 6)-8-1] = White_pawn

	game := Game{Board: board, WhiteToPlay: false, Moves: []Move{}}
	moves := game.GetPawnMoves()

	if len(expected) != len(moves) {
		t.Errorf("Incorrect amount of moves generated")
	}

	for i, move := range expected {
		AsertMoveEquals(t, moves[i], move)
	}
}

func TestWhitePawnsCapturePromote(t *testing.T) {
	board := NewChessboard()
	expected := []Move{
		{From: CoordsToSquareIndex(1, 6), To: CoordsToSquareIndex(1, 6) + 8, Promote: White_queen, Capture: None},
		{From: CoordsToSquareIndex(1, 6), To: CoordsToSquareIndex(1, 6) + 8, Promote: White_knight, Capture: None},
		{From: CoordsToSquareIndex(1, 6), To: CoordsToSquareIndex(1, 6) + 8, Promote: White_bishop, Capture: None},
		{From: CoordsToSquareIndex(1, 6), To: CoordsToSquareIndex(1, 6) + 8, Promote: White_rook, Capture: None},

		{From: CoordsToSquareIndex(1, 6), To: CoordsToSquareIndex(1, 6) + 8 + 1, Promote: White_queen, Capture: Black_pawn},
		{From: CoordsToSquareIndex(1, 6), To: CoordsToSquareIndex(1, 6) + 8 + 1, Promote: White_knight, Capture: Black_pawn},
		{From: CoordsToSquareIndex(1, 6), To: CoordsToSquareIndex(1, 6) + 8 + 1, Promote: White_bishop, Capture: Black_pawn},
		{From: CoordsToSquareIndex(1, 6), To: CoordsToSquareIndex(1, 6) + 8 + 1, Promote: White_rook, Capture: Black_pawn},

		{From: CoordsToSquareIndex(1, 6), To: CoordsToSquareIndex(1, 6) + 8 - 1, Promote: White_queen, Capture: Black_pawn},
		{From: CoordsToSquareIndex(1, 6), To: CoordsToSquareIndex(1, 6) + 8 - 1, Promote: White_knight, Capture: Black_pawn},
		{From: CoordsToSquareIndex(1, 6), To: CoordsToSquareIndex(1, 6) + 8 - 1, Promote: White_bishop, Capture: Black_pawn},
		{From: CoordsToSquareIndex(1, 6), To: CoordsToSquareIndex(1, 6) + 8 - 1, Promote: White_rook, Capture: Black_pawn},
	}

	board.Squares[CoordsToSquareIndex(1, 6)] = White_pawn
	board.Squares[CoordsToSquareIndex(1, 6)+8+1] = Black_pawn
	board.Squares[CoordsToSquareIndex(1, 6)+8-1] = Black_pawn

	game := Game{Board: board, WhiteToPlay: true, Moves: []Move{}}
	moves := game.GetPawnMoves()

	if len(expected) != len(moves) {
		t.Errorf("Incorrect amount of moves generated")
	}

	for i, move := range expected {
		AsertMoveEquals(t, moves[i], move)
	}
}

func TestBlackPawnsCapturePrompte(t *testing.T) {
	board := NewChessboard()
	expected := []Move{
		{From: CoordsToSquareIndex(1, 1), To: CoordsToSquareIndex(1, 1) - 8, Promote: Black_queen, Capture: None},
		{From: CoordsToSquareIndex(1, 1), To: CoordsToSquareIndex(1, 1) - 8, Promote: Black_knight, Capture: None},
		{From: CoordsToSquareIndex(1, 1), To: CoordsToSquareIndex(1, 1) - 8, Promote: Black_bishop, Capture: None},
		{From: CoordsToSquareIndex(1, 1), To: CoordsToSquareIndex(1, 1) - 8, Promote: Black_rook, Capture: None},

		{From: CoordsToSquareIndex(1, 1), To: CoordsToSquareIndex(1, 1) - 8 + 1, Promote: Black_queen, Capture: White_pawn},
		{From: CoordsToSquareIndex(1, 1), To: CoordsToSquareIndex(1, 1) - 8 + 1, Promote: Black_knight, Capture: White_pawn},
		{From: CoordsToSquareIndex(1, 1), To: CoordsToSquareIndex(1, 1) - 8 + 1, Promote: Black_bishop, Capture: White_pawn},
		{From: CoordsToSquareIndex(1, 1), To: CoordsToSquareIndex(1, 1) - 8 + 1, Promote: Black_rook, Capture: White_pawn},

		{From: CoordsToSquareIndex(1, 1), To: CoordsToSquareIndex(1, 1) - 8 - 1, Promote: Black_queen, Capture: White_pawn},
		{From: CoordsToSquareIndex(1, 1), To: CoordsToSquareIndex(1, 1) - 8 - 1, Promote: Black_knight, Capture: White_pawn},
		{From: CoordsToSquareIndex(1, 1), To: CoordsToSquareIndex(1, 1) - 8 - 1, Promote: Black_bishop, Capture: White_pawn},
		{From: CoordsToSquareIndex(1, 1), To: CoordsToSquareIndex(1, 1) - 8 - 1, Promote: Black_rook, Capture: White_pawn},
	}

	board.Squares[CoordsToSquareIndex(1, 1)] = Black_pawn
	board.Squares[CoordsToSquareIndex(1, 1)-8+1] = White_pawn
	board.Squares[CoordsToSquareIndex(1, 1)-8-1] = White_pawn

	game := Game{Board: board, WhiteToPlay: false, Moves: []Move{}}
	moves := game.GetPawnMoves()

	if len(expected) != len(moves) {
		t.Errorf("Incorrect amount of moves generated")
	}

	for i, move := range expected {
		AsertMoveEquals(t, moves[i], move)
	}
}

func TestWhitePawnsCaptureEnPassant(t *testing.T) {
	t.Errorf("Not implemented")
}

func TestBlackPawnsCaptureEnPassant(t *testing.T) {
	t.Errorf("Not implemented")
}
