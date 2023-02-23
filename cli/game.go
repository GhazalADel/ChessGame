package main

import (
	ChessGame "Chess"
	"fmt"
)

func getPieceSymbol(piece ChessGame.ChessPiece) string {
	symbol := ""
	if piece == nil {
		symbol = ""
	} else {
		switch piece.(type) {
		case *ChessGame.Pawn:
			if piece.GetColor().IsWhite() {
				symbol = "♙"
			} else {
				symbol = "♟"
			}
		case *ChessGame.Knight:
			if piece.GetColor().IsWhite() {
				symbol = "♘"
			} else {
				symbol = "♞"
			}
		case *ChessGame.Bishop:
			if piece.GetColor().IsWhite() {
				symbol = "♗"
			} else {
				symbol = "♝"
			}
		case *ChessGame.Rook:
			if piece.GetColor().IsWhite() {
				symbol = "♖"
			} else {
				symbol = "♜"
			}
		case *ChessGame.Queen:
			if piece.GetColor().IsWhite() {
				symbol = "♕"
			} else {
				symbol = "♛"
			}
		case *ChessGame.King:
			if piece.GetColor().IsWhite() {
				symbol = "♔"
			} else {
				symbol = "♚"
			}
		default:
			symbol = ""
		}
	}
	return symbol
}

func showBoard(board *ChessGame.Board) {
	fmt.Println("┌────┐────┐────┐────┐────┐────┐────┐────┐")
	for i := 0; i < 7; i++ {
		var pieceOnRow [8]string
		for j := 0; j < 8; j++ {
			c := ChessGame.Cell{
				X: j,
				Y: i,
			}
			piece := board.GetPieceOnCell(c)
			pieceOnRow[j] = getPieceSymbol(piece)

		}
		fmt.Printf("│ %s  │ %s  │ %s  │ %s  │ %s  │ %s  │ %s  │ %s  │\n", pieceOnRow[0], pieceOnRow[1], pieceOnRow[2], pieceOnRow[3], pieceOnRow[4], pieceOnRow[5], pieceOnRow[6], pieceOnRow[7])
		fmt.Printf("├────├────├────├────├────├────├────├────┤\n")
	}
	var pieceOnRow [8]string
	for j := 0; j < 8; j++ {
		c := ChessGame.Cell{
			X: j,
			Y: 7,
		}
		piece := board.GetPieceOnCell(c)
		pieceOnRow[j] = getPieceSymbol(piece)
		fmt.Printf("│ %s  │ %s  │ %s  │ %s  │ %s  │ %s  │ %s  │ %s  │\n", pieceOnRow[0], pieceOnRow[1], pieceOnRow[2], pieceOnRow[3], pieceOnRow[4], pieceOnRow[5], pieceOnRow[6], pieceOnRow[7])
		fmt.Println("└────└────└────└────└────└────└────└────┘")
	}
}

func main() {

}
