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
