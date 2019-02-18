package tictacgo

import (
	"errors"
)

const GAME_BOARD_SIZE = 3

type GameTile string

const (
	Empty GameTile = ""
	X     GameTile = "X"
	O     GameTile = "O"
)

type GameBoard [GAME_BOARD_SIZE][GAME_BOARD_SIZE]GameTile

/**
 * Rotates a board clockwise
 */
func (board *GameBoard) Rotate() *GameBoard {
	var newBoard = GameBoard{
		{board[2][0], board[1][0], board[0][0]},
		{board[2][1], board[1][1], board[0][1]},
		{board[2][2], board[1][2], board[0][2]},
	}

	return &newBoard
}

/**
 * Flip board horizontally
 */
func (board *GameBoard) FlipH() *GameBoard {
	var newBoard = GameBoard{
		{board[0][2], board[0][1], board[0][0]},
		{board[1][2], board[1][1], board[1][0]},
		{board[2][2], board[2][1], board[2][0]},
	}

	return &newBoard
}

/**
 * Flip board vertically
 */
func (board *GameBoard) FlipV() *GameBoard {
	var newBoard = GameBoard{
		{board[2][0], board[2][1], board[2][2]},
		{board[1][0], board[1][1], board[1][2]},
		{board[0][0], board[0][1], board[0][2]},
	}

	return &newBoard
}

/*
 * Tests the equivalence of two game boards. Game boards that are
 * rotation/reflection transforms of each other are deemed equivalent
 */
func (board1 *GameBoard) Equals(board2 *GameBoard) bool {
	board_rot := board2

	var board_transforms []GameBoard

	for i := 0; i < 4; i++ {
		board_transforms = append(board_transforms, *board_rot)
		board_transforms = append(board_transforms, *board_rot.FlipH())
		board_transforms = append(board_transforms, *board_rot.FlipV())
		board_transforms = append(board_transforms, *board_rot.FlipH().FlipV())
		board_rot = board_rot.Rotate()
	}

	for _, board := range board_transforms {
		if *board1 == board {
			return true
		}
	}
	return false
}

func (board *GameBoard) Serialize() string {
	var s string = ""
	for _, row := range board {
		for _, col := range row {
			switch col {
			case Empty:
				s += "E"
			case X:
				s += "X"
			case O:
				s += "O"
			}
		}
	}
	return s
}

func (board *GameBoard) Deserialize(s string) (*GameBoard, error) {
	if len(s) != GAME_BOARD_SIZE*GAME_BOARD_SIZE {
		return nil, errors.New("String " + s + " cannot be deserialized into GameBoard.")
	}

	ptr := 0
	for row := 0; row < GAME_BOARD_SIZE; row++ {
		for col := 0; col < GAME_BOARD_SIZE; col++ {
			c := string(s[ptr])
			ptr++
			switch c {
			case "E":
				board[row][col] = Empty
			case "X":
				board[row][col] = X
			case "O":
				board[row][col] = O
			default:
				return nil, errors.New("String " + s + " cannot be deserialized into GameBoard.")
			}
		}
	}
	return board, nil
}
