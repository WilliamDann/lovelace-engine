package main

import (
	"fmt"
	"math"

	"github.com/notnil/chess"
)

func search(game chess.Game, depth int, a, b float64, white bool) (float64, []chess.Move) {
	if white {
		return maximize(game, depth, a, b)
	} else {
		return minimize(game, depth, a, b)
	}
}

// TODO
//   instead of a fixed depth, ordering moves by best at last depth and ++ing depth
//   will allow a search based on time instead of fixed depth. This has
//   the added benefit of having a great guess of the best moves first
//   which is a huge advantage for pruning
//
//   undo move instead of cloning maybe?
//
//   sort so checks, lower->higher captures, and threats are evaled first
func maximize(game chess.Game, depth int, a, b float64) (float64, []chess.Move) {
	if depth == 0 || game.Outcome() != chess.NoOutcome {
		return eval(game), []chess.Move{}
	}

	maxScore := math.Inf(-1)
	maxScoreMoves := []chess.Move{}
	allLegalMoves := game.ValidMoves()

	for _, legalMove := range allLegalMoves {

		analysisBoard := game.Clone()
		analysisBoard.Move(legalMove)

		lineScore, lineMoves := minimize(*analysisBoard, depth-1, a, b)

		r := legalMove.String()
		ev := fmt.Sprintf("%f", lineScore)
		fmt.Println(r + " :" + ev)

		if lineScore > maxScore {
			maxScoreMoves = append([]chess.Move{*legalMove}, lineMoves...)
			maxScore = lineScore
		}

		a = math.Max(a, lineScore)

		if b <= a {
			break
		}
	}

	return maxScore, maxScoreMoves
}

func minimize(game chess.Game, depth int, a, b float64) (float64, []chess.Move) {
	if depth == 0 || game.Outcome() != chess.NoOutcome {
		return eval(game), []chess.Move{}
	}

	minScore := math.Inf(1)
	minScoreMoves := []chess.Move{}
	allLegalMoves := game.ValidMoves()

	for _, legalMove := range allLegalMoves {
		analysisBoard := game.Clone()
		analysisBoard.Move(legalMove)

		lineScore, lineMoves := maximize(*analysisBoard, depth-1, a, b)

		r := legalMove.String()
		ev := fmt.Sprintf("%f", lineScore)
		fmt.Println("  " + r + " :" + ev)

		if lineScore < minScore {
			minScoreMoves = append([]chess.Move{*legalMove}, lineMoves...)
			minScore = lineScore
		}

		b = math.Min(b, lineScore)

		if b <= a {
			break
		}
	}

	return minScore, minScoreMoves
}
