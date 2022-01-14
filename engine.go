package main

import "math/rand"
import "time"

type Board [4][4]int
type Move int

const (
	MoveUp Move = iota
	MoveDown
	MoveRight
	MoveLeft
)

type Player interface {
	Play(Board) Move
}

type Engine struct {
	board  *Board
	player Player
}

func (e *Engine) InitiateBoard() {
	e.board = new(Board)

	rand.Seed(time.Now().UnixNano())

	e.putRandomlyOnBoard(2)
	e.putRandomlyOnBoard(2)
}

func (e *Engine) RunGame(p Player) {
	e.player = p

	moveFuncs := map[Move]func(){
		MoveUp:    e.doMoveUp,
		MoveDown:  e.doMoveDown,
		MoveLeft:  e.doMoveLeft,
		MoveRight: e.doMoveRight,
	}

	for e.canMove() {
		copiedBoard := e.copyBoard()
		nextMove := e.player.Play(copiedBoard)
		
		if (nextMove == MoveUp || nextMove == MoveDown) && e.canMoveVertical() == false {
			continue
		}
		if (nextMove == MoveRight || nextMove == MoveLeft) && e.canMoveHorizontal() == false {
			continue
		}
		
		moveFuncs[nextMove]()
		e.putRandomlyOnBoard([]int{2, 4}[rand.Intn(1)])
	}
}

func (e *Engine) putRandomlyOnBoard(number int) {
	type cell struct {
		row int
		col int
	}

	emptyCells := *new([]cell)
	for i := 0; i < len(e.board); i++ {
		for j := 0; j < len(e.board[i]); j++ {
			if e.board[i][j] == 0 {
				emptyCells = append(emptyCells, cell{i, j})
			}
		}
	}

	if len(emptyCells) == 0 {
		return
	}

	selectedCell := emptyCells[rand.Intn(len(emptyCells))]
	e.board[selectedCell.row][selectedCell.col] = number
}

func (e *Engine) copyBoard() Board {
	copiedBoard := new(Board)
	for i := 0; i < len(e.board); i++ {
		for j := 0; j < len(e.board[i]); j++ {
			copiedBoard[i][j] = e.board[i][j]
		}
	}
	return *copiedBoard
}

func (e *Engine) canMove() bool {
	return e.canMoveHorizontal() || e.canMoveVertical()
}

func (e *Engine) canMoveHorizontal() bool {
	for i := 0; i < len(e.board); i++ {
		for j := 0; j < len(e.board[i]); j++ {
			if e.board[i][j] == 0 {
				return true
			}

			if j+1 < len(e.board[i]) && e.board[i][j] == e.board[i][j+1] {
				return true
			}
		}
	}
	return false
}

func (e *Engine) canMoveVertical() bool {
	for i := 0; i < len(e.board); i++ {
		for j := 0; j < len(e.board[i]); j++ {
			if e.board[i][j] == 0 {
				return true
			}

			if i+1 < len(e.board) && e.board[i][j] == e.board[i+1][j] {
				return true
			}
		}
	}
	return false
}

func (e *Engine) doMoveUp() {
	for j := 0; j < len(e.board[0]); j++ {
		cur := 0
		for i := 0; i < len(e.board); i++ {
			if e.board[i][j] == 0 || i == cur {
				continue
			}

			if e.board[cur][j] == 0 {
				e.board[cur][j] = e.board[i][j]
				e.board[i][j] = 0
				continue
			}

			if e.board[i][j] == e.board[cur][j] {
				e.board[cur][j] *= 2
				e.board[i][j] = 0
				cur += 1
				continue
			} else {
				cur += 1
				temp := e.board[i][j]
				e.board[i][j] = 0
				e.board[cur][j] = temp
				continue
			}
		}
	}
}

func (e *Engine) doMoveDown() {
	for j := 0; j < len(e.board[0]); j++ {
		cur := len(e.board) - 1
		for i := len(e.board) - 1; i >= 0; i-- {
			if e.board[i][j] == 0 || i == cur {
				continue
			}

			if e.board[cur][j] == 0 {
				e.board[cur][j] = e.board[i][j]
				e.board[i][j] = 0
				continue
			}

			if e.board[i][j] == e.board[cur][j] {
				e.board[cur][j] *= 2
				e.board[i][j] = 0
				cur -= 1
				continue
			} else {
				cur -= 1
				temp := e.board[i][j]
				e.board[i][j] = 0
				e.board[cur][j] = temp
				continue
			}
		}
	}
}

func (e *Engine) doMoveRight() {
	for i := 0; i < len(e.board); i++ {
		cur := len(e.board[i]) - 1
		for j := len(e.board[i]) - 1; j >= 0; j-- {
			if e.board[i][j] == 0 || j == cur {
				continue
			}

			if e.board[i][cur] == 0 {
				e.board[i][cur] = e.board[i][j]
				e.board[i][j] = 0
				continue
			}

			if e.board[i][j] == e.board[i][cur] {
				e.board[i][cur] *= 2
				e.board[i][j] = 0
				cur -= 1
				continue
			} else {
				cur -= 1
				temp := e.board[i][j]
				e.board[i][j] = 0
				e.board[i][cur] = temp
				continue
			}
		}
	}
}

func (e *Engine) doMoveLeft() {
	for i := 0; i < len(e.board); i++ {
		cur := 0
		for j := 0; j < len(e.board[i]); j++ {
			if e.board[i][j] == 0 || j == cur {
				continue
			}

			if e.board[i][cur] == 0 {
				e.board[i][cur] = e.board[i][j]
				e.board[i][j] = 0
				continue
			}

			if e.board[i][j] == e.board[i][cur] {
				e.board[i][cur] *= 2
				e.board[i][j] = 0
				cur += 1
				continue
			} else {
				cur += 1
				temp := e.board[i][j]
				e.board[i][j] = 0
				e.board[i][cur] = temp
				continue
			}
		}
	}
}
