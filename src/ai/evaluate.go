package ai

import "chess"

const (
	DRAW float64 = 0
	WHITEWIN float64		= 255
	BLACKWIN float64		= -255
	
	KINGINCORNER			= .15  // king in a castled position
	KINGONOPENFILE		= -.3  // king not protected by a pawn
	KINGPROTECTED		= .1   // king protected by a pawn, applies to pawns on files near king

	HUNGPIECE			= 0
	ADVANCEDPAWN			= .075 // how far a pawn is from its starting rank
	LONGPAWNCHAIN		= .03  // per pawn
	ISOLATEDPAWN			= -.3
	DOUBLEDPAWN			= -.4  // increases for tripled, etc. pawns
	
	PASSEDPAWN			= .75 // pawn has no opposing pawns blocking it from promoting
	
	CENTRALKNIGHT		= .5   // knight close to center of board
	BISHOPSQUARES		= .125 // per square a bishop attacks
	ROOKONSEVENTH		= .8   // rook is on the second to last rank relative to color
	CONNECTEDROOKS		= .5   // both rooks share the same rank or file
	
	IMPORTANTSQUARE		= .28  // the central squares
	WEAKSQUARE			= .03  // outer squares
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
			Y: p.Position.Y + dir[1]*n,
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
	
	score += pawnStructureAnalysis(whitePawns, 1)
	score -= pawnStructureAnalysis(blackPawns, -1)
	
	whiteRooks := []chess.Square{}
	blackRooks := []chess.Square{}
	
	for _, piece := range b.Board {
		if !piece.Captured {
			if piece.Name != 'q' && piece.Name != 'k' {
				if attackArray[piece.Position.X - 1][piece.Position.Y - 1] * piece.Color < 1 {
					score += float64(piece.Color) * HUNGPIECE
				}
			}
			
			switch piece.Name {
				case 'k':
					if piece.Color == 1 {
						score += checkKingSafety(piece.Position.X, whitePawns)
					} else {
						score -= checkKingSafety(piece.Position.X, blackPawns)
					}
				case 'p':
					// reward passed pawn
					if piece.Color == 1 {
						if pawnIsPassed(piece, blackPawns){
							score += PASSEDPAWN
						} 
					} else {
						if pawnIsPassed(piece, whitePawns){
							score -= PASSEDPAWN
						}
					}
				case 'n':
					if 3 <= piece.Position.X && piece.Position.X <= 6 && 3 <= piece.Position.Y && piece.Position.Y <= 6 {
						score += float64(piece.Color) * CENTRALKNIGHT
					}
				case 'b':
					var numAttacking int
					for _, dir := range piece.Directions {
						numAttacking += AttackRay(piece, b, dir)
					}
					score += float64(piece.Color*numAttacking) * BISHOPSQUARES
				case 'r':
					if (piece.Color == -1 && piece.Position.Y == 2) || (piece.Color == 1 && piece.Position.Y == 7){
						score += float64(piece.Color) * ROOKONSEVENTH
					}
					is piece.Color == 1 {
						whiteRooks = append(whiteRooks, piece.Position)
					} else {
						blackRooks = append(blackRooks, piece.Position)
					}
			}
		}
	}
	score += rookAnalysis(whiteRooks)
	score -= rookAnalysis(blackRooks)
	
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			if attackArray[x][y] > 0 {
				if x >= 2 && x <= 5 && y >= 2 && y <= 5 {
					score += IMPORTANTSQUARE
				} else {
					score += WEAKSQUARE
				}
			} else if attackArray[x][y] < 0 {
				if x >= 2 && x <= 5 && y >= 2 && y <= 5 {
					score -= IMPORTANTSQUARE
				} else {
					score -= WEAKSQUARE
				}
			}
		}
	}
	return score
}

// Reward players for protecting their KING with pawns and being in corner
func checkKingSafety(file int, pawns []chess.Square) float64 {
	pawnArray := [8]int{}
	for _, p := range pawns {
		pawnArray[p.X - 1] += 1
	}
	
	var score float64
	for i := -1; i < 2; i++ {
		if location := file + 1; location > -1 && location < 8 {
			if pawnArray[location] == 0 {
				score += KINGONOPENFILE
			} else {
				score += KINGPROTECTED
			}
		}
	}
	
	if file == 1 || file == 2 || file == 7 || file == 8 {
		score += KINGINCORNER
	} else {
		score -= KINGINCORNER
	}
	
	return score
}

// Used in pawnStructureAnalysis to update a score given a discovered to be broken pawn chain
func updatePawnChainScore(pawnChain int) float64 {
	var score float64
	if pawnChain > 2 {
		score += float64(pawnChain) * LONGPAWNCHAIN
	} else if pawnChain != 0 {
		score += ISOLATEDPAWN / float64(pawnChain)
	}
	
	return score
}

// Returns appropriate penalties for double and isolated  pawns
func pawnStructureAnalysis(pawns []chess.Square, color int) {
	pawnArray := [8]int{}
	var score float64
	
	for _, p := range pawns {
		pawnArray[p.X - 1] += 1
		
		if color == 1 {
			score += float64(p.Y - 2) * ADVANCEDPAWN
		} else {
			score += float64(7 - p.Y) * ADVANCEDPAWN
		}
	}
	
	var pawnChain int
	
	for _, count := range pawnArray {
		if count >= 2 {
			score += float64(count) * DOUBLEDPAWN
			pawnChain += 1
		} else if count == 1 {
			pawnChain += 1
		} else if count == 0 {
			score += updatePawnChainScore(pawnChain)
			pawnChain = 0
		}
	}
	
	score += updatePawnChainScore(pawnChain)
	
	return score
}

func rookAnalysis(rooks []chess.Square) float64 {
	if len(rooks) != 2 {
		return 0
	}
	if rooks[0].X == rooks[1].X || rooks[0].Y == rooks[1].Y {
		return CONNECTEDROOKS
	}
	
	return 0
}

// Return wheter a given pawn has no opposing pawns blocking its path in any of its adjacent files
func pawnIsPassed(pawn *chess.Piece, oppFullPawns []chess.Square) bool {
	for _, p := range oppFullPawns {
		if absInt(p.X - pawn.Position.X) <= 1 {
			if (pawn.Color == 1 && p.Y > pawn.Position.Y) || (pawn.Color == -1 && p.Y < pawn.Position.Y) {
				return false
			}
		}
	}
	
	return true
}









