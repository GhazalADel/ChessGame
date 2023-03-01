package ChessGame

type Bishop struct {
	Position Cell
	Color    chessPieceColor

	board       *Board
	TmpPosition Cell
}

func (bishop *Bishop) getType() chessPieceType {
	return typeBishop
}

func (bishop *Bishop) GetPosition() *Cell {
	if &bishop.TmpPosition != nil {
		return &bishop.TmpPosition
	}
	return &bishop.Position
}

func (bishop *Bishop) GetColor() chessPieceColor {
	return bishop.Color
}

func (bishop *Bishop) attachToBoard(board *Board) {
	bishop.board = board
}
func (bishop *Bishop) setTmpPosition(cell *Cell) {
	bishop.TmpPosition = *cell
}

func (bishop *Bishop) GetAvailableMoves() []Cell {
	moves := make([]Cell, 0, 4)
	paths := [4][2]int{
		{+1, -1}, {-1, -1},
		{-1, +1}, {+1, +1},
	}
	for _, path := range paths {
		cell := Cell{X: bishop.Position.X + path[0], Y: bishop.Position.Y + path[1]}
		for !cell.isUndefined() &&
			validateAndAddMove(&moves, bishop, bishop.board.GetPieceOnCell(cell), cell, *bishop.board) {
			cell.X += path[0]
			cell.Y += path[1]
		}
	}
	return moves
}

func (bishop *Bishop) moveTo(cell Cell) {
	bishop.Position.update(cell)
}
