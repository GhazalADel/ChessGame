package ChessGame

type Rook struct {
	Position Cell
	Color    chessPieceColor

	board *Board
}

func (rook *Rook) getType() chessPieceType {
	return typeRook
}

func (rook *Rook) GetPosition() *Cell {
	return &rook.Position
}

func (rook *Rook) GetColor() chessPieceColor {
	return rook.Color
}

func (rook *Rook) attachToBoard(board *Board) {
	rook.board = board
}

func (rook *Rook) GetAvailableMoves() []Cell {
	return nil
}

func (rook *Rook) moveTo(cell Cell) {
	rook.Position.update(cell)
}
