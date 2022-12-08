package main

import "math/rand"

type RandomAI struct{}

func (r *RandomAI) Play(board Board) Move {
	move := Move(rand.Intn(4))
	return move
}
