package main

import "math"
import "math/rand"

type GreedyAI struct{}

func (g *GreedyAI) Play(board Board) Move {
	_, upScore := doMoveUp(board)
	_, downScore := doMoveDown(board)
	_, rightScore := doMoveRight(board)
	_, leftScore := doMoveLeft(board)

	maxScore := int(math.Max(float64(upScore), math.Max(float64(downScore), math.Max(float64(rightScore), float64(leftScore)))))
	maxes := *new([]Move)
	if upScore == maxScore {
		maxes = append(maxes, MoveUp)
	}
	if leftScore == maxScore {
		maxes = append(maxes, MoveLeft)
	}
	if rightScore == maxScore {
		maxes = append(maxes, MoveRight)
	}
	if downScore == maxScore {
		maxes = append(maxes, MoveDown)
	}

	return maxes[rand.Intn(len(maxes))]
}
