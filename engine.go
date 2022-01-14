package main

import "math/rand"
import "time"

type Board [4][4]int;
type Move int;
const (
	MoveUp Move = iota
	MoveDown
	MoveRight
	MoveLeft
)
type Player interface {
	Play(Board) Move;
}

type Engine struct {
	board *Board;
	player Player;
}

func (e *Engine) InitiateBoard() {
	e.board = new(Board);

	rand.Seed(time.Now().UnixNano());
	
	e.putRandomlyOnBoard(2);
	e.putRandomlyOnBoard(2);
}

func (e *Engine) RunGame(p Player) {
	e.player = p;

	moveFuncs := map[Move]func(){
		MoveUp: e.doMoveUp,
		MoveDown: e.doMoveDown,
		MoveLeft: e.doMoveLeft,
		MoveRight: e.doMoveRight,
	};

	for e.canMove() {
		copiedBoard := e.copyBoard()
		nextMove := e.player.Play(copiedBoard)
		moveFuncs[nextMove]()

		e.putRandomlyOnBoard([]int{2, 4}[rand.Intn(1)]);
	}
}

func (e *Engine) putRandomlyOnBoard(number int) {
	type cell struct {
		row int;
		col int;
	}

	emptyCells := *new([]cell)
	for i := 0; i < len(e.board); i++ {
		for j := 0; j < len(e.board[i]); j++ {
			if e.board[i][j] == 0 {
				emptyCells = append(emptyCells, cell{i, j});
			}
		}
	}

	if len(emptyCells) == 0 {
		return
	}

	selectedCell := emptyCells[rand.Intn(len(emptyCells))];
	e.board[selectedCell.row][selectedCell.col] = number;
}

func (e *Engine) copyBoard() Board {
	copiedBoard := new(Board);
	for i := 0; i < len(e.board); i++ {
		for j := 0; j < len(e.board[i]); j++ {
			copiedBoard[i][j] = e.board[i][j];
		}
	}
	return *copiedBoard;
}

func (e *Engine) canMove() bool {
	for i := 0; i < len(e.board); i++ {
		for j := 0; j < len(e.board[i]); j++ {
			if e.board[i][j] == 0 {
				return true;
			}

			if i + 1 < len(e.board) && e.board[i][j] == e.board[i+1][j] {
				return true;
			}

			if j + 1 < len(e.board[i]) && e.board[i][j] == e.board[i][j+1] {
				return true;
			}
		}
	}
	return false;
}

func (e *Engine) doMoveUp() {
	for j := 0; j < len(e.board[0]); j++ {
		top := 0;
		for i := 0; i < len(e.board); i++ {
			if e.board[i][j] == 0 || i == top {
				continue;
			}

			if e.board[top][j] == 0 {
				e.board[top][j] = e.board[i][j];
				e.board[i][j] = 0;
				continue;
			}

			if e.board[i][j] == e.board[top][j] {
				e.board[top][j] *= 2;
				e.board[i][j] = 0;
				top += 1;
				continue;
			} else {
				top += 1;
				temp := e.board[i][j];
				e.board[i][j] = 0;
				e.board[top][j] = temp;
				continue;
			}
		}

	}
}

func (e *Engine) doMoveDown() {

}

func (e *Engine) doMoveLeft() {

}

func (e *Engine) doMoveRight() {

}