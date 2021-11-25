package main

// There are 64 square on a chessboard, and 64 bits in a int64
// https://www.chessprogramming.org/Bitboards
//
// for some reason in Go you cannot set the last bit of a uint64
//     beacuse of this reason I am using a [64]bool instead.
//     this is simpler to use than bit hacks anyway
type Bitboard struct {
	board [64]bool
}

func coords_to_square_num(x, y int) int {
	return y*8 + x
}

func (board Bitboard) String() (output string) {
	output = ""

	for i := 0; i < 64; i++ {
		if board.board[i] {
			output += "1 "
		} else {
			output += "0 "
		}

		if (i+1)%8 == 0 && i != 0 {
			output += "\n"
		}
	}

	return
}
