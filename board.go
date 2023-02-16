package ChessGame

type chessPieceColor bool

const BoardGrid = 8

const (
	White chessPieceColor = true
	Black chessPieceColor = false
)

type Board struct {
	BlackPieces chessPieces
	WhitePieces chessPieces
	Turn        chessPieceColor

	piecesOnCell [BoardGrid][BoardGrid]ChessPiece
}

func (board *Board) AddPiece(piece ChessPiece) {
	position := piece.getPosition()
	if board.piecesOnCell[position.X][position.Y] != nil {
		panic("There is a piece already!")
	}

	if piece.getColor() == White {
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
				Position: Cell{i, 1},
			})

			board.AddPiece(&Pawn{
				Color:    White,
				Position: Cell{i, BoardGrid - 2},
			})
		}

		board.AddPiece(&Rook{
			Color:    Black,
			Position: Cell{0, 0},
		})
		board.AddPiece(&Rook{
			Color:    Black,
			Position: Cell{7, 0},
		})
		board.AddPiece(&Rook{
			Color:    White,
			Position: Cell{0, 7},
		})
		board.AddPiece(&Rook{
			Color:    White,
			Position: Cell{7, 7},
		})

		board.AddPiece(&Knight{
			Color:    Black,
			Position: Cell{1, 0},
		})
		board.AddPiece(&Knight{
			Color:    Black,
			Position: Cell{6, 0},
		})
		board.AddPiece(&Knight{
			Color:    White,
			Position: Cell{1, 7},
		})
		board.AddPiece(&Knight{
			Color:    White,
			Position: Cell{6, 7},
		})

		board.AddPiece(&Bishop{
			Color:    Black,
			Position: Cell{2, 0},
		})
		board.AddPiece(&Bishop{
			Color:    Black,
			Position: Cell{5, 0},
		})
		board.AddPiece(&Bishop{
			Color:    White,
			Position: Cell{2, 7},
		})
		board.AddPiece(&Bishop{
			Color:    White,
			Position: Cell{5, 7},
		})

		board.AddPiece(&Queen{
			Color:    Black,
			Position: Cell{3, 0},
		})
		board.AddPiece(&Queen{
			Color:    White,
			Position: Cell{3, 7},
		})

		board.AddPiece(&King{
			Color:    Black,
			Position: Cell{4, 0},
		})
		board.AddPiece(&Queen{
			Color:    White,
			Position: Cell{4, 7},
		})

	}

	return board
}
