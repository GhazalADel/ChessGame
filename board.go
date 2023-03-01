package ChessGame

type chessPieceColor bool

const BoardGrid = 8

const (
	White chessPieceColor = true
	Black chessPieceColor = false
)

func (color chessPieceColor) IsWhite() bool {
	return color == White
}

func (color chessPieceColor) IsBlack() bool {
	return color == Black
}

type Board struct {
	BlackPieces chessPieces
	WhitePieces chessPieces
	Turn        chessPieceColor

	piecesOnCell [BoardGrid][BoardGrid]ChessPiece
	isSimulating bool
}

func (board *Board) AddPiece(piece ChessPiece) {
	position := piece.GetPosition()
	if board.piecesOnCell[position.X][position.Y] != nil {
		panic("There is a piece already!")
	}

	if piece.GetColor() == White {
		board.WhitePieces.add(piece)
	} else {
		board.BlackPieces.add(piece)
	}

	piece.attachToBoard(board)
	board.piecesOnCell[position.X][position.Y] = piece
}

func CreateBoard() *Board {
	board := new(Board)
	board.Turn = Options.StartTurnWith

	if Options.InitializePieces {
		// add default chess pieces here

		for i := 0; i < BoardGrid; i++ {
			board.AddPiece(&Pawn{
				Color:    Black,
				Position: Cell{X: i, Y: 1},
			})

			board.AddPiece(&Pawn{
				Color:    White,
				Position: Cell{X: i, Y: BoardGrid - 2},
			})
		}

		board.AddPiece(&Rook{
			Color:    Black,
			Position: Cell{X: 0, Y: 0},
		})
		board.AddPiece(&Rook{
			Color:    Black,
			Position: Cell{X: 7, Y: 0},
		})
		board.AddPiece(&Rook{
			Color:    White,
			Position: Cell{X: 0, Y: 7},
		})
		board.AddPiece(&Rook{
			Color:    White,
			Position: Cell{X: 7, Y: 7},
		})

		board.AddPiece(&Knight{
			Color:    Black,
			Position: Cell{X: 1, Y: 0},
		})
		board.AddPiece(&Knight{
			Color:    Black,
			Position: Cell{X: 6, Y: 0},
		})
		board.AddPiece(&Knight{
			Color:    White,
			Position: Cell{X: 1, Y: 7},
		})
		board.AddPiece(&Knight{
			Color:    White,
			Position: Cell{X: 6, Y: 7},
		})

		board.AddPiece(&Bishop{
			Color:    Black,
			Position: Cell{X: 2, Y: 0},
		})
		board.AddPiece(&Bishop{
			Color:    Black,
			Position: Cell{X: 5, Y: 0},
		})
		board.AddPiece(&Bishop{
			Color:    White,
			Position: Cell{X: 2, Y: 7},
		})
		board.AddPiece(&Bishop{
			Color:    White,
			Position: Cell{X: 5, Y: 7},
		})

		board.AddPiece(&Queen{
			Color:    Black,
			Position: Cell{X: 3, Y: 0},
		})
		board.AddPiece(&Queen{
			Color:    White,
			Position: Cell{X: 3, Y: 7},
		})

		board.AddPiece(&King{
			Color:    Black,
			Position: Cell{X: 4, Y: 0},
		})
		board.AddPiece(&King{
			Color:    White,
			Position: Cell{X: 4, Y: 7},
		})
	}
	return board
}

func (board *Board) GetPieceOnCell(c Cell) ChessPiece {
	if c.isUndefined() {
		return nil
	}
	return board.piecesOnCell[c.X][c.Y]
}

type BoardInterface interface {
	ShowBoard()
	SelectCell(cell Cell)
	getMyBoard() *Board
}
