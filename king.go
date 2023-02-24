package ChessGame

type King struct {
	Position Cell
	Color    chessPieceColor

	board *Board
}

func (king *King) getType() chessPieceType {
	return typeKing
}

func (king *King) GetPosition() *Cell {
	return &king.Position
}

func (king *King) GetColor() chessPieceColor {
	return king.Color
}

func (king *King) attachToBoard(board *Board) {
	king.board = board
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
		if !Check(king.board, king.Color, &cell) {
			validateAndAddMove(&moves, king, replace, cell)
		}
	}
	if king.Color {
		for _, piece := range king.board.WhitePieces.data[2] {
			if king.canCasting(piece) {
				cell := Cell{X: piece.GetPosition().X, Y: piece.GetPosition().Y}
				replace := king.board.GetPieceOnCell(cell)
				validateAndAddMove(&moves, king, replace, cell)
			}
		}
	} else {
		for _, piece := range king.board.BlackPieces.data[2] {
			if king.canCasting(piece) {
				cell := Cell{X: piece.GetPosition().X, Y: piece.GetPosition().Y}
				replace := king.board.GetPieceOnCell(cell)
				validateAndAddMove(&moves, king, replace, cell)
			}
		}
	}

	return moves
}

func (king *King) moveTo(cell Cell) {
	king.Position.update(cell)
}
func (king *King) canCasting(rook ChessPiece) bool {
	isEmpty := true
	kingCell := king.Position
	if !rook.GetPosition().hasMoved && !king.Position.hasMoved && !Check(king.board, king.Color, &kingCell) {
		currentY := king.Position.Y
		start := 0
		end := 0
		if king.Position.X < rook.GetPosition().X {
			start = king.Position.X + 1
			end = rook.GetPosition().X
		} else {
			start = rook.GetPosition().X + 1
			end = king.Position.X
		}
		for i := start; i < end; i++ {
			cell := Cell{
				X: i,
				Y: currentY,
			}
			replace := king.board.GetPieceOnCell(cell)
			if replace != nil {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			rookCell := rook.GetPosition()
			if !Check(king.board, king.Color, rookCell) {
				for j := start; j < end; j++ {
					c := Cell{
						X: j,
						Y: currentY,
					}
					if Check(king.board, king.Color, &c) {
						return false
					}
				}
				return true
			}
		}
	}
	return false
}
