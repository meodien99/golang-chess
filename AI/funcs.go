package ai

/*
	Return true if a move places the mover in check
*/
func moveIsCheck(b *Board, m *Move) bool {
	var pieceIndex int
	var capture bool
	var capturePieceIndex int
	
	for i,p := range b.Board {
		if p.Position == m.Begin && p.Name == m.Piece && p.Color == b.Turn && !p.Captured {
			pieceIndex = i
		} else if p.Position == m.End && !p.Captured {
			capture = true
			capturePieceIndex = i
		}
	}
	
	b.Board[pieceIndex].Position = m.End
	if capture {
		b.Board[capturePieceIndex].Captured = true
	}
	
	passed := true
	
	if !b.IsCheck(b.Turn){
		passed = false
	}
	
	b.Board[pieceIndex].Position = m.Begin
	if capture {
		b.Board[capturePieceIndex].Captured = false
	}
	
	return passed
}




func maxInt(x, y int) int {
	if x > y {
		return x
	}
	
	return y
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	
	return y
}