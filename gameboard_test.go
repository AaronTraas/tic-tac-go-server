package tictacgo_test

import (
	. "github.com/AaronTraas/tic-tac-go-server"
	"testing"
)

var emptyBoard GameBoard

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
		t.Error("Board should be equal to its rotational eqiuvalent")
	}

	if !board.Equals(board.Rotate()) {
		t.Error("Board should be equal to its rotational eqiuvalent")
	}

	if !board.Equals(board.FlipH()) {
		t.Error("Board should be equal to its horizontal reflection")
	}

	if !board.Equals(board.FlipV()) {
		t.Error("Board should be equal to its vertical reflection")
	}

	if !board.Equals(board.Rotate().Rotate().FlipH().FlipV()) {
		t.Error("Board should be equal to its rotational and reflectional eqiuvalent")
	}

	var board_empty GameBoard
	if board.Equals(&board_empty) {
		t.Error("Board should not be equal to empty board")
	}
}

func TestSerialize(t *testing.T) {
	if emptyBoard.Serialize() != "EEEEEEEEE" {
		t.Error("Empty board should serialize to 'EEEEEEEEE'")
	}

	if board.Serialize() != "XXOOOXXOO" {
		t.Error("Board serialization result incorrect")
	}
}

func TestDeserialize(t *testing.T) {
	var newBoard GameBoard

	newEmptyBoard, err := newBoard.Deserialize("EEEEEEEEE")
	if err != nil {
		t.Error("'EEEEEEEEE' is valid serialized board string")
	}
	if *newEmptyBoard != emptyBoard {
		t.Error("'EEEEEEEEE' should validly deserialize")
	}

	newFullBoard, err := newBoard.Deserialize("XXOOOXXOO")
	if err != nil {
		t.Error("'XXOOOXXOO' is valid serialized board string")
	}
	if *newFullBoard != board {
		t.Error("'XXOOOXXOO' should validly deserialize")
	}

	_, err = newBoard.Deserialize("XXXO OXXX")
	if err == nil {
		t.Error("'XXXO OXXX' is not a valid serialized board string")
	}

	_, err = newBoard.Deserialize("X")
	if err == nil {
		t.Error("'X' is not a valid serialized board string")
	}

	_, err = newBoard.Deserialize("XXXXXXXXXXXXXXXXXXXXXXXXX")
	if err == nil {
		t.Error("'XXXXXXXXXXXXXXXXXXXXXXXXX' is not a valid serialized board string")
	}
}
