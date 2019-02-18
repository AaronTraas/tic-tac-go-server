package tictacgo_test

import (
	. "github.com/AaronTraas/tic-tac-go-server"
	"testing"
)

var firstMoveBoard GameBoard = GameBoard{
	{Empty, Empty, Empty},
	{Empty, X, Empty},
	{Empty, Empty, Empty},
}
var secondMoveBoard GameBoard = GameBoard{
	{O, Empty, Empty},
	{Empty, X, Empty},
	{Empty, Empty, Empty},
}

var partial_board GameBoard = GameBoard{
	{Empty, X, O},
	{O, O, X},
	{X, O, Empty},
}

var newGameState GameState = GameState{Board: emptyBoard}

var firstMoveGameState GameState = GameState{
	Board: firstMoveBoard,
	Move:  GameMove{Row: 1, Col: 1, MoveType: X},
}
var secondMoveGameState GameState = GameState{
	Board: secondMoveBoard,
	Move:  GameMove{Row: 0, Col: 0, MoveType: O},
}

var newGame TicTacToeGame
var gameInProgress TicTacToeGame = TicTacToeGame{
	ID:     "myGameID",
	States: []GameState{firstMoveGameState, secondMoveGameState},
}

func TestIsLegalMove(t *testing.T) {
	if partial_board.IsLegalMove(GameMove{Row: 0, Col: 0, MoveType: X}) != true {
		t.Error("Empty space should be legal")
	}

	if partial_board.IsLegalMove(GameMove{Row: 1, Col: 0, MoveType: X}) != false {
		t.Error("Filled space should not be legal")
	}

	if partial_board.IsLegalMove(GameMove{Row: -1, Col: 0, MoveType: X}) != false {
		t.Error("Out of range row should not be legal")
	}

	if partial_board.IsLegalMove(GameMove{Row: GAME_BOARD_SIZE, Col: 0, MoveType: X}) != false {
		t.Error("Out of range row should not be legal")
	}

	if partial_board.IsLegalMove(GameMove{Row: 0, Col: -1, MoveType: X}) != false {
		t.Error("Out of range column should not be legal")
	}

	if partial_board.IsLegalMove(GameMove{Row: 0, Col: GAME_BOARD_SIZE, MoveType: X}) != false {
		t.Error("Out of range column should not be legal")
	}

	if partial_board.IsLegalMove(GameMove{Row: 0, Col: 0}) != false {
		t.Error("Empty move type is not legal")
	}
}

func TestPlaceMove(t *testing.T) {
	nextState, err := newGameState.PlaceMove(GameMove{Row: 1, Col: 1, MoveType: X})
	if err != nil {
		t.Error("First move of X to 1, 1 on empty board should be legal")
	} else if *nextState != firstMoveGameState {
		t.Error("First move of X to 1, 1 should result in correct state")
	}

	oobState, err := newGameState.PlaceMove(GameMove{Row: -1, Col: 1, MoveType: X})
	if err == nil {
		t.Error("Should be out of bounds")
	}
	if oobState != nil {
		t.Error("Invalid move should result nil being returned")
	}

	fullState, err := firstMoveGameState.PlaceMove(GameMove{Row: 1, Col: 1, MoveType: O})
	if err == nil {
		t.Error("Piece should already be in this location")
	}
	if fullState != nil {
		t.Error("Invalid move should result nil being returned")
	}
}

func TestMove(t *testing.T) {
	moveX22 := GameMove{Row: 2, Col: 2, MoveType: X}
	game, err := gameInProgress.Move(moveX22)
	if err != nil {
		t.Error("Should be a valid move")
	}
	numStates := len(game.States)
	if numStates != 3 {
		t.Error("There were 2 previous states. Performing a legal move should result in 3 states.")
	}
	if game.States[numStates-1].Move != moveX22 {
		t.Error("Move in the last state should be equivalent to move made.")
	}

	badGame, err := gameInProgress.Move(GameMove{Row: -1, Col: 1, MoveType: X})
	if err == nil {
		t.Error("Should be an invalid move.")
	}
	if len(badGame.States) != 2 {
		println("STATES: ", len(badGame.States))
		t.Error("There were 2 previous states. Performing an illegal move should not change this.")
	}
}
