package main

import "fmt"

func main() {
	board := ParseFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR")
	fmt.Println(board)
}
