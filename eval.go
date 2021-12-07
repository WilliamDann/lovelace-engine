package main

import (
	"math"

	"github.com/notnil/chess"
)

var piece_values map[chess.PieceType]float64 = map[chess.PieceType]float64{
	chess.Pawn:   1.00,
	chess.Knight: 3.00,
	chess.Bishop: 3.00,
	chess.Rook:   5.00,
	chess.Queen:  9.00,
}

func eval(game chess.Game) (score float64) {
	if game.Outcome() == chess.WhiteWon {
		return math.Inf(1)
	}
	if game.Outcome() == chess.BlackWon {
		return math.Inf(-1)
	}
	if game.Outcome() == chess.Draw {
		return 0.00
	}

	for _, piece := range game.Position().Board().SquareMap() {
		if piece.Color() == chess.White {
			score += piece_values[piece.Type()]
		} else {
			score -= piece_values[piece.Type()]
		}
	}

	return
}
