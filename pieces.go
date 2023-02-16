package ChessGame

type chessPieces map[chessPieceType][]ChessPiece

func (pieces chessPieces) King() ChessPiece {
	return pieces[typeKing][0]
}

func (pieces chessPieces) Add(piece ChessPiece) {
	pieceType := piece.getType()

	if _, contains := pieces[pieceType]; !contains {
		pieces[pieceType] = make([]ChessPiece, 0, 1)
	}

	if pieceType == typeKing && len(pieces[typeKing]) != 0 {
		panic("You can only have one King!")
	}

	pieces[pieceType] = append(pieces[pieceType], piece)
}
