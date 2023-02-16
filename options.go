package ChessGame

type chessGameOptions struct {
	InitializePieces bool
	StartTurnWith    chessPieceColor
}

var Options = chessGameOptions{
	InitializePieces: true,
	StartTurnWith:    White,
}
