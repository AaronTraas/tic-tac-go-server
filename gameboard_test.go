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

	if !board.Equals(&board_rotcw) {
		t.Error("Rotation clockwise once failed")
	}

}
