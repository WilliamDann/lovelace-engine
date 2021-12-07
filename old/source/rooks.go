package source

type BoardDirection int

const (
	North BoardDirection = -8
	South BoardDirection = 8
	East  BoardDirection = 1
	West  BoardDirection = -1

	Northeast BoardDirection = -7
	Southeast BoardDirection = 9
	Southwest BoardDirection = 7
	Northwest BoardDirection = -9
)

func BoardRay(start int, direction BoardDirection) (hits Bitboard) {
	for i := start; i < 64 && i >= 0; i += int(direction) {
		// if a board edge is reached on east or west directions it will continue on because that index is still on the board
		//    this stops the ray from extending into those indexes
		if direction == East || direction == Northeast || direction == Southeast {
			if i%8 == 0 && i != 0 {
				break
			}
		}
		if direction == West || direction == Northwest || direction == Southwest {
			hits.Board[i] = true
			if i%8 == 0 && i != 0 {
				break
			}
		}

		hits.Board[i] = true
	}
	return
}

func RookStep(board Game, start int, direction BoardDirection) (moves []Move) {
	ray := BoardRay(start, direction)
	for i, square := range ray.Board {
		if square && i != start {
			if board.Board.Squares[i] != None {
				moves = append(moves, Move{From: start, To: i, Capture: board.Board.Squares[i], Promote: None})
				break
			}

			moves = append(moves, Move{From: start, To: i, Capture: None, Promote: None})
		}
	}
	return
}

func (game Game) GetRookMoves() (moves []Move) {
	seeking := White_rook
	if !game.WhiteToPlay {
		seeking = Black_rook
	}

	for square, rook := range game.Board.GetPieceLocations(seeking).Board {
		if rook {
			moves = append(moves, RookStep(game, square, North)...)
			moves = append(moves, RookStep(game, square, South)...)
			moves = append(moves, RookStep(game, square, East)...)
			moves = append(moves, RookStep(game, square, West)...)
		}
	}

	return
}
