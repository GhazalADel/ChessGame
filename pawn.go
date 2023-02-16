package ChessGame

type Pawn struct {
	Position Cell
	Color    chessPieceColor

	board *Board
}

func (pawn *Pawn) getType() chessPieceType {
	return typePawn
}

func (pawn *Pawn) getPosition() *Cell {
	return &pawn.Position
}

func (pawn *Pawn) getColor() chessPieceColor {
	return pawn.Color
}

func (pawn *Pawn) attachToBoard(board *Board) {
	pawn.board = board
}
