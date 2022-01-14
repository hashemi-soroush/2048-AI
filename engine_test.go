package main

import "reflect"
import "testing"

func TestInitiateBoard(t *testing.T) {
	e := new(Engine)
	e.InitiateBoard()

	nonZeroCount := 0
	for i := 0; i < len(e.board); i++ {
		for j := 0; j < len(e.board[i]); j++ {
			if e.board[i][j] != 0 {
				nonZeroCount++
			}

			if !(e.board[i][j] == 0 || e.board[i][j] == 2) {
				t.Errorf("Initiated board must only have 2s and 0s. board[%d][%d] is %d", i, j, e.board[i][j])
			}
		}
	}
	if nonZeroCount != 2 {
		t.Errorf("Initiated board must contain only 2 non-zero elements")
	}
}

func TestCanMove(t *testing.T) {
	e := new(Engine)
	
	///////////////////////////////////
	// test sample
	e.board = &Board {
		{2, 4, 2, 4},
		{4, 2, 4, 2},
		{2, 4, 2, 4},
		{4, 2, 4, 2},
	}
	if e.canMove() == true {
		t.Errorf("No valid move is possible in this board")
	}
	///////////////////////////////////

	///////////////////////////////////
	// test sample
	e.board = &Board {
		{2, 4, 2, 4},
		{4, 2, 4, 2},
		{2, 0, 2, 4},
		{4, 2, 4, 2},
	}
	if e.canMove() == false {
		t.Errorf("There is a valid move in this board")
	}
	///////////////////////////////////
}

func TestDoMoveUp(t *testing.T) {
	e := new(Engine)

	///////////////////////////////////
	// test sample
	e.board = &Board {
		{0, 2, 2, 2},
		{0, 0, 2, 0},
		{0, 0, 0, 2},
		{0, 0, 0, 0},
	}
	targetBoard := &Board {
		{0, 2, 4, 4},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	e.doMoveUp()
	if reflect.DeepEqual(e.board, targetBoard) == false {
		t.Errorf("doMoveUp is making a mistake")
	}
	///////////////////////////////////

	///////////////////////////////////
	// test sample
	e.board = &Board {
		{0, 2, 0, 0},
		{0, 0, 2, 2},
		{2, 0, 2, 2},
		{0, 2, 0, 2},
	}
	targetBoard = &Board {
		{2, 4, 4, 4},
		{0, 0, 0, 2},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	e.doMoveUp()
	if reflect.DeepEqual(e.board, targetBoard) == false {
		t.Errorf("doMoveUp is making a mistake")
	}
	///////////////////////////////////

	///////////////////////////////////
	// test sample
	e.board = &Board {
		{2, 0, 0, 4},
		{2, 2, 4, 0},
		{2, 2, 2, 2},
		{2, 4, 2, 2},
	}
	targetBoard = &Board {
		{4, 4, 4, 4},
		{4, 4, 4, 4},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	e.doMoveUp()
	if reflect.DeepEqual(e.board, targetBoard) == false {
		t.Errorf("doMoveUp is making a mistake")
	}
	///////////////////////////////////
}