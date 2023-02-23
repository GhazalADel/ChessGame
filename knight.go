package ChessGame

type Knight struct {
	Position Cell
	Color    chessPieceColor

	board *Board
}

func (knight *Knight) getType() chessPieceType {
	return typeKnight
}

func (knight *Knight) GetPosition() *Cell {
	return &knight.Position
}

func (knight *Knight) GetColor() chessPieceColor {
	return knight.Color
}

func (knight *Knight) attachToBoard(board *Board) {
	knight.board = board
}

func (knight *Knight) GetAvailableMoves() []Cell {
	return nil
}

func (knight *Knight) moveTo(cell Cell) {
	knight.Position.update(cell)
}
