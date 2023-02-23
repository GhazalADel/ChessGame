package ChessGame

type Bishop struct {
	Position Cell
	Color    chessPieceColor

	board *Board
}

func (bishop *Bishop) getType() chessPieceType {
	return typeBishop
}

func (bishop *Bishop) GetPosition() *Cell {
	return &bishop.Position
}

func (bishop *Bishop) GetColor() chessPieceColor {
	return bishop.Color
}

func (bishop *Bishop) attachToBoard(board *Board) {
	bishop.board = board
}

func (bishop *Bishop) GetAvailableMoves() []Cell {
	return nil
}

func (bishop *Bishop) moveTo(cell Cell) {
	bishop.Position.update(cell)
}
