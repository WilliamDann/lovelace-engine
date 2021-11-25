package main

type Move struct {
	from int
	to   int

	promote Piece
	capture Piece
}

func NewMove(from, to int) Move {
	return Move{from, to, none, none}
}

func PawnStep(square int, white bool) Move {
	if white {
		return NewMove(square, square+8)
	}
	return NewMove(square, square-8)
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

	return game
}

// TODO captures en passant
func (game Game) GetPawnMoves(white bool) (moves []Move) {
	movements := []Move{}
	captures := []Move{}

	seeking := white_pawn

	startRankLow := 8
	startRankHigh := 15

	promoteRankLow := 48
	promoteRankHigh := 55

	captureRight := 8 + 1
	captureLeft := 8 - 1

	if !white {
		seeking = black_pawn
		startRankLow, startRankHigh, promoteRankLow, promoteRankHigh = promoteRankLow, promoteRankHigh, startRankLow, startRankHigh

		captureRight = -8 + 1
		captureLeft = -8 - 1
	}

	for square, pawn := range game.board.GetPieceLocations(seeking).board {
		if pawn {
			// moves
			if square >= startRankLow && square <= startRankHigh {
				// double move
				movements = append(movements, PawnStep(square, white))
				movements = append(movements, NewMove(square, PawnStep(PawnStep(square, white).to, white).to))
			} else if square >= promoteRankLow && square <= promoteRankHigh {
				// promotion move
				movements = append(movements,
					Move{square, PawnStep(square, white).to, white_queen, none},
					Move{square, PawnStep(square, white).to, white_knight, none},
					Move{square, PawnStep(square, white).to, white_bishop, none},
					Move{square, PawnStep(square, white).to, white_rook, none},
				)

			} else {
				// single move
				movements = append(movements, PawnStep(square, white))
			}

			// captues
			if square+captureRight >= 0 && square+captureRight < 64 && game.board.squares[square+captureRight] != none {
				captures = append(captures, Move{square, square + captureRight, none, game.board.squares[square+captureRight]})
			}
			if square+captureLeft >= 0 && square+captureLeft < 64 && game.board.squares[square+captureLeft] != none {
				captures = append(captures, Move{square, square + captureLeft, none, game.board.squares[square+captureLeft]})
			}
		}
	}

	for _, move := range movements {
		if (move.to > 63 || move.to < 0) || game.board.squares[move.to] != none {
			break
		}
		moves = append(moves, move)
	}

	for _, move := range captures {
		if move.to >= promoteRankLow && move.to <= promoteRankHigh {
			moves = append(moves,
				Move{move.from, move.to, white_queen, move.capture},
				Move{move.from, move.to, white_knight, move.capture},
				Move{move.from, move.to, white_bishop, move.capture},
				Move{move.from, move.to, white_rook, move.capture},
			)
		} else {
			moves = append(moves, move)
		}
	}

	return
}

func (game Game) String() string {
	return game.board.String()
}
