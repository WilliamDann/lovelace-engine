package main

import "strconv"

type Move struct {
	from int
	to   int

	promote Piece
	capture Piece
}

func (move Move) String() (str string) {
	str = strconv.FormatInt(int64(move.from), 10)

	if move.capture != none {
		str += "x"
	} else {
		str += "->"
	}

	str = strconv.FormatInt(int64(move.to), 10)

	if move.capture != none {
		str += "=" + piece_symbols[move.promote]
	}

	return
}

func NewMove(from, to int) Move {
	return Move{from, to, none, none}
}

type Game struct {
	board       Chessboard
	whiteToPlay bool
	moves       []Move
}

func (game Game) PushMove(move Move) Game {
	if move.capture != none {
		game.board.squares[move.to] = none
	}

	game.board.squares[move.to], game.board.squares[move.from] = game.board.squares[move.from], game.board.squares[move.to]
	game.moves = append(game.moves, move)

	game.whiteToPlay = !game.whiteToPlay

	return game
}

func (game Game) String() string {
	return game.board.String()
}
