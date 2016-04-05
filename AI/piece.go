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
	Return all legal moves for a given piece
	check: 
		is true when :
			moves that would place the player in check are not returned
	Ex: if a pinned piece is giving check
*/
func (p *Piece) legalMoves(b *Board, check bool) []*Move {
	legals := make([]*Move, 0)
	
	if p.Captured {
		return legals
	}
	
	if p.Name == 'k' {
		var castley int 
		if b.Turn == 1 {
			castley = 1
		} else if b.Turn == -1 {
			castley = 8
		}
		
		if b.can_castle(1){
			m := p.makeMoveTo(3, castley)
			legals = append(legals, m)
		}
		
		if b.can_castle(8){
			m := p.makeMoveTo(7, castley)
			legals = append(legals, m)
		}
	}
	
	if p.Infinite_direction{
		for _, direction := range p.Directions {
			for i := 1; i < 8; i++ {
				s := Square {
					X: p.Position.X + direction[0]*i,
					Y: p.Position.Y + direction[1]*i,
				}
				
				if o, capname := b.Occupied(s); o == -2 || o == p.Color {
					break
				} else if o == p.Color*-1{
					m := p.makeMoveTo(s.X, s.Y)
					m.Capture = capname
					
					if check {
						if !moveIsCheck(b, m) {
							legals = append(legals, m)
						}
					} else {
						legals = append(legals, m)
					}
					break
				} else {
					m := p.makeMoveTo(s.X, s.Y)
					if check {
						if !moveIsCheck(b, m){
							legals = append(legals, m)
						}
					} else {
						legals = append(legals, m)
					}
				}
			}
		}
	} else {
		for _, direction := range p.Directions {
			s := Square {
				X: p.Position.X + direction[0],
				Y: p.Position.Y + direction[1],
			}
			
			if o, capname := b.Occupied(&s); o == 0 || (o == p.Color*-1 && p.Name != 'p'){
				m := p.makeMoveTo(s.X, s.Y)
				m.Capture = capname
				if p.Name == 'p' && ((p.Color == 1 && s.Y == 8)||(p.Color == -1 && s.Y == 1)){
					for _, promotion := range [4]byte{'q', 'r', 'n', 'b'} {
						move := m.CopyMove()
						move.Promotion = promotion
						if check {
							if !moveIsCheck(b, m){
								legals = append(legals, m)
							}
						} else {
							legals = append(legals, m)
						}
					}
				} else {
					if check {
						if !moveIsCheck(b. m) {
							legals = append(legals, m)
						}
					} else {
						legals = append(legals, m)
					}
				}
			}
		}
	}
	
	if p.Name == 'p' {
		captures := [2][2]int{{1, -1}, {1, 1}}
		for _, val := range captures {
			capture := Square {
				X: p.Position.X + val[1],
				Y: p.Position.Y + val[0]*p.Color,
			}
			
			if o, capname := b.Occupied(&capture); o == p.Color*-1 {
				m := p.makeMoveTo(capture.X, capture.Y)
				m.Capture = capname
				
				if p.Name == 'p' && ((p.Color == 1 && capture.Y == 8)|| (p.Color == -1 && capture.Y == 1)){
					for _, promotion := range [4]byte{'q', 'b', 'n', 'r'} {
						move := m.CopyMove()
						move.Promotion = promotion
						if check {
							if !moveIsCheck(b, m) {
								legals = append(legals, m)
							}
						} else {
							legals = append(legals, m)
						}
					}
				} else {
					if check {
						if !moveIsCheck(b, m) {
							legals = append(legals, m)
						}
					} else {
						legals = append(legals, m)
					}
				}
			}
		}
		
		if (p.Color == 1 && p.Position.Y == 2) || (p.Color == -1 && p.Position.Y == 7) {
			singleSquare := Square{
				X: p.Position.X,
				Y: p.Position.Y + p.Color,
			}
			doubleSquare := Square{
				X: p.Position.X,
				Y: p.Position.Y + 2*p.Color,
			}
			
			if so, _ := b.Occupied(&singleSquare); so == 0 {
				if do, _ := b.Occupied(&doubleSquare); do == 0 {
					m := p.makeMoveTo(doubleSquare.X, doubleSquare.Y)
					if check {
						if !moveIsCheck(b, m) {
							legals = append(legals, m)
						}
					} else {
						legals = append(legals, m)
					}
				} 
			}
		} else {
			en_passants := [2][2]int{{1,0}, {-1, 0}}
			for _, val := range en_passants {
				s := Square{
					X: p.Position + val[0],
					Y: p.Position.Y,
				}
				if o,_ := b.Occupied(&s); o == p.Color*-1 {
					for _, piece := range b.Board {
						if piece.Position == s && piece.Can_en_passant == true {
							captureSquare := Square{
								X: p.Position.X + val[0],
								Y: p.Position.Y + p.Color,
							}
							m := p.makeMoveTo(capturesquare.X, capturesquare.Y)
							m.Capture = 'p'
							if checkcheck {
								if !moveIsCheck(b, m) {
									legals = append(legals, m)
								}
							} else {
								legals = append(legals, m)
							}
						}
					}
				}
			}
		}
	}
	return legals
}

// Given a piece and a destination, constructs a valid move struct.
// Promotion must be added after the fact.
func (p *Piece) makeMoveTo(x, y int) *Move {
	m := &Move {
		Piece: p.Name,
		Begin: p.Position,
		End : Square{
			X: x,
			Y: y,
		},
	}
	
	return m
}















