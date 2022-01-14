package main

import "fmt"
import "testing"
import "math/rand"

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

type dummyPlayer struct{}
func (d *dummyPlayer) Play(board Board) Move {
	move := Move(rand.Intn(4))
	return move
}

func TestRunGame(t *testing.T) {

	for i:=0; i < 10; i++ {
		e := new(Engine)
		e.InitiateBoard()
		fmt.Println(e.board)
		e.RunGame(new(dummyPlayer))
		fmt.Println(e.board)
		fmt.Println(e.score)
		fmt.Println()
	}
}