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
	moves := make([]Cell, 0, 4)
	paths := [8][2]int{
		{-1, -2}, {+1, -2},
		{-2, -1}, {+2, -1},
		{-2, +1}, {+2, +1},
		{-1, +2}, {+1, +2},
	}

	for _, path := range paths {
		cell := Cell{X: knight.Position.X + path[0], Y: knight.Position.Y + path[1]}
		replace := knight.board.GetPieceOnCell(cell)
		validateAndAddMove(&moves, knight, replace, cell)

	}
	return moves
}

func (knight *Knight) moveTo(cell Cell) {
	knight.Position.update(cell)
}
