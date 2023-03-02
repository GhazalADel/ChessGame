package main

import (
	ChessGame "Chess"
	ChessGameCLI "Chess/cli"
	"fmt"
	"strconv"
	"strings"
)

func resetCellColors(ui ChessGame.BoardInterface) {
	ui.SetMoves(nil)
	ui.SetSelectedCell(ChessGame.Cell{X: -1, Y: -1})
}
func checkInputValidation(inputCell string) bool {
	if len(inputCell) != 2 || inputCell[0] < 65 || inputCell[0] > (65+ChessGame.BoardGrid-1) || string(inputCell[1]) < "1" || string(inputCell[1]) > strconv.FormatInt(int64(ChessGame.BoardGrid), 10) {
		return false
	}
	return true
}
func main() {
	var ui ChessGame.BoardInterface = &ChessGameCLI.ChessGameCLI{
		Board: ChessGame.CreateBoard(),
	}

	board := ui.GetBoard()
	board.Turn = true

	for {
		fmt.Println("\n\n")
		resetCellColors(ui)
		ui.ShowBoard()
		fmt.Println("\n")
		chooseAnotherPiece := false

		turn := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 32, "Black's turn")
		if board.Turn {
			turn = fmt.Sprintf("\x1b[%dm%s\x1b[0m", 32, "White's turn")
		}
		fmt.Println(turn)
		fmt.Print(fmt.Sprintf("\x1b[%dm%s\x1b[0m", 32, "Choose your piece"))
		fmt.Print(fmt.Sprintf("\x1b[%dm%s\x1b[0m", 31, " (FORMAT : [ROW][COLUMN])\n"))

	selectMove:
		for {
			var selectedCellStr string
			fmt.Scanln(&selectedCellStr)
			selectedCellStr = strings.ToUpper(selectedCellStr)
			if !checkInputValidation(selectedCellStr) {
				fmt.Println(fmt.Sprintf("\x1b[%dm%s\x1b[0m", 31, "WRONG cell format!"))
				continue
			}
			cell := ChessGame.Cell{X: int(selectedCellStr[0] - 65), Y: int(selectedCellStr[1]) - 49}
			selectedPiece := ui.GetBoard().GetPieceOnCell(cell)
			if selectedPiece == nil {
				fmt.Println(fmt.Sprintf("\x1b[%dm%s\x1b[0m", 31, "This cell is empty!"))
				continue
			}
			if selectedPiece.GetColor() != board.Turn {
				fmt.Println(fmt.Sprintf("\x1b[%dm%s\x1b[0m", 31, "Choose your piece!"))
				continue
			}
			moves := selectedPiece.GetAvailableMoves()
			if len(moves) == 0 {
				fmt.Println(fmt.Sprintf("\x1b[%dm%s\x1b[0m", 31, "you can't move this piece!"))
				continue
			}
			ui.SelectCell(cell)
			ui.ShowBoard()
			fmt.Println("\n\n")
			fmt.Print(fmt.Sprintf("\x1b[%dm%s\x1b[0m", 32, "Select destination"))
			fmt.Print(fmt.Sprintf("\x1b[%dm%s\x1b[0m", 31, " (FORMAT : [ROW][COLUMN]) or write `back`\n"))
			var destinationCellStr string
			for {
				fmt.Scanln(&destinationCellStr)
				if destinationCellStr == "back" {
					chooseAnotherPiece = true
					break selectMove
				}
				if !checkInputValidation(destinationCellStr) {
					fmt.Println(fmt.Sprintf("\x1b[%dm%s\x1b[0m", 31, "WRONG destination cell format!"))
					continue
				}
				destinationCell := ChessGame.Cell{X: int(destinationCellStr[0] - 65), Y: int(destinationCellStr[1]) - 49}
				validMove := false
				for _, move := range moves {
					if move.Equals(&destinationCell) {
						board.SetPieceOnCell(move, selectedPiece)
						board.SetPieceOnCell(cell, nil)
						selectedPiece.MoveTo(move)
						validMove = true
						break selectMove
					}
				}
				if !validMove {
					fmt.Println(fmt.Sprintf("\x1b[%dm%s\x1b[0m", 31, "you can't go to this cell!"))
					continue
				}
			}
		}
		if chooseAnotherPiece {
			resetCellColors(ui)
			ui.ShowBoard()
			continue
		}

		board.Turn = !board.Turn
		// TODO : check end game
	}
}
