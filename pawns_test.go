package main

import "testing"

func TestWhitePawnsOn2(t *testing.T) {
	board := NewChessboard()
	expected := []Move{
		{CoordsToSquareIndex(0, 1), CoordsToSquareIndex(0, 1) + 8, none, none},
		{CoordsToSquareIndex(0, 1), CoordsToSquareIndex(0, 1) + 16, none, none},
	}

	board.squares[CoordsToSquareIndex(0, 1)] = white_pawn

	game := Game{board, true, []Move{}}
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
		{CoordsToSquareIndex(0, 6), CoordsToSquareIndex(0, 6) - 8, none, none},
		{CoordsToSquareIndex(0, 6), CoordsToSquareIndex(0, 6) - 16, none, none},
	}

	board.squares[CoordsToSquareIndex(0, 6)] = black_pawn

	game := Game{board, false, []Move{}}
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
		{CoordsToSquareIndex(0, 2), CoordsToSquareIndex(0, 2) + 8, none, none},
	}

	board.squares[CoordsToSquareIndex(0, 2)] = white_pawn

	game := Game{board, true, []Move{}}
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
		{CoordsToSquareIndex(0, 5), CoordsToSquareIndex(0, 5) - 8, none, none},
	}

	board.squares[CoordsToSquareIndex(0, 5)] = black_pawn

	game := Game{board, false, []Move{}}
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
		{CoordsToSquareIndex(0, 6), CoordsToSquareIndex(0, 6) + 8, white_queen, none},
		{CoordsToSquareIndex(0, 6), CoordsToSquareIndex(0, 6) + 8, white_knight, none},
		{CoordsToSquareIndex(0, 6), CoordsToSquareIndex(0, 6) + 8, white_bishop, none},
		{CoordsToSquareIndex(0, 6), CoordsToSquareIndex(0, 6) + 8, white_rook, none},
	}

	board.squares[CoordsToSquareIndex(0, 6)] = white_pawn

	game := Game{board, true, []Move{}}
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
		{CoordsToSquareIndex(0, 1), CoordsToSquareIndex(0, 1) - 8, black_queen, none},
		{CoordsToSquareIndex(0, 1), CoordsToSquareIndex(0, 1) - 8, black_knight, none},
		{CoordsToSquareIndex(0, 1), CoordsToSquareIndex(0, 1) - 8, black_bishop, none},
		{CoordsToSquareIndex(0, 1), CoordsToSquareIndex(0, 1) - 8, black_rook, none},
	}

	board.squares[CoordsToSquareIndex(0, 1)] = black_pawn

	game := Game{board, false, []Move{}}
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
		{CoordsToSquareIndex(1, 1), CoordsToSquareIndex(1, 1) + 8, none, none},
		{CoordsToSquareIndex(1, 1), CoordsToSquareIndex(1, 1) + 16, none, none},
		{CoordsToSquareIndex(1, 1), CoordsToSquareIndex(1, 1) + 8 + 1, none, black_pawn},
		{CoordsToSquareIndex(1, 1), CoordsToSquareIndex(1, 1) + 8 - 1, none, black_pawn},
	}

	board.squares[CoordsToSquareIndex(1, 1)] = white_pawn
	board.squares[CoordsToSquareIndex(1, 1)+8+1] = black_pawn
	board.squares[CoordsToSquareIndex(1, 1)+8-1] = black_pawn

	game := Game{board, true, []Move{}}
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
		{CoordsToSquareIndex(1, 6), CoordsToSquareIndex(1, 6) - 8, none, none},
		{CoordsToSquareIndex(1, 6), CoordsToSquareIndex(1, 6) - 16, none, none},
		{CoordsToSquareIndex(1, 6), CoordsToSquareIndex(1, 6) - 8 + 1, none, white_pawn},
		{CoordsToSquareIndex(1, 6), CoordsToSquareIndex(1, 6) - 8 - 1, none, white_pawn},
	}

	board.squares[CoordsToSquareIndex(1, 6)] = black_pawn
	board.squares[CoordsToSquareIndex(1, 6)-8+1] = white_pawn
	board.squares[CoordsToSquareIndex(1, 6)-8-1] = white_pawn

	game := Game{board, false, []Move{}}
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
		{CoordsToSquareIndex(1, 6), CoordsToSquareIndex(1, 6) + 8, white_queen, none},
		{CoordsToSquareIndex(1, 6), CoordsToSquareIndex(1, 6) + 8, white_knight, none},
		{CoordsToSquareIndex(1, 6), CoordsToSquareIndex(1, 6) + 8, white_bishop, none},
		{CoordsToSquareIndex(1, 6), CoordsToSquareIndex(1, 6) + 8, white_rook, none},

		{CoordsToSquareIndex(1, 6), CoordsToSquareIndex(1, 6) + 8 + 1, white_queen, black_pawn},
		{CoordsToSquareIndex(1, 6), CoordsToSquareIndex(1, 6) + 8 + 1, white_knight, black_pawn},
		{CoordsToSquareIndex(1, 6), CoordsToSquareIndex(1, 6) + 8 + 1, white_bishop, black_pawn},
		{CoordsToSquareIndex(1, 6), CoordsToSquareIndex(1, 6) + 8 + 1, white_rook, black_pawn},

		{CoordsToSquareIndex(1, 6), CoordsToSquareIndex(1, 6) + 8 - 1, white_queen, black_pawn},
		{CoordsToSquareIndex(1, 6), CoordsToSquareIndex(1, 6) + 8 - 1, white_knight, black_pawn},
		{CoordsToSquareIndex(1, 6), CoordsToSquareIndex(1, 6) + 8 - 1, white_bishop, black_pawn},
		{CoordsToSquareIndex(1, 6), CoordsToSquareIndex(1, 6) + 8 - 1, white_rook, black_pawn},
	}

	board.squares[CoordsToSquareIndex(1, 6)] = white_pawn
	board.squares[CoordsToSquareIndex(1, 6)+8+1] = black_pawn
	board.squares[CoordsToSquareIndex(1, 6)+8-1] = black_pawn

	game := Game{board, true, []Move{}}
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
		{CoordsToSquareIndex(1, 1), CoordsToSquareIndex(1, 1) - 8, black_queen, none},
		{CoordsToSquareIndex(1, 1), CoordsToSquareIndex(1, 1) - 8, black_knight, none},
		{CoordsToSquareIndex(1, 1), CoordsToSquareIndex(1, 1) - 8, black_bishop, none},
		{CoordsToSquareIndex(1, 1), CoordsToSquareIndex(1, 1) - 8, black_rook, none},

		{CoordsToSquareIndex(1, 1), CoordsToSquareIndex(1, 1) - 8 + 1, black_queen, white_pawn},
		{CoordsToSquareIndex(1, 1), CoordsToSquareIndex(1, 1) - 8 + 1, black_knight, white_pawn},
		{CoordsToSquareIndex(1, 1), CoordsToSquareIndex(1, 1) - 8 + 1, black_bishop, white_pawn},
		{CoordsToSquareIndex(1, 1), CoordsToSquareIndex(1, 1) - 8 + 1, black_rook, white_pawn},

		{CoordsToSquareIndex(1, 1), CoordsToSquareIndex(1, 1) - 8 - 1, black_queen, white_pawn},
		{CoordsToSquareIndex(1, 1), CoordsToSquareIndex(1, 1) - 8 - 1, black_knight, white_pawn},
		{CoordsToSquareIndex(1, 1), CoordsToSquareIndex(1, 1) - 8 - 1, black_bishop, white_pawn},
		{CoordsToSquareIndex(1, 1), CoordsToSquareIndex(1, 1) - 8 - 1, black_rook, white_pawn},
	}

	board.squares[CoordsToSquareIndex(1, 1)] = black_pawn
	board.squares[CoordsToSquareIndex(1, 1)-8+1] = white_pawn
	board.squares[CoordsToSquareIndex(1, 1)-8-1] = white_pawn

	game := Game{board, false, []Move{}}
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
