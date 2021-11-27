package source

func _SliderStep(square, stepX, stepY, i int) (moves []Move) {
	if stepX*i < 0 || stepX*i > 7 || stepY*i < 0 || stepY*i > 7 {
		return
	}
	step := stepX*i + (stepY*i)*8

	moves = append(moves, Move{square, step, None, None})
	return append(moves, _SliderStep(square, stepX, stepY, i+1)...)
}

func SliderStep(square, stepX, stepY int) (moves []Move) {
	return _SliderStep(square, stepX, stepY, 1)
}

func (game Game) GetRookMoves() (moves []Move) {
	seeking := White_rook
	if !game.WhiteToPlay {
		seeking = Black_rook
	}

	for square, rook := range game.Board.GetPieceLocations(seeking).Board {
		if rook {
			moves = append(moves, SliderStep(square, 1, 0)...)
			moves = append(moves, SliderStep(square, -1, 0)...)
			moves = append(moves, SliderStep(square, 0, 1)...)
			moves = append(moves, SliderStep(square, -1, 0)...)
		}
	}

	return
}
