package ChessGame

type Queen struct {
	Position Cell
	Color    chessPieceColor

	board *Board
}

func (queen *Queen) getType() chessPieceType {
	return typeQueen
}

func (queen *Queen) GetPosition() *Cell {
	return &queen.Position
}

func (queen *Queen) GetColor() chessPieceColor {
	return queen.Color
}

func (queen *Queen) attachToBoard(board *Board) {
	queen.board = board
}

func (queen *Queen) GetAvailableMoves() []Cell {
	return nil
}

func (queen *Queen) moveTo(cell Cell) {
	queen.Position.update(cell)
}
