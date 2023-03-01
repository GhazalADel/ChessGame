package ChessGameCLI

import (
	ChessGame "Chess"
	"fmt"
	"strconv"
)

type ChessGameCLI struct {
	Board    *ChessGame.Board
	Moves    []ChessGame.Cell
	Selected ChessGame.Cell
}

func GetPieceSymbol(piece ChessGame.ChessPiece) string {
	symbol := " "
	if piece != nil {
		switch piece.(type) {
		case *ChessGame.Pawn:
			if piece.GetColor().IsBlack() {
				symbol = "♙"
			} else {
				symbol = "♟"
			}
		case *ChessGame.Knight:
			if piece.GetColor().IsBlack() {
				symbol = "♘"
			} else {
				symbol = "♞"
			}
		case *ChessGame.Bishop:
			if piece.GetColor().IsBlack() {
				symbol = "♗"
			} else {
				symbol = "♝"
			}
		case *ChessGame.Rook:
			if piece.GetColor().IsBlack() {
				symbol = "♖"
			} else {
				symbol = "♜"
			}
		case *ChessGame.Queen:
			if piece.GetColor().IsBlack() {
				symbol = "♕"
			} else {
				symbol = "♛"
			}
		case *ChessGame.King:
			if piece.GetColor().IsBlack() {
				symbol = "♔"
			} else {
				symbol = "♚"
			}
		}
	}

	return symbol
}

func (cli *ChessGameCLI) ShowBoard() {
	var boardMap [ChessGame.BoardGrid*2 + 1][ChessGame.BoardGrid*5 + 1]any
	var verticlaNum int64 = 8
	var printNum bool = false
	for j := 0; j < len(boardMap); j++ {
		for i := 0; i < len(boardMap[j]); i++ {
			if j == 0 {
				if i == 0 {
					boardMap[j][i] = "┌"
				} else if i%5 == 0 {
					boardMap[j][i] = "┐"
				} else {
					boardMap[j][i] = "─"
				}
			} else if j == len(boardMap)-1 {
				if i == len(boardMap[j])-1 {
					boardMap[j][i] = "┘"
				} else if i%5 == 0 {
					boardMap[j][i] = "└"
				} else {
					boardMap[j][i] = "─"
				}
			} else if j%2 == 0 {
				if i == len(boardMap[j])-1 {
					boardMap[j][i] = "┤"
				} else if i%5 == 0 {
					boardMap[j][i] = "├"
				} else {
					boardMap[j][i] = "─"
				}
			} else {
				if i%5 == 0 {
					boardMap[j][i] = "|"
				} else {
					cell := ChessGame.Cell{X: i / 5, Y: (j - 1) / 2}

					if (i-2)%5 == 0 {
						boardMap[j][i] = GetPieceSymbol(cli.Board.GetPieceOnCell(cell))
					} else {
						boardMap[j][i] = " "
					}

					var color int
					if cell.X%2 == cell.Y%2 {
						color = 47
					} else {
						color = 40
					}

					if cell.Equals(&cli.Selected) {
						color = 44
					} else if cli.Moves != nil {
						for _, m := range cli.Moves {
							if cell.Equals(&m) {
								color = 42
								break
							}
						}
					}
					boardMap[j][i] = fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, boardMap[j][i])
				}
			}
		}
		if printNum {
			numbers := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 33, strconv.FormatInt(int64(verticlaNum), 10))
			fmt.Print(numbers)
			verticlaNum -= 1
			printNum = false
		} else {
			printNum = true
		}
		if printNum {
			fmt.Print("  ")
		} else {
			fmt.Print(" ")
		}
		fmt.Print(boardMap[j][:]...)
		fmt.Println()

	}
	letters := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 33, "    A    B    C    D    E    F    G    H")
	fmt.Print(letters)

}

func (cli *ChessGameCLI) SelectCell(cell ChessGame.Cell) {
	piece := cli.Board.GetPieceOnCell(cell)
	cli.Selected = cell
	if piece != nil {
		cli.Moves = piece.GetAvailableMoves()
	}
}
