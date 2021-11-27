package main

func PawnStep(square int, white bool) Move {
	if white {
		return NewMove(square, square+8)
	}
	return NewMove(square, square-8)
}

func PawnCaptures(square int, white bool) (moves []Move) {
	captureRight := 8 + 1
	captureLeft := 8 - 1

	if !white {
		captureRight = -8 + 1
		captureLeft = -8 - 1
	}

	moves = append(moves, Move{square, square + captureRight, none, none})
	moves = append(moves, Move{square, square + captureLeft, none, none})

	return
}

// TODO captures en passant
func (game Game) GetPawnMoves() (moves []Move) {
	movements := []Move{}
	captures := []Move{}

	seeking := white_pawn

	startRankLow := 8
	startRankHigh := 15

	promoteRankLow := 48
	promoteRankHigh := 55

	if !game.whiteToPlay {
		seeking = black_pawn
		startRankLow, startRankHigh, promoteRankLow, promoteRankHigh = promoteRankLow, promoteRankHigh, startRankLow, startRankHigh
	}

	for square, pawn := range game.board.GetPieceLocations(seeking).board {
		if pawn {
			// moves
			if square >= startRankLow && square <= startRankHigh {
				// double move
				movements = append(movements, PawnStep(square, game.whiteToPlay))
				movements = append(movements, NewMove(square, PawnStep(PawnStep(square, game.whiteToPlay).to, game.whiteToPlay).to))
			} else if square >= promoteRankLow && square <= promoteRankHigh {
				// promotion move
				if game.whiteToPlay {
					movements = append(movements,
						Move{square, PawnStep(square, game.whiteToPlay).to, white_queen, none},
						Move{square, PawnStep(square, game.whiteToPlay).to, white_knight, none},
						Move{square, PawnStep(square, game.whiteToPlay).to, white_bishop, none},
						Move{square, PawnStep(square, game.whiteToPlay).to, white_rook, none},
					)
				} else {
					movements = append(movements,
						Move{square, PawnStep(square, game.whiteToPlay).to, black_queen, none},
						Move{square, PawnStep(square, game.whiteToPlay).to, black_knight, none},
						Move{square, PawnStep(square, game.whiteToPlay).to, black_bishop, none},
						Move{square, PawnStep(square, game.whiteToPlay).to, black_rook, none},
					)
				}

			} else {
				// single move
				movements = append(movements, PawnStep(square, game.whiteToPlay))
			}

			// captues
			captures = append(captures, PawnCaptures(square, game.whiteToPlay)...)
		}
	}

	for _, move := range movements {
		if (move.to > 63 || move.to < 0) || game.board.squares[move.to] != none {
			break
		}
		moves = append(moves, move)
	}

	for _, move := range captures {
		if (move.to < 0 || move.to >= 64) || game.board.squares[move.to] == none {
			continue
		}

		if move.from >= promoteRankLow && move.from <= promoteRankHigh {
			if game.whiteToPlay {
				moves = append(moves,
					Move{move.from, move.to, white_queen, game.board.squares[move.to]},
					Move{move.from, move.to, white_knight, game.board.squares[move.to]},
					Move{move.from, move.to, white_bishop, game.board.squares[move.to]},
					Move{move.from, move.to, white_rook, game.board.squares[move.to]},
				)
			} else {
				moves = append(moves,
					Move{move.from, move.to, black_queen, game.board.squares[move.to]},
					Move{move.from, move.to, black_knight, game.board.squares[move.to]},
					Move{move.from, move.to, black_bishop, game.board.squares[move.to]},
					Move{move.from, move.to, black_rook, game.board.squares[move.to]},
				)

			}
		} else {
			moves = append(moves, Move{move.from, move.to, none, game.board.squares[move.to]})
		}
	}

	return
}
