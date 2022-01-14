package main

import "fmt"
import "testing"
import "math/rand"

func TestInitiateBoard(t *testing.T) {
	g := new(Game)
	g.InitiateBoard()

	nonZeroCount := 0
	for i := 0; i < len(g.board); i++ {
		for j := 0; j < len(g.board[i]); j++ {
			if g.board[i][j] != 0 {
				nonZeroCount++
			}

			if !(g.board[i][j] == 0 || g.board[i][j] == 2) {
				t.Errorf("Initiated board must only have 2s and 0s. board[%d][%d] is %d", i, j, g.board[i][j])
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
		g := new(Game)
		g.InitiateBoard()
		fmt.Println(g.board)
		g.RunGame(new(dummyPlayer))
		fmt.Println(g.board)
		fmt.Println(g.score)
		fmt.Println()
	}
}