package main

import "fmt"

func main() {
	board := ParseFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR")
	game := Game{board, true, []Move{}}
	fmt.Println(game)

	game = game.PushMove(game.GetPawnMoves()[7])

	fmt.Println(game)
}
