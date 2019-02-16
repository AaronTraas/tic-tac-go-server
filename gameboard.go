package main

const SIDE = 3

type GameTile string

const (
    Empty GameTile = ""
    X     GameTile = "X"
    O     GameTile = "O"
)

type GameBoard [SIDE][SIDE]GameTile

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

/*
 * Tests the equivalence of two game boards. Game boards that are
 * rotational transforms of each other are deemed equivalent
 */
func (board1 *GameBoard) Equals(board2 *GameBoard) bool {
	if (*board1 == *board2) ||
	   (*board1 == *board2.Rotate()) ||
	   (*board1 == *board2.Rotate().Rotate()) ||
	   (*board1 == *board2.Rotate().Rotate().Rotate()) {
	   	return true
	} else {
		return false
	}
}

