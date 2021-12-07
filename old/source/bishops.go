package source

func (game Game) GetBishopMoves() (moves []Move) {
	seeking := White_bishop
	if !game.WhiteToPlay {
		seeking = Black_bishop
	}

	for _, bishop := range game.Board.GetPieceLocations(seeking).Board {
		if bishop {
			return
		}
	}

	return
}
