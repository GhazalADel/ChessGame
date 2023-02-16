package ChessGame

type ChessPieces map[ChessPieceType][]ChessPiece

func (pieces ChessPieces) King() ChessPiece {
	return pieces[King][0]
}

type Board struct {
	blackPieces ChessPieces
	whitePieces ChessPieces
	turnWhite   bool
}
