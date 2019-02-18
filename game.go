package tictacgo

import (
	"errors"
)

type GameMove struct {
	Row      int      `json:"row"`
	Col      int      `json:"col"`
	MoveType GameTile `json:"type"`
}

type GameState struct {
	Board GameBoard `json:"board"`
	Move  GameMove  `json:"move"`
}

type TicTacToeGame struct {
	ID     string      `json:"ID"`
	States []GameState `json:"states"`
}

func (board *GameBoard) IsLegalMove(move GameMove) bool {
	if (move.Row >= 0) && (move.Row < GAME_BOARD_SIZE) &&
		(move.Col >= 0) && (move.Col < GAME_BOARD_SIZE) &&
		(move.MoveType == X || move.MoveType == O) &&
		(board[move.Row][move.Col] == Empty) {
		return true
	}
	return false
}

func (oldState *GameState) PlaceMove(move GameMove) (*GameState, error) {
	if oldState.Board.IsLegalMove(move) {
		newBoard := oldState.Board
		newBoard[move.Row][move.Col] = move.MoveType
		var newState *GameState = &GameState{Board: newBoard, Move: move}
		return newState, nil
	}
	return nil, errors.New("invalid move")
}

func (game *TicTacToeGame) Move(move GameMove) (*TicTacToeGame, error) {
	newGame := *game

	lastState := newGame.States[len(newGame.States)-1]
	newState, err := lastState.PlaceMove(move)
	if err != nil {
		return &newGame, err
	}
	newGame.States = append(newGame.States, *newState)
	return &newGame, nil
}
