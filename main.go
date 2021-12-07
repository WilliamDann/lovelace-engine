package main

import (
	"fmt"
	"math"

	"github.com/notnil/chess"
)

func main() {
	fen, _ := chess.FEN("7r/p5pp/R4p2/2N2n2/5k2/P4r1P/5P2/4R1K1 w - - 2 31")
	game := chess.NewGame(fen)

	score, move := minimaxRoot(*game, 2, math.Inf(-1), math.Inf(1), true)

	fmt.Println(score)
	fmt.Println(move.String())
}
