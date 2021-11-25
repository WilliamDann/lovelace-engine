package main

type Move struct {
	from int
	to   int

	promote Piece
	capture bool
}

func NewMove(from, to int) Move {
	return Move{from, to, none, false}
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
					Move{square, PawnStep(square, white).to, white_queen, false},
					Move{square, PawnStep(square, white).to, white_knight, false},
					Move{square, PawnStep(square, white).to, white_bishop, false},
					Move{square, PawnStep(square, white).to, white_rook, false},
				)

			} else {
				// single move
				movements = append(movements, PawnStep(square, white))
			}

			// captues
			if square+captureRight >= 0 && square+captureRight < 64 && game.board.squares[square+captureRight] != none {
				captures = append(captures, Move{square, square + captureRight, none, true})
			}
			if square+captureLeft >= 0 && square+captureLeft < 64 && game.board.squares[square+captureLeft] != none {
				captures = append(captures, Move{square, square + captureLeft, none, true})
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
				Move{move.from, move.to, white_queen, true},
				Move{move.from, move.to, white_knight, true},
				Move{move.from, move.to, white_bishop, true},
				Move{move.from, move.to, white_rook, true},
			)
		} else {
			moves = append(moves, move)
		}
	}

	return
}
