package rules

//piece name + beginning end ending square
type Move struct {
	Piece byte // Piece.Name
	Begin, End Square
	Score float64
	Promotion byte
	Capture byte
}

//Return a pointer to a copy of a move
//Does not copy move's score
func (m *Move) CopyMove() *Move {
	newMove := &Move{
		Piece: m.Piece,
	}
	
	newMove.Begin.X, newMove.Begin.Y = m.Begin.X, m.Begin.Y
	newMove.End.X, newMove.End.Y = m.End.X, m.End.Y
	
	return newMove
}

// Translate move to form "nb1-c3"
func (m *Move) ToString() string {
	return string(m.Piece) + m.Begin.ToString() + "-" + m.End.ToString()
}