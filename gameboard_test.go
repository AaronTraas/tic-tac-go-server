package main

import (
	"testing"
)

var board GameBoard = GameBoard{
	{X, X, O},
	{O, O, X},
	{X, O, O},
}
var board_rotcw GameBoard = GameBoard{
	{X, O, X},
	{O, O, X},
	{O, X, O},
}
var board_fliph GameBoard = GameBoard{
	{O, X, X},
	{X, O, O},
	{O, O, X},
}
var board_flipv GameBoard = GameBoard{
	{X, O, O},
	{O, O, X},
	{X, X, O},
}

func TestRotate(t *testing.T) {

	if *board.Rotate() != board_rotcw {
		t.Error("Rotation clockwise once failed")
	}

	if *board.Rotate().Rotate().Rotate().Rotate() != board {
		t.Error("Rotation clockwise 4 times should result in the same board")
	}
}

func TestFlipH(t *testing.T) {

	if *board.FlipH() != board_fliph {
		t.Error("Horizontal flip failed")
	}

	if *board.FlipH().FlipH() != board {
		t.Error("Flipping twice horizontally should result in the same board")
	}
}

func TestFlipV(t *testing.T) {

	if *board.FlipV() != board_flipv {
		t.Error("Vertical flip failed")
	}

	if *board.FlipV().FlipV() != board {
		t.Error("Flipping twice vertically should result in the same board")
	}
}

func TestEq(t *testing.T) {

	if !board.Equals(&board) {
		t.Error("Board should be equal to itself")
	}

	if !board.Equals(&board_rotcw) {
		t.Error("Board should be equal to its rotational equvalent")
	}

	if !board.Equals(board.Rotate()) {
		t.Error("Board should be equal to its rotational equvalent")
	}

	var board_empty GameBoard
	if board.Equals(&board_empty) {
		t.Error("Board should not be equal to empty board")
	}
}
