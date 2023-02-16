package ChessGame

type Queen struct {
	Position Cell
	Color    chessPieceColor

	board *Board
}

func (queen *Queen) getType() chessPieceType {
	return typeQueen
}

func (queen *Queen) getPosition() *Cell {
	return &queen.Position
}

func (queen *Queen) getColor() chessPieceColor {
	return queen.Color
}

func (queen *Queen) attachToBoard(board *Board) {
	queen.board = board
}
