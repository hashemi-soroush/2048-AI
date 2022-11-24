package main

import "math/rand"

type Player interface {
	Play(Board) Move
}

type Game struct {
	board  Board
	player Player
	score  int
}

func (g *Game) InitiateBoard() {
	g.board = *new(Board)

	g.putRandomlyOnBoard(2)
	g.putRandomlyOnBoard(2)
}

func (g *Game) RunGame(p Player) int {
	g.player = p

	doMoveFuncs := map[Move]func(Board) (Board, int){
		MoveUp:    doMoveUp,
		MoveDown:  doMoveDown,
		MoveLeft:  doMoveLeft,
		MoveRight: doMoveRight,
	}
	canMoveFuncs := map[Move]func(Board) bool{
		MoveUp:    canMoveUp,
		MoveDown:  canMoveDown,
		MoveRight: canMoveRight,
		MoveLeft:  canMoveLeft,
	}

	for canMove(g.board) {
		copiedBoard := copyBoard(g.board)
		nextMove := g.player.Play(copiedBoard)

		if canMoveFuncs[nextMove](g.board) == false {
			continue
		}

		copiedBoard = copyBoard(g.board)
		newBoard, score := doMoveFuncs[nextMove](copiedBoard)
		g.board = newBoard
		g.score += score
		g.putRandomlyOnBoard([]int{2, 4}[rand.Intn(1)])
	}

	return g.score
}

func (g *Game) putRandomlyOnBoard(number int) {
	type cell struct {
		row int
		col int
	}

	emptyCells := *new([]cell)
	for i := 0; i < len(g.board); i++ {
		for j := 0; j < len(g.board[i]); j++ {
			if g.board[i][j] == 0 {
				emptyCells = append(emptyCells, cell{i, j})
			}
		}
	}

	if len(emptyCells) == 0 {
		return
	}

	selectedCell := emptyCells[rand.Intn(len(emptyCells))]
	g.board[selectedCell.row][selectedCell.col] = number
}
