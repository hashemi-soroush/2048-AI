package main

import "testing"
import "reflect"

func TestCanMove(t *testing.T) {
	board := *new(Board)

	///////////////////////////////////
	// test sample
	board = Board{
		{2, 4, 2, 4},
		{4, 2, 4, 2},
		{2, 4, 2, 4},
		{4, 2, 4, 2},
	}
	if canMove(board) == true {
		t.Errorf("No valid move is possible in this board")
	}
	///////////////////////////////////

	///////////////////////////////////
	// test sample
	board = Board{
		{2, 4, 2, 4},
		{4, 2, 4, 2},
		{2, 0, 2, 4},
		{4, 2, 4, 2},
	}
	if canMove(board) == false {
		t.Errorf("There is a valid move in this board")
	}
	///////////////////////////////////

	///////////////////////////////////
	// test sample
	board = Board{
		{2, 4, 2, 4},
		{4, 2, 4, 2},
		{8, 8, 2, 4},
		{4, 2, 4, 2},
	}
	if canMove(board) == false {
		t.Errorf("There is a valid move in this board")
	}
	///////////////////////////////////

	///////////////////////////////////
	// test sample
	board = Board{
		{2, 4, 2, 4},
		{4, 2, 4, 2},
		{8, 4, 2, 4},
		{8, 2, 4, 2},
	}
	if canMove(board) == false {
		t.Errorf("There is a valid move in this board")
	}
	///////////////////////////////////
}

func TestDoMoveUp(t *testing.T) {
	startBoard := *new(Board)
	endBoard := *new(Board)
	targetBoard := *new(Board)

	///////////////////////////////////
	// test sample
	startBoard = Board{
		{0, 2, 2, 2},
		{0, 0, 2, 0},
		{0, 0, 0, 2},
		{0, 0, 0, 0},
	}
	targetBoard = Board{
		{0, 2, 4, 4},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	endBoard, _ = doMoveUp(startBoard)
	if reflect.DeepEqual(endBoard, targetBoard) == false {
		t.Errorf("doMoveUp is making a mistake")
	}
	///////////////////////////////////

	///////////////////////////////////
	// test sample
	startBoard = Board{
		{0, 2, 0, 0},
		{0, 0, 2, 2},
		{2, 0, 2, 2},
		{0, 2, 0, 2},
	}
	targetBoard = Board{
		{2, 4, 4, 4},
		{0, 0, 0, 2},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	endBoard, _ = doMoveUp(startBoard)
	if reflect.DeepEqual(endBoard, targetBoard) == false {
		t.Errorf("doMoveUp is making a mistake")
	}
	///////////////////////////////////

	///////////////////////////////////
	// test sample
	startBoard = Board{
		{2, 0, 0, 4},
		{2, 2, 4, 0},
		{2, 2, 2, 2},
		{2, 4, 2, 2},
	}
	targetBoard = Board{
		{4, 4, 4, 4},
		{4, 4, 4, 4},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	endBoard, _ = doMoveUp(startBoard)
	if reflect.DeepEqual(endBoard, targetBoard) == false {
		t.Errorf("doMoveUp is making a mistake")
	}
	///////////////////////////////////
}

func TestDoMoveDown(t *testing.T) {
	startBoard := *new(Board)
	endBoard := *new(Board)
	targetBoard := *new(Board)

	///////////////////////////////////
	// test sample
	startBoard = Board{
		{0, 2, 2, 2},
		{0, 0, 2, 0},
		{0, 0, 0, 2},
		{0, 0, 0, 0},
	}
	targetBoard = Board{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 2, 4, 4},
	}
	endBoard, _ = doMoveDown(startBoard)
	if reflect.DeepEqual(endBoard, targetBoard) == false {
		t.Errorf("doMoveDown is making a mistake")
	}
	///////////////////////////////////

	///////////////////////////////////
	// test sample
	startBoard = Board{
		{0, 2, 0, 0},
		{0, 0, 2, 2},
		{2, 0, 2, 2},
		{0, 2, 0, 2},
	}
	targetBoard = Board{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 2},
		{2, 4, 4, 4},
	}
	endBoard, _ = doMoveDown(startBoard)
	if reflect.DeepEqual(endBoard, targetBoard) == false {
		t.Errorf("doMoveDown is making a mistake")
	}
	///////////////////////////////////

	///////////////////////////////////
	// test sample
	startBoard = Board{
		{2, 0, 0, 4},
		{2, 2, 4, 0},
		{2, 2, 2, 2},
		{2, 4, 2, 2},
	}
	targetBoard = Board{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{4, 4, 4, 4},
		{4, 4, 4, 4},
	}
	endBoard, _ = doMoveDown(startBoard)
	if reflect.DeepEqual(endBoard, targetBoard) == false {
		t.Errorf("doMoveDown is making a mistake")
	}
	///////////////////////////////////
}

func TestDoMoveRight(t *testing.T) {
	startBoard := *new(Board)
	endBoard := *new(Board)
	targetBoard := *new(Board)

	///////////////////////////////////
	// test sample
	startBoard = Board{
		{0, 0, 0, 0},
		{2, 0, 0, 0},
		{2, 2, 0, 0},
		{2, 0, 2, 0},
	}
	targetBoard = Board{
		{0, 0, 0, 0},
		{0, 0, 0, 2},
		{0, 0, 0, 4},
		{0, 0, 0, 4},
	}
	endBoard, _ = doMoveRight(startBoard)
	if reflect.DeepEqual(endBoard, targetBoard) == false {
		t.Errorf("doMoveRight is making a mistake")
	}
	///////////////////////////////////

	///////////////////////////////////
	// test sample
	startBoard = Board{
		{0, 0, 2, 0},
		{2, 0, 0, 2},
		{0, 2, 2, 0},
		{0, 2, 2, 2},
	}
	targetBoard = Board{
		{0, 0, 0, 2},
		{0, 0, 0, 4},
		{0, 0, 0, 4},
		{0, 0, 2, 4},
	}
	endBoard, _ = doMoveRight(startBoard)
	if reflect.DeepEqual(endBoard, targetBoard) == false {
		t.Errorf("doMoveRight is making a mistake")
	}
	///////////////////////////////////

	///////////////////////////////////
	// test sample
	startBoard = Board{
		{2, 2, 2, 2},
		{0, 2, 2, 4},
		{0, 4, 2, 2},
		{4, 0, 2, 2},
	}
	targetBoard = Board{
		{0, 0, 4, 4},
		{0, 0, 4, 4},
		{0, 0, 4, 4},
		{0, 0, 4, 4},
	}
	endBoard, _ = doMoveRight(startBoard)
	if reflect.DeepEqual(endBoard, targetBoard) == false {
		t.Errorf("doMoveRight is making a mistake")
	}
	///////////////////////////////////
}

func TestDoMoveLeft(t *testing.T) {
	startBoard := *new(Board)
	endBoard := *new(Board)
	targetBoard := *new(Board)

	///////////////////////////////////
	// test sample
	startBoard = Board{
		{0, 0, 0, 0},
		{2, 0, 0, 0},
		{2, 2, 0, 0},
		{2, 0, 2, 0},
	}
	targetBoard = Board{
		{0, 0, 0, 0},
		{2, 0, 0, 0},
		{4, 0, 0, 0},
		{4, 0, 0, 0},
	}
	endBoard, _ = doMoveLeft(startBoard)
	if reflect.DeepEqual(endBoard, targetBoard) == false {
		t.Errorf("doMoveLeft is making a mistake")
	}
	///////////////////////////////////

	///////////////////////////////////
	// test sample
	startBoard = Board{
		{0, 0, 2, 0},
		{2, 0, 0, 2},
		{0, 2, 2, 0},
		{0, 2, 2, 2},
	}
	targetBoard = Board{
		{2, 0, 0, 0},
		{4, 0, 0, 0},
		{4, 0, 0, 0},
		{4, 2, 0, 0},
	}
	endBoard, _ = doMoveLeft(startBoard)
	if reflect.DeepEqual(endBoard, targetBoard) == false {
		t.Errorf("doMoveLeft is making a mistake")
	}
	///////////////////////////////////

	///////////////////////////////////
	// test sample
	startBoard = Board{
		{2, 2, 2, 2},
		{0, 2, 2, 4},
		{0, 4, 2, 2},
		{4, 0, 2, 2},
	}
	targetBoard = Board{
		{4, 4, 0, 0},
		{4, 4, 0, 0},
		{4, 4, 0, 0},
		{4, 4, 0, 0},
	}
	endBoard, _ = doMoveLeft(startBoard)
	if reflect.DeepEqual(endBoard, targetBoard) == false {
		t.Errorf("doMoveLeft is making a mistake")
	}
	///////////////////////////////////
}
