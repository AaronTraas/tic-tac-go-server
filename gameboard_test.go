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

func TestRotate(t *testing.T) {

	if *board.Rotate() != board_rotcw {
		t.Error("Rotation clockwise once failed")
	}

	if *board.Rotate().Rotate().Rotate().Rotate() != board {
		t.Error("Rotation clockwise 4 times should result in the same board")
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
