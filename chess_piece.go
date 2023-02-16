package ChessGame

type ChessPieceType int
type ChessPiece interface{}

const (
	_ChessPieceType = iota
	Pawn
	Rook
	Knight
	Bishop
	Queen
	King
)
