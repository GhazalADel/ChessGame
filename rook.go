package ChessGame

type Rook struct {
	Position Cell
	Color    chessPieceColor

	board *Board
}

func (rook *Rook) getType() chessPieceType {
	return typeRook
}

func (rook *Rook) GetPosition() *Cell {
	return &rook.Position
}

func (rook *Rook) GetColor() chessPieceColor {
	return rook.Color
}

func (rook *Rook) attachToBoard(board *Board) {
	rook.board = board
}

func (rook *Rook) GetAvailableMoves() []Cell {
	moves := []Cell{}
	currentX := rook.Position.X
	currentY := rook.Position.Y
	//vertical - up
	nextY := currentY - 1
	for nextY >= 0 {
		cell := Cell{X: currentX, Y: nextY}
		replace := rook.board.GetPieceOnCell(cell)
		if replace == nil {
			validateAndAddMove(moves, rook, replace, cell)
		} else if isEnemy(rook, replace) && replace.getType() != 6 {
			validateAndAddMove(moves, rook, replace, cell)
			break
		} else {
			break
		}
		nextY -= 1
	}
	//vertical - down
	nextY = currentY + 1
	for nextY <= 7 {
		cell := Cell{X: currentX, Y: nextY}
		replace := rook.board.GetPieceOnCell(cell)
		if replace == nil {
			validateAndAddMove(moves, rook, replace, cell)
		} else if isEnemy(rook, replace) && replace.getType() != 6 {
			validateAndAddMove(moves, rook, replace, cell)
			break
		} else {
			break
		}
		nextY += 1
	}
	//horizontal - left
	nextX := currentX - 1
	for nextX >= 0 {
		cell := Cell{X: nextX, Y: currentY}
		replace := rook.board.GetPieceOnCell(cell)
		if replace == nil {
			validateAndAddMove(moves, rook, replace, cell)
		} else if isEnemy(rook, replace) && replace.getType() != 6 {
			validateAndAddMove(moves, rook, replace, cell)
			break
		} else {
			break
		}
		nextX -= 1
	}
	//horizontal - right
	nextX = currentX + 1
	for nextX <= 7 {
		cell := Cell{X: nextX, Y: currentY}
		replace := rook.board.GetPieceOnCell(cell)
		if replace == nil {
			validateAndAddMove(moves, rook, replace, cell)
		} else if isEnemy(rook, replace) && replace.getType() != 6 {
			validateAndAddMove(moves, rook, replace, cell)
			break
		} else {
			break
		}
		nextX += 1
	}
	return moves
}

func (rook *Rook) moveTo(cell Cell) {
	rook.Position.update(cell)
}
