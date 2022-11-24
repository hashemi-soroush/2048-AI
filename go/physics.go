package main

import "reflect"

type Board [4][4]int
type Move int

const (
	MoveUp Move = iota
	MoveDown
	MoveRight
	MoveLeft
)

func copyBoard(board Board) Board {
	copiedBoard := new(Board)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			copiedBoard[i][j] = board[i][j]
		}
	}
	return *copiedBoard
}

func canMove(board Board) bool {
	return canMoveUp(board) || canMoveDown(board) || canMoveLeft(board) || canMoveRight(board)
}

func canMoveUp(board Board) bool {
	copiedBoard := copyBoard(board)
	boardAfterMove, _ := doMoveUp(copiedBoard)
	return reflect.DeepEqual(board, boardAfterMove) == false
}

func canMoveDown(board Board) bool {
	copiedBoard := copyBoard(board)
	boardAfterMove, _ := doMoveDown(copiedBoard)
	return reflect.DeepEqual(board, boardAfterMove) == false
}

func canMoveRight(board Board) bool {
	copiedBoard := copyBoard(board)
	boardAfterMove, _ := doMoveRight(copiedBoard)
	return reflect.DeepEqual(board, boardAfterMove) == false
}

func canMoveLeft(board Board) bool {
	copiedBoard := copyBoard(board)
	boardAfterMove, _ := doMoveLeft(copiedBoard)
	return reflect.DeepEqual(board, boardAfterMove) == false
}

func doMoveUp(board Board) (Board, int) {
	score := 0
	for j := 0; j < len(board[0]); j++ {
		cur := 0
		for i := 0; i < len(board); i++ {
			if board[i][j] == 0 || i == cur {
				continue
			}

			if board[cur][j] == 0 {
				board[cur][j] = board[i][j]
				board[i][j] = 0
				continue
			}

			if board[i][j] == board[cur][j] {
				board[cur][j] *= 2
				score += board[cur][j]
				board[i][j] = 0
				cur += 1
				continue
			} else {
				cur += 1
				temp := board[i][j]
				board[i][j] = 0
				board[cur][j] = temp
				continue
			}
		}
	}
	return board, score
}

func doMoveDown(board Board) (Board, int) {
	score := 0
	for j := 0; j < len(board[0]); j++ {
		cur := len(board) - 1
		for i := len(board) - 1; i >= 0; i-- {
			if board[i][j] == 0 || i == cur {
				continue
			}

			if board[cur][j] == 0 {
				board[cur][j] = board[i][j]
				board[i][j] = 0
				continue
			}

			if board[i][j] == board[cur][j] {
				board[cur][j] *= 2
				score += board[cur][j]
				board[i][j] = 0
				cur -= 1
				continue
			} else {
				cur -= 1
				temp := board[i][j]
				board[i][j] = 0
				board[cur][j] = temp
				continue
			}
		}
	}
	return board, score
}

func doMoveRight(board Board) (Board, int) {
	score := 0
	for i := 0; i < len(board); i++ {
		cur := len(board[i]) - 1
		for j := len(board[i]) - 1; j >= 0; j-- {
			if board[i][j] == 0 || j == cur {
				continue
			}

			if board[i][cur] == 0 {
				board[i][cur] = board[i][j]
				board[i][j] = 0
				continue
			}

			if board[i][j] == board[i][cur] {
				board[i][cur] *= 2
				score += board[i][cur]
				board[i][j] = 0
				cur -= 1
				continue
			} else {
				cur -= 1
				temp := board[i][j]
				board[i][j] = 0
				board[i][cur] = temp
				continue
			}
		}
	}
	return board, score
}

func doMoveLeft(board Board) (Board, int) {
	score := 0
	for i := 0; i < len(board); i++ {
		cur := 0
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 0 || j == cur {
				continue
			}

			if board[i][cur] == 0 {
				board[i][cur] = board[i][j]
				board[i][j] = 0
				continue
			}

			if board[i][j] == board[i][cur] {
				board[i][cur] *= 2
				score += board[i][cur]
				board[i][j] = 0
				cur += 1
				continue
			} else {
				cur += 1
				temp := board[i][j]
				board[i][j] = 0
				board[i][cur] = temp
				continue
			}
		}
	}
	return board, score
}
