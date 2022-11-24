package main

type SmartAI struct{}

func (s *SmartAI) Play(board Board) Move {
	if canMoveUp(board) {
		return MoveUp
	} else if canMoveLeft(board) {
		return MoveLeft
	} else if canMoveRight(board) {
		return MoveRight
	} else {
		return MoveDown
	}
	return MoveUp
}
