package ChessGame

type King struct {
	Position Cell
	Color    chessPieceColor

	board *Board
}

func (king *King) getType() chessPieceType {
	return typeKing
}

func (king *King) getPosition() *Cell {
	return &king.Position
}

func (king *King) getColor() chessPieceColor {
	return king.Color
}

func (king *King) attachToBoard(board *Board) {
	king.board = board
}
