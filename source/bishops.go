package source

func (game Game) GetBishopMoves() (moves []Move) {
	seeking := White_rook
	if !game.WhiteToPlay {
		seeking = Black_rook
	}

	for square, rook := range game.Board.GetPieceLocations(seeking).Board {
		if rook {
			moves = append(moves, SliderStep(game, square, 1, 1)...)
			moves = append(moves, SliderStep(game, square, 1, -1)...)
			moves = append(moves, SliderStep(game, square, -1, 1)...)
			moves = append(moves, SliderStep(game, square, -1, -1)...)
		}
	}

	return
}
