package main

import "fmt"

func main() {
	board := NewChessboard()

	board.squares[coords_to_square_num(1, 1)] = white_pawn
	fmt.Println(board)
}
