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
	}

	return board
}
