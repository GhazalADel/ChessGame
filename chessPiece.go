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

func validateAndAddMove(moves *[]Cell, piece, replace ChessPiece, cell Cell, board *Board) bool {
	if cell.isUndefined() {
		return false
	}

	if replace == nil || replace.GetColor() != piece.GetColor() {
		if board.isSimulating || !simulateNextMove(piece, replace, cell, board) {
			*moves = append(*moves, cell)
		}
	}
	return replace == nil
}

func simulateNextMove(piece, replace ChessPiece, cell Cell, board *Board) bool {
	tmpBoard := Board{}
	oldCell := Cell{
		X:        piece.GetPosition().X,
		Y:        piece.GetPosition().Y,
		hasMoved: piece.GetPosition().hasMoved,
	}

	piece.moveTo(cell)

	doOnAllPieces(board, func(bp ChessPiece) {
		if bp != replace {
			tmpBoard.AddPiece(bp)
		}
	})

	check := isCheck(&tmpBoard, piece.GetColor())

	piece.GetPosition().update(oldCell)
	piece.GetPosition().hasMoved = oldCell.hasMoved

	doOnAllPieces(board, func(bp ChessPiece) {
		bp.attachToBoard(board)
	})

	return check
}

func doOnAllPieces(board *Board, action func(ChessPiece)) {
	pieces := [2]chessPiecesData{
		board.BlackPieces.data,
		board.WhitePieces.data,
	}

	for _, data := range pieces {
		for _, value := range data {
			for _, piece := range value {
				action(piece)
			}
		}
	}
}

func isEnemy(piece1, piece2 ChessPiece) bool {
	if piece2 == nil {
		return false
	}

	return piece1.GetColor() != piece2.GetColor()
}

func getKing(board *Board, color chessPieceColor) (king ChessPiece) {
	if color == White {
		king = board.BlackPieces.King()
	} else {
		king = board.WhitePieces.King()
	}
	return
}

func getPieces(board *Board, color chessPieceColor) (data chessPiecesData) {
	if color == White {
		data = board.BlackPieces.data
	} else {
		data = board.WhitePieces.data
	}
	return
}

func isUnderAttack(board *Board, color chessPieceColor, cell *Cell) bool {
	board.isSimulating = true
	data := getPieces(board, color)

	for _, value := range data {
		for _, piece := range value {
			moves := piece.GetAvailableMoves()
			for _, move := range moves {
				if move.Equals(cell) {
					return true
				}
			}
		}
	}
	board.isSimulating = false
	return false
}

func isCheck(board *Board, color chessPieceColor) bool {
	return isUnderAttack(board, color, getKing(board, color).GetPosition())
}

func isMate(piece ChessPiece, board *Board) bool {
	if isUnderAttack(board, piece.GetColor(), getKing(board, piece.GetColor()).GetPosition()) {
		for _, move := range getKing(board, piece.GetColor()).GetAvailableMoves() {
			if !isUnderAttack(board, piece.GetColor(), &move) {
				return false
			}
		}
		return true
	}
	return false
}
