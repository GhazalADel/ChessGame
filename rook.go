package ChessGame

type Rook struct {
	Position Cell
	Color    chessPieceColor

	board *Board
}

func (rook *Rook) getType() chessPieceType {
	return typeRook
}

func (rook *Rook) getPosition() *Cell {
	return &rook.Position
}

func (rook *Rook) getColor() chessPieceColor {
	return rook.Color
}

func (rook *Rook) attachToBoard(board *Board) {
	rook.board = board
}
