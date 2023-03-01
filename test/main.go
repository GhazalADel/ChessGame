package main

import (
	ChessGame "Chess"
	ChessGameCLI "Chess/cli"
)

func main() {
	var ui ChessGame.BoardInterface = &ChessGameCLI.ChessGameCLI{
		Board: ChessGame.CreateBoard(),
	}
	//ui.SelectCell(ChessGame.Cell{X: 0, Y: 1})
	ui.ShowBoard()
	board := ui.GetBoard()
	board.Turn = true
	//for {
	//
	//}
}
