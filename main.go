package main

import "fmt"

func main() {
	board := NewChessboard()

	board.place(white_king, coords_to_square_num(1, 1))
	fmt.Println(board)
}
