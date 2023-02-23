package ChessGame

type King struct {
	Position Cell
	Color    chessPieceColor

	board *Board
}

func (king *King) getType() chessPieceType {
	return typeKing
}

func (king *King) GetPosition() *Cell {
	return &king.Position
}

func (king *King) GetColor() chessPieceColor {
	return king.Color
}

func (king *King) attachToBoard(board *Board) {
	king.board = board
}

func (king *King) GetAvailableMoves() []Cell {
	return nil
}

func (king *King) moveTo(cell Cell) {
	king.Position.update(cell)
}
