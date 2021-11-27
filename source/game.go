package source

import "strconv"

type Move struct {
	From int
	To   int

	Promote Piece
	Capture Piece
}

func (move Move) String() (str string) {
	str = strconv.FormatInt(int64(move.From), 10)

	if move.Capture != None {
		str += "x"
	} else {
		str += "->"
	}

	str = strconv.FormatInt(int64(move.To), 10)

	if move.Capture != None {
		str += "=" + piece_symbols[move.Promote]
	}

	return
}

func NewMove(from, to int) Move {
	return Move{from, to, None, None}
}

type Game struct {
	Board       Chessboard
	WhiteToPlay bool
	Moves       []Move
}

func (game Game) PushMove(move Move) Game {
	if move.Capture != None {
		game.Board.Squares[move.To] = None
	}

	game.Board.Squares[move.To], game.Board.Squares[move.From] = game.Board.Squares[move.From], game.Board.Squares[move.To]
	game.Moves = append(game.Moves, move)

	game.WhiteToPlay = !game.WhiteToPlay

	return game
}

func (game Game) String() string {
	return game.Board.String()
}
