package ChessGame

type King struct {
	Position Cell
	Color    chessPieceColor

	board       *Board
	TmpPosition Cell
}

func (king *King) getType() chessPieceType {
	return typeKing
}

func (king *King) GetPosition() *Cell {
	if &king.TmpPosition != nil {
		return &king.TmpPosition
	}
	return &king.Position
}

func (king *King) GetColor() chessPieceColor {
	return king.Color
}

func (king *King) attachToBoard(board *Board) {
	king.board = board
}
func (king *King) setTmpPosition(cell *Cell) {
	king.TmpPosition = *cell
}

func (king *King) GetAvailableMoves() []Cell {
	moves := make([]Cell, 0, 4)
	paths := [8][2]int{
		{0, +1}, {0, -1},
		{+1, 0}, {-1, 0},
		{-1, -1}, {-1, +1},
		{+1, +1}, {+1, -1},
	}

	for _, path := range paths {
		cell := Cell{X: king.Position.X + path[0], Y: king.Position.Y + path[1]}
		replace := king.board.GetPieceOnCell(cell)
		validateAndAddMove(&moves, king, replace, cell, *king.board)
	}

	if !king.Position.hasMoved {
		if king.canCastling(true) {
			moves = append(moves, Cell{X: 6, Y: king.Position.Y})
		}

		if king.canCastling(false) {
			moves = append(moves, Cell{X: 2, Y: king.Position.Y})
		}
	}
	return moves
}

func (king *King) moveTo(cell Cell) {
	king.Position.update(cell)
}

func (king *King) canCastling(kingSide bool) bool {
	if isCheck(king.board, king.GetColor()) {
		return false
	}

	var rookX int
	if kingSide {
		rookX = BoardGrid - 1
	} else {
		rookX = 0
	}

	piece := king.board.GetPieceOnCell(Cell{X: rookX, Y: king.Position.Y})
	if piece != nil && piece.getType() == typeRook && !piece.GetPosition().hasMoved {
		var xs [2]int
		if kingSide {
			xs = [2]int{5, 6}
		} else {
			xs = [2]int{2, 3}
		}

		for _, x := range xs {
			cell := Cell{X: x, Y: king.Position.Y}
			if king.board.GetPieceOnCell(cell) != nil ||
				isUnderAttack(king.board, king.GetColor(), &cell) {
				return false
			}
		}
	}
	return true
}
