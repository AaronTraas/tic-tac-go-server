//+build !test

package tictacgo

func (board *GameBoard) DebugPrint() {
	println(board[0][0], "|", board[0][1], "|", board[0][2])
	println("---------")
	println(board[1][0], "|", board[1][1], "|", board[1][2])
	println("---------")
	println(board[2][0], "|", board[2][1], "|", board[2][2])
	println()
}
