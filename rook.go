package ChessGame

type Rook struct {
	Position Cell
	Color    chessPieceColor

	board       *Board
	TmpPosition Cell
}

func (rook *Rook) getType() chessPieceType {
	return typeRook
}

func (rook *Rook) GetPosition() *Cell {
	if &rook.TmpPosition != nil {
		return &rook.TmpPosition
	}
	return &rook.Position
}

func (rook *Rook) GetColor() chessPieceColor {
	return rook.Color
}

func (rook *Rook) attachToBoard(board *Board) {
	rook.board = board
}
func (rook *Rook) setTmpPosition(cell *Cell) {
	rook.TmpPosition = *cell
}

func (rook *Rook) GetAvailableMoves() []Cell {
	moves := make([]Cell, 0, 4)
	paths := [4][2]int{
		{-1, 0}, {+1, 0},
		{0, -1}, {0, +1},
	}

	for _, path := range paths {
		cell := Cell{X: rook.Position.X + path[0], Y: rook.Position.Y + path[1]}

		for !cell.isUndefined() &&
			validateAndAddMove(&moves, rook, rook.board.GetPieceOnCell(cell), cell, *rook.board) {
			cell.X += path[0]
			cell.Y += path[1]
		}
	}
	return moves
}

func (rook *Rook) moveTo(cell Cell) {
	rook.Position.update(cell)
}
