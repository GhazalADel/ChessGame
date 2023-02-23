package ChessGame

type Cell struct {
	X int
	Y int

	hasMoved bool
}

func (cell *Cell) update(new Cell) {
	cell.X = new.X
	cell.Y = new.Y
	cell.hasMoved = true
}

func (cell *Cell) updateStatic(x, y int) {
	cell.X = x
	cell.Y = y
	cell.hasMoved = true
}

func (cell *Cell) isUndefined() bool {
	return cell.X < 0 || cell.Y < 0 ||
		cell.X >= BoardGrid || cell.Y >= BoardGrid
}

func (cell *Cell) Equals(other *Cell) bool {
	return cell.X == other.X && cell.Y == other.Y
}
