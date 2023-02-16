package ChessGame

type chessPiecesData map[chessPieceType][]ChessPiece

type chessPieces struct {
	data  chessPiecesData
	board *Board
}

func (pieces chessPieces) King() ChessPiece {
	return pieces.data[typeKing][0]
}

func (pieces chessPieces) add(piece ChessPiece) {
	if pieces.data == nil {
		pieces.data = make(chessPiecesData)
	}

	pieceType := piece.getType()

	if _, contains := pieces.data[pieceType]; !contains {
		pieces.data[pieceType] = make([]ChessPiece, 0, 1)
	}

	if pieceType == typeKing && len(pieces.data[typeKing]) != 0 {
		panic("You can only have one King!")
	}

	pieces.data[pieceType] = append(pieces.data[pieceType], piece)
}
