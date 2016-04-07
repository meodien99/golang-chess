package ai

import "chess"

const (
	DRAW float64 = 0
	WHITEWIN float64 = 255
	BLACKWIN float64 = -255
)

var (
	VALUES = map[byte]int{'p':1, 'n':3, 'b':3, 'r': 5, 'q': 9}
)

// default math package uses float64
func absInt(i int) int {
	if i > 0 {
		return i
	}
	
	return i*-1
}


/*
Based heavily off of the analysis function :

http://www.frayn.net/beowulf/theory.html#analysis
*/

// Represents the board as an array of aggression.
// Each value is how many times the mover attacks the square minus how many times the other player defends it.
func updateAttackArray(b *chess.Board, p *chess.Piece, a *[8][8]int){
	if p.Name == 'p' {
		captures := [2][2]int{{1, 1*p.Color}, {-1, 1*p.Color}}
		for _, capture := range captures {
			capx := p.Position.X + capture[0]
			capy := p.Position.Y + capture[1]
			
			if 0 < capx && capx < 9 && 0 < capy && capy < 9 {
				a[capx - 1][capy - 1] += p.Color * b.Turn
			}
		}
	} else {
		for _, dir := range p.Directions {
			if p.Infinite_direction {
				for i:= 1; i <= AttackRay(p, b, dir); i++ {
					a[p.Position.X + dir[0]*i - 1][p.Position.Y + dir[1]*i - 1] += p.Color * b.Turn
				}
			} else {
				destx := p.Position.X + dir[0]
				desty := p.Position.Y + dir[1]
				
				if 0 < destx && destx < 9 && 0 < desty && desty < 9 {
					a[destx-1][desty -1] += p.Color * b.Turn
				}
			}
		}
	}
}

// Measures how many square a piece can attack in a given direction
func AttackRay(b *chess.Board, p *chess.Piece, dir [2]int) int {
	if p.Captured {
		return 0
	}
	
	if !p.Infinite_direction {
		return 1
	}
	
	if n:=1; n < 8; n++ {
		s := &chess.Square{
			X: p.Position.X + dir[0]*n,
			Y: p.Position.Y + dir[0]*n,
		}
		
		if ocuppied, _ := b.Occupied(s); ocuppied != 0 {
			if ocuppied == -2 {
				return n - 1
			}
			return n
		}
	}
	return 7
}

// Return a score from the point of view of the person whose turn it is
// Positive numbers indicate a stronger position
func EvalBoard(b *chess.Board) float64 {
	if over := b.IsOver(); over != 0 {
		if over == 1 {
			return DRAW
		} else {
			if over > 0 {
				return WHITEWIN
			} else {
				return BLACKWIN
			}
		}
	}
	attackArray := [8][8]int{}
	whitePawns := []chess.Square{}
	blackPawns := []chess.Square{}
	
	var score float64
	
	for _, piece := range b.Board {
		if !piece.Captured {
			score += float64(VALUES[piece.Name] * piece.Color)
			updateAttackArray(b, piece, &attackArray)
			
			if piece.Name == 'p' {
				if piece.Color == 1 {
					whitePawns = append(whitePawns, piece.Position)
				} else {
					blackPawns = append(blackPawns, piece.Position)
				}
			}
		}
	}
	
	return score
}















