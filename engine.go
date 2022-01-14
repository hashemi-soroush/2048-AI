package main

import "math/rand"
import "time"

type Player interface {
	Play(Board) Move
}

type Engine struct {
	board  Board
	player Player
	score int
}

func (e *Engine) InitiateBoard() {
	e.board = *new(Board)

	rand.Seed(time.Now().UnixNano())

	e.putRandomlyOnBoard(2)
	e.putRandomlyOnBoard(2)
}

func (e *Engine) RunGame(p Player) {
	e.player = p

	doMoveFuncs := map[Move]func(Board)(Board, int){
		MoveUp:    doMoveUp,
		MoveDown:  doMoveDown,
		MoveLeft:  doMoveLeft,
		MoveRight: doMoveRight,
	}
	canMoveFuncs := map[Move]func(Board)bool{
		MoveUp: canMoveUp,
		MoveDown: canMoveDown,
		MoveRight: canMoveRight,
		MoveLeft: canMoveLeft,
	}

	for canMove(e.board) {
		copiedBoard := copyBoard(e.board)
		nextMove := e.player.Play(copiedBoard)
		
		if canMoveFuncs[nextMove](e.board) == false {
			continue
		}

		copiedBoard = copyBoard(e.board)
		newBoard, score := doMoveFuncs[nextMove](copiedBoard)
		e.board = newBoard
		e.score += score
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
