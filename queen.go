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
	moves := make([]Cell, 0, 4)
	paths := [8][2]int{
		{-1, 0}, {+1, 0},
		{0, -1}, {0, +1},
		{+1, -1}, {-1, -1},
		{-1, +1}, {+1, +1},
	}

	for _, path := range paths {
		cell := Cell{X: queen.Position.X + path[0], Y: queen.Position.Y + path[1]}

		for !cell.isUndefined() &&
			validateAndAddMove(&moves, queen, queen.board.GetPieceOnCell(cell), cell, queen.board) {
			cell.X += path[0]
			cell.Y += path[1]
		}
	}
	return moves
}

func (queen *Queen) MoveTo(cell Cell) {
	queen.Position.update(cell)
}
