# ChessGame

A chess engine and CLI written in Go. The project is split into a core library package (`Chess`) and a command-line interface package (`Chess/cli`), keeping game logic cleanly separated from presentation.

## Features

- Full chess rules for all six piece types: Pawn, Rook, Knight, Bishop, Queen, and King
- En passant captures
- Castling (king-side and queen-side), with proper check and move-history validation
- Check and checkmate detection via move simulation
- Color-coded terminal board rendered with Unicode chess symbols and ANSI escape codes
- Visual highlighting for the selected piece (blue), valid moves (green), and capturable enemies (red)
- Turn-based two-player gameplay in a single terminal session
- Configurable starting options (whose turn goes first, whether to auto-populate pieces)

## Project Structure

```
ChessGame-master/
├── board.go          # Board struct, piece placement, CreateBoard()
├── cell.go           # Cell struct and coordinate helpers
├── chessPiece.go     # ChessPiece interface, move validation, check/checkmate logic
├── pieces.go         # chessPieces collection type
├── options.go        # Global game options (Options var)
├── pawn.go           # Pawn movement, double-advance, en passant
├── rook.go           # Rook movement
├── knight.go         # Knight movement
├── bishop.go         # Bishop movement
├── queen.go          # Queen movement
├── king.go           # King movement and castling
├── cli/
│   └── game.go       # CLI renderer and BoardInterface implementation
├── test/
│   └── main.go       # Entry point / game loop
└── go.mod
```

## Requirements

- Go 1.19 or later
- A terminal with Unicode and ANSI color support (most modern terminals qualify)

## Getting Started

```bash
git clone <repo-url>
cd ChessGame-master
go run ./test/main.go
```

## How to Play

Coordinates use the format `[column][row]`, where columns are letters A–H and rows are numbers 1–8.

1. On your turn, enter the coordinate of the piece you want to move (e.g. `A2`).
2. The board redraws with valid destination squares highlighted.
3. Enter the destination coordinate (e.g. `A4`), or type `back` to pick a different piece.
4. Play alternates between White and Black.

## Architecture

The core library exposes a `BoardInterface` so that alternative frontends (GUI, network, etc.) can be plugged in without touching game logic:

```go
type BoardInterface interface {
    ShowBoard()
    SelectCell(cell Cell)
    GetBoard() *Board
    SetMoves(cells []Cell)
    SetSelectedCell(cell Cell)
}
```

Move legality is enforced by simulating each candidate move on a temporary board copy and checking whether it leaves the moving side's king in check. This same simulation path powers `isCheck`, `isMate`, and castling validation.

Starting configuration can be adjusted via the global `Options` variable before calling `CreateBoard()`:

```go
ChessGame.Options.InitializePieces = true   // populate the standard starting position
ChessGame.Options.StartTurnWith    = ChessGame.White
```

## Known Limitations / TODOs

- End-game detection (stalemate, insufficient material) is not yet implemented — the game loop contains a `// TODO : check end game` comment.
- Pawn promotion is not handled; pawns that reach the back rank remain pawns.
- There is no AI opponent; the game requires two human players sharing a terminal.
