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
