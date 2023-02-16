package ChessGame

type turn bool

const (
	White turn = true
	Black turn = false
)

type Board struct {
	blackPieces chessPieces
	whitePieces chessPieces
	turn        turn
}

func CreateBoard() *Board {

	return nil
}
