package ChessGame

type Pawn struct {
	Position Cell
	Color    chessPieceColor

	board       *Board
	enPassant   bool
	TmpPosition Cell
}

func (pawn *Pawn) getType() chessPieceType {
	return typePawn
}

func (pawn *Pawn) GetPosition() *Cell {
	if &pawn.TmpPosition != nil {
		return &pawn.TmpPosition
	}
	return &pawn.Position
}

func (pawn *Pawn) GetColor() chessPieceColor {
	return pawn.Color
}

func (pawn *Pawn) attachToBoard(board *Board) {
	pawn.board = board
}
func (pawn *Pawn) setTmpPosition(cell *Cell) {
	pawn.TmpPosition = *cell
}

func (pawn *Pawn) GetAvailableMoves() []Cell {
	moves := make([]Cell, 0, 4)
	nextY := pawn.Position.Y
	if pawn.Color.IsBlack() {
		nextY++
	} else {
		nextY--
	}

	cell := Cell{X: pawn.Position.X, Y: nextY}
	replace := pawn.board.GetPieceOnCell(cell)
	if replace == nil {
		validateAndAddMove(&moves, pawn, replace, cell, *pawn.board)

		if pawn.isFirstMove() {
			if pawn.Color.IsBlack() {
				cell.Y++
			} else {
				cell.Y--
			}

			replace = pawn.board.GetPieceOnCell(cell)
			if replace == nil {
				validateAndAddMove(&moves, pawn, replace, cell, *pawn.board)
			}
		}
	}

	cell.X = pawn.Position.X + 1
	cell.Y = nextY
	replace = pawn.board.GetPieceOnCell(cell)
	if isEnemy(pawn, replace) {
		validateAndAddMove(&moves, pawn, replace, cell, *pawn.board)
	}

	cell.Y = pawn.Position.Y
	replace = pawn.board.GetPieceOnCell(cell)
	if replace != nil {
		if p, ok := replace.(*Pawn); ok && p.enPassant {
			validateAndAddMove(&moves, pawn, replace, cell, *pawn.board)
		}
	}

	cell.X = pawn.Position.X - 1
	cell.Y = nextY
	replace = pawn.board.GetPieceOnCell(cell)
	if isEnemy(pawn, replace) {
		validateAndAddMove(&moves, pawn, replace, cell, *pawn.board)
	}

	cell.Y = pawn.Position.Y
	replace = pawn.board.GetPieceOnCell(cell)
	if replace != nil {
		if p, ok := replace.(*Pawn); ok && p.enPassant {
			validateAndAddMove(&moves, pawn, replace, cell, *pawn.board)
		}
	}
	return moves
}

func (pawn *Pawn) moveTo(cell Cell) {
	pawn.enPassant = pawn.isFirstMove() &&
		((pawn.Color.IsBlack() && cell.Y == 3) ||
			(pawn.Color.IsWhite() && cell.Y == BoardGrid-4))

	pawn.Position.update(cell)
}

func (pawn *Pawn) isFirstMove() bool {
	return !pawn.Position.hasMoved &&
		((pawn.Color.IsBlack() && pawn.Position.Y == 1) ||
			(pawn.Color.IsWhite() && pawn.Position.Y == BoardGrid-2))
}
