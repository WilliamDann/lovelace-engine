package main

import (
	"math"

	"github.com/notnil/chess"
)

func minimaxRoot(game chess.Game, depth, a, b float64, white bool) (float64, chess.Move) {
	if white {
		evalMax := math.Inf(-1)
		var moveMax chess.Move

		moves := game.ValidMoves()
		for _, move := range moves {
			clone := game.Clone()
			clone.Move(move)

			nodeEval := minimax(*clone, depth-1, a, b, false)
			if nodeEval > evalMax {
				moveMax = *move
				evalMax = nodeEval
			}
			a = math.Max(a, nodeEval)

			if b <= a {
				break
			}
		}
		return evalMax, moveMax
	} else {
		evalMin := math.Inf(1)
		var moveMin chess.Move

		moves := game.ValidMoves()
		for _, move := range moves {
			clone := game.Clone()
			clone.Move(move)

			nodeEval := minimax(*clone, depth-1, a, b, true)
			if nodeEval < evalMin {
				moveMin = *move
				evalMin = nodeEval
			}
			b = math.Min(b, nodeEval)

			if b <= a {
				break
			}
		}
		return evalMin, moveMin
	}
}

// TODO
//   instead of a fixed depth, ordering moves by best last depth and ++ing will allow
//   a search based on time instead of fixed depth. This has the added benefit of
//   having a great guess of the best moves first which is a huge advantage for
//   pruning
//
//   undo move instead of cloning maybe?
func minimax(game chess.Game, depth, a, b float64, white bool) float64 {
	if depth == 0 || game.Outcome() != chess.NoOutcome {
		return eval(game)
	}

	if white {
		evalMax := math.Inf(-1)
		moves := game.ValidMoves()
		for _, move := range moves {
			clone := game.Clone()
			clone.Move(move)

			nodeEval := minimax(*clone, depth-1, a, b, false)
			evalMax = math.Max(evalMax, nodeEval)
			a = math.Max(a, nodeEval)

			if b <= a {
				break
			}
		}
		return evalMax
	} else {
		evalMin := math.Inf(1)
		moves := game.ValidMoves()
		for _, move := range moves {
			clone := game.Clone()
			clone.Move(move)

			nodeEval := minimax(*clone, depth-1, a, b, true)
			evalMin = math.Min(evalMin, nodeEval)
			b = math.Min(b, nodeEval)
			if b <= a {
				break
			}
		}
		return evalMin
	}
}
