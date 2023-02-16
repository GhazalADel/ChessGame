package ChessGame

type Bishop struct {
	Position Cell
	Color    chessPieceColor

	board *Board
}

func (bishop *Bishop) getType() chessPieceType {
	return typeBishop
}

func (bishop *Bishop) getPosition() *Cell {
	return &bishop.Position
}

func (bishop *Bishop) getColor() chessPieceColor {
	return bishop.Color
}

func (bishop *Bishop) attachToBoard(board *Board) {
	bishop.board = board
}
