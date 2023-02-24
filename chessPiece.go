package ChessGame

type chessPieceType int

const (
	_ chessPieceType = iota
	typePawn
	typeRook
	typeKnight
	typeBishop
	typeQueen
	typeKing
)

type ChessPiece interface {
	GetPosition() *Cell
	GetColor() chessPieceColor
	GetAvailableMoves() []Cell

	moveTo(cell Cell)
	getType() chessPieceType
	attachToBoard(board *Board)
}

func validateAndAddMove(moves *[]Cell, piece, replace ChessPiece, cell Cell) bool {
	if replace == nil || replace.GetColor() != piece.GetColor() {
		// TODO: validate checkmate here
		*moves = append(*moves, cell)
	}

	return replace == nil
}

func isEnemy(piece1, piece2 ChessPiece) bool {
	if piece2 == nil {
		return false
	}

	return piece1.GetColor() != piece2.GetColor()
}
func Check(board *Board, color chessPieceColor, cell *Cell) bool {
	isChecked := false
	if color {
		for key, value := range board.BlackPieces.data {
			if key != 6 {
				for _, piece := range value {
					moves := piece.GetAvailableMoves()
					for _, move := range moves {
						if board.GetPieceOnCell(move).GetPosition().X == cell.X && board.GetPieceOnCell(move).GetPosition().Y == cell.Y {
							isChecked = true
							break
						}
					}
				}
			}
		}
	} else {
		for key, value := range board.WhitePieces.data {
			if key != 6 {
				for _, piece := range value {
					moves := piece.GetAvailableMoves()
					for _, move := range moves {
						if board.GetPieceOnCell(move).GetPosition().X == cell.X && board.GetPieceOnCell(move).GetPosition().Y == cell.Y {
							isChecked = true
							break
						}
					}
				}
			}
		}
	}
	return isChecked
}
