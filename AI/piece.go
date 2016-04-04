/*
	Piece
*/

package ai

type Piece struct {
	Position   Square
	Color      int    // 1: white | -1: black
	Name       string // [p]awn, k[n]ignt, [b]ishop, [r]ook, [q]ueen, [k]ing
	Can_castle bool   // rooks and kings, default true, set to false when piece makes a non-castle move

	Can_en_passant  bool // only applicable to pawns

	Directions         [][2]int // slice of {0 or 1, 0 or 1} indicating how piece moves
	Infinite_direction bool     // if piece can move as far as it wants in given direction
	
	Captured bool
}

/*
	Return true if a piece p is attacking a square s
	("Attacking" means it could capture an opposing piece on that square)
	Ex: A rook is attacking its own pawn next to it, but a pawn is not attacking a piece directly in front end of it
*/
func (p *Piece) Attacking(s *Square, b *Board) bool {
	return
}

/*
	Return true if a move places the mover in check
*/
func moveIsCheck(b *Board, m *Move) bool {
	return
}

/*
	Return all legal moves for a given piece
	check: 
		is true when :
			moves that would place the player in check are not returned
	Ex: if a pinned piece is giving check
*/
func (p *Piece) legalMoves(b *Board, check bool) []*Move {
	return
}