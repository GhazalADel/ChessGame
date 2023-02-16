package main

type ChessPieceType int
type ChessPiece interface{}
type ChessPieces map[ChessPieceType]ChessPiece

type Board struct {
	blackPieces ChessPieces
	whitePieces ChessPieces
	turnWhite   bool
}

const (
	_ChessPieceType = iota
	Pawn
	Rook
	Knight
	Bishop
	Queen
	King
)

func main() {

}
