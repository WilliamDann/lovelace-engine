package source

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

	moves = append(moves, Move{square, square + captureRight, None, None})
	moves = append(moves, Move{square, square + captureLeft, None, None})

	return
}

// TODO captures en passant
func (game Game) GetPawnMoves() (moves []Move) {
	movements := []Move{}
	captures := []Move{}

	seeking := White_pawn

	startRankLow := 8
	startRankHigh := 15

	promoteRankLow := 48
	promoteRankHigh := 55

	if !game.WhiteToPlay {
		seeking = Black_pawn
		startRankLow, startRankHigh, promoteRankLow, promoteRankHigh = promoteRankLow, promoteRankHigh, startRankLow, startRankHigh
	}

	for square, pawn := range game.Board.GetPieceLocations(seeking).Board {
		if pawn {
			// moves
			if square >= startRankLow && square <= startRankHigh {
				// double move
				movements = append(movements, PawnStep(square, game.WhiteToPlay))
				movements = append(movements, NewMove(square, PawnStep(PawnStep(square, game.WhiteToPlay).To, game.WhiteToPlay).To))
			} else if square >= promoteRankLow && square <= promoteRankHigh {
				// promotion move
				if game.WhiteToPlay {
					movements = append(movements,
						Move{square, PawnStep(square, game.WhiteToPlay).To, White_queen, None},
						Move{square, PawnStep(square, game.WhiteToPlay).To, White_knight, None},
						Move{square, PawnStep(square, game.WhiteToPlay).To, White_bishop, None},
						Move{square, PawnStep(square, game.WhiteToPlay).To, White_rook, None},
					)
				} else {
					movements = append(movements,
						Move{square, PawnStep(square, game.WhiteToPlay).To, Black_queen, None},
						Move{square, PawnStep(square, game.WhiteToPlay).To, Black_knight, None},
						Move{square, PawnStep(square, game.WhiteToPlay).To, Black_bishop, None},
						Move{square, PawnStep(square, game.WhiteToPlay).To, Black_rook, None},
					)
				}

			} else {
				// single move
				movements = append(movements, PawnStep(square, game.WhiteToPlay))
			}

			// captues
			captures = append(captures, PawnCaptures(square, game.WhiteToPlay)...)
		}
	}

	for _, move := range movements {
		if (move.To > 63 || move.To < 0) || game.Board.Squares[move.To] != None {
			break
		}
		moves = append(moves, move)
	}

	for _, move := range captures {
		if (move.To < 0 || move.To >= 64) || game.Board.Squares[move.To] == None {
			continue
		}

		if move.From >= promoteRankLow && move.From <= promoteRankHigh {
			if game.WhiteToPlay {
				moves = append(moves,
					Move{move.From, move.To, White_queen, game.Board.Squares[move.To]},
					Move{move.From, move.To, White_knight, game.Board.Squares[move.To]},
					Move{move.From, move.To, White_bishop, game.Board.Squares[move.To]},
					Move{move.From, move.To, White_rook, game.Board.Squares[move.To]},
				)
			} else {
				moves = append(moves,
					Move{move.From, move.To, Black_queen, game.Board.Squares[move.To]},
					Move{move.From, move.To, Black_knight, game.Board.Squares[move.To]},
					Move{move.From, move.To, Black_bishop, game.Board.Squares[move.To]},
					Move{move.From, move.To, Black_rook, game.Board.Squares[move.To]},
				)

			}
		} else {
			moves = append(moves, Move{move.From, move.To, None, game.Board.Squares[move.To]})
		}
	}

	return
}
