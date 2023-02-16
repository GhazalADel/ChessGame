package ChessGame

type Knight struct {
	Position Cell
	Color    chessPieceColor

	board *Board
}

func (knight *Knight) getType() chessPieceType {
	return typeKnight
}

func (knight *Knight) getPosition() *Cell {
	return &knight.Position
}

func (knight *Knight) getColor() chessPieceColor {
	return knight.Color
}

func (knight *Knight) attachToBoard(board *Board) {
	knight.board = board
}
