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
	getType() chessPieceType
	getPosition() *Cell
	getColor() chessPieceColor

	attachToBoard(board *Board)
}
