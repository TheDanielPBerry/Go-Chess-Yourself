package rules

import (
// "fmt"
)

func IsPawn(piece rune) bool {
	return piece == '♙' || piece == '♟'
}

func IsWhite(piece rune) bool {
	return piece >= '♔' && piece <= '♙'
	//return piece == '♖' || piece == '♘' || piece == '♗' || piece == '♔' || piece == '♕' || piece == '♙'
}

func InitializeMovesRepresentation() [8][8]bool {
	return [8][8]bool{
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
	}
}

func SetValidMove(pos [2]int, board *[8][8]rune, moves *[8][8]bool, possible [2]int) {
	piece := board[pos[1]][pos[0]]
	if possible[0] >= 0 && possible[0] <= 7 && possible[1] >= 0 && possible[1] <= 7 {
		space := board[possible[1]][possible[0]]
		if space == ' ' || IsWhite(space) != IsWhite(piece) {
			moves[possible[1]][possible[0]] = true
		}
	}
}

func GetPossibleMoves(pos [2]int, board [8][8]rune) [8][8]bool {
	moves := InitializeMovesRepresentation()
	piece := board[pos[1]][pos[0]]
	var direction int = -1
	if IsWhite(piece) {
		direction = 1
	}
	if IsPawn(piece) {
		if pos[1] > 0 && pos[1] < 7 {
			if board[pos[1]+direction][pos[0]] == ' ' {
				moves[pos[1]+direction][pos[0]] = true
			}
		} else if direction-1 == pos[1] || (8+direction) == pos[1] {
			//Handle promotion
		}

		//Check if first move of the pawn
		if direction == pos[1] || (7+direction) == pos[1] {
			if board[pos[1]+(direction*2)][pos[0]] == ' ' {
				moves[pos[1]+(direction*2)][pos[0]] = true
			}
		}

		//attack right of piece
		if pos[0] < 7 && board[pos[1]+direction][pos[0]+1] != ' ' {
			moves[pos[1]+direction][pos[0]+1] = true
		}
		//attack left of piece
		if pos[0] > 0 && board[pos[1]+direction][pos[0]-1] != ' ' {
			moves[pos[1]+direction][pos[0]-1] = true
		}
		//Implement en passant
		//TODO
	}

	//Bishop
	if piece == '♝' || piece == '♗' || piece == '♛' || piece == '♕' {
		direction := [2]int{1, 1}
		for pos[1]+direction[1] <= 7 && pos[0]+direction[0] <= 7 {
			space := board[pos[1]+direction[1]][pos[0]+direction[0]]
			if space == ' ' {
				moves[pos[1]+direction[1]][pos[0]+direction[0]] = true
			} else {
				if IsWhite(space) != IsWhite(piece) {
					moves[pos[1]+direction[1]][pos[0]+direction[0]] = true
				}
				break
			}
			direction[0]++
			direction[1]++
		}

		direction = [2]int{-1, 1}
		for pos[1]+direction[1] <= 7 && pos[0]+direction[0] >= 0 {
			space := board[pos[1]+direction[1]][pos[0]+direction[0]]
			if space == ' ' {
				moves[pos[1]+direction[1]][pos[0]+direction[0]] = true
			} else {
				if IsWhite(space) != IsWhite(piece) {
					moves[pos[1]+direction[1]][pos[0]+direction[0]] = true
				}
				break
			}
			direction[0]--
			direction[1]++
		}

		direction = [2]int{-1, -1}
		for pos[1]+direction[1] >= 0 && pos[0]+direction[0] >= 0 {
			space := board[pos[1]+direction[1]][pos[0]+direction[0]]
			if space == ' ' {
				moves[pos[1]+direction[1]][pos[0]+direction[0]] = true
			} else {
				if IsWhite(space) != IsWhite(piece) {
					moves[pos[1]+direction[1]][pos[0]+direction[0]] = true
				}
				break
			}
			direction[0]--
			direction[1]--
		}

		direction = [2]int{1, -1}
		for pos[1]+direction[1] >= 0 && pos[0]+direction[0] <= 7 {
			space := board[pos[1]+direction[1]][pos[0]+direction[0]]
			if space == ' ' {
				moves[pos[1]+direction[1]][pos[0]+direction[0]] = true
			} else {
				if IsWhite(space) != IsWhite(piece) {
					moves[pos[1]+direction[1]][pos[0]+direction[0]] = true
				}
				break
			}
			direction[0]++
			direction[1]--
		}
	}

	if piece == '♜' || piece == '♖' || piece == '♛' || piece == '♕' {
		for x := pos[0] + 1; x <= 7; x++ {
			space := board[pos[1]][x]
			if space == ' ' {
				moves[pos[1]][x] = true
			} else {
				if IsWhite(piece) != IsWhite(space) {
					moves[pos[1]][x] = true
				}
				break
			}
		}
		for x := pos[0] - 1; x >= 0; x-- {
			space := board[pos[1]][x]
			if space == ' ' {
				moves[pos[1]][x] = true
			} else {
				if IsWhite(piece) != IsWhite(space) {
					moves[pos[1]][x] = true
				}
				break
			}
		}

		for y := pos[1] + 1; y <= 7; y++ {
			space := board[y][pos[0]]
			if space == ' ' {
				moves[y][pos[0]] = true
			} else {
				if IsWhite(piece) != IsWhite(space) {
					moves[y][pos[0]] = true
				}
				break
			}
		}
		for y := pos[1] - 1; y >= 0; y-- {
			space := board[y][pos[0]]
			if space == ' ' {
				moves[y][pos[0]] = true
			} else {
				if IsWhite(piece) != IsWhite(space) {
					moves[y][pos[0]] = true
				}
				break
			}
		}
	}

	//Knight
	if piece == '♘' || piece == '♞' {
		offset := [2]int{2, 1}
		possible := [2]int{pos[0] + offset[0], pos[1] + offset[1]}
		SetValidMove(pos, &board, &moves, possible)

		offset = [2]int{-2, 1}
		possible = [2]int{pos[0] + offset[0], pos[1] + offset[1]}
		SetValidMove(pos, &board, &moves, possible)

		offset = [2]int{-2, -1}
		possible = [2]int{pos[0] + offset[0], pos[1] + offset[1]}
		SetValidMove(pos, &board, &moves, possible)

		offset = [2]int{2, -1}
		possible = [2]int{pos[0] + offset[0], pos[1] + offset[1]}
		SetValidMove(pos, &board, &moves, possible)

		offset = [2]int{1, 2}
		possible = [2]int{pos[0] + offset[0], pos[1] + offset[1]}
		SetValidMove(pos, &board, &moves, possible)

		offset = [2]int{-1, 2}
		possible = [2]int{pos[0] + offset[0], pos[1] + offset[1]}
		SetValidMove(pos, &board, &moves, possible)

		offset = [2]int{-1, -2}
		possible = [2]int{pos[0] + offset[0], pos[1] + offset[1]}
		SetValidMove(pos, &board, &moves, possible)

		offset = [2]int{1, -2}
		possible = [2]int{pos[0] + offset[0], pos[1] + offset[1]}
		SetValidMove(pos, &board, &moves, possible)
	}

	//King
	if piece == '♔' || piece == '♚' {
		offset := [2]int{-1, -1}
		possible := [2]int{pos[0] + offset[0], pos[1] + offset[1]}
		SetValidMove(pos, &board, &moves, possible)

		offset = [2]int{-1, 0}
		possible = [2]int{pos[0] + offset[0], pos[1] + offset[1]}
		SetValidMove(pos, &board, &moves, possible)

		offset = [2]int{-1, 1}
		possible = [2]int{pos[0] + offset[0], pos[1] + offset[1]}
		SetValidMove(pos, &board, &moves, possible)

		offset = [2]int{0, -1}
		possible = [2]int{pos[0] + offset[0], pos[1] + offset[1]}
		SetValidMove(pos, &board, &moves, possible)

		offset = [2]int{0, 1}
		possible = [2]int{pos[0] + offset[0], pos[1] + offset[1]}
		SetValidMove(pos, &board, &moves, possible)

		offset = [2]int{1, -1}
		possible = [2]int{pos[0] + offset[0], pos[1] + offset[1]}
		SetValidMove(pos, &board, &moves, possible)

		offset = [2]int{1, 0}
		possible = [2]int{pos[0] + offset[0], pos[1] + offset[1]}
		SetValidMove(pos, &board, &moves, possible)

		offset = [2]int{1, 1}
		possible = [2]int{pos[0] + offset[0], pos[1] + offset[1]}
		SetValidMove(pos, &board, &moves, possible)
	}

	return moves
}

func HasPossibleMoves() {

}

func ValidateMove(piece rune, board [8][8]rune, move [2]int) bool {
	switch piece {
	case '♙':
	case '♟':

		break
	case '♔':
	case '♚':
		break
	}
	return false
}
