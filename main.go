package main

import (
	"fmt"
	"math"

	"github.com/notnil/chess"
)

func main() {
	fen, _ := chess.FEN("1r4r1/1p3Rpp/pN1pk3/8/3PP3/1n6/5RK1/8 w - - 1 2")
	game := chess.NewGame(fen)

	score, moves := search(*game, 4, math.Inf(-1), math.Inf(1), true)

	for _, move := range moves {
		game.Move(&move)
	}

	fmt.Print(game.String() + " : ")
	fmt.Println(score)
}
