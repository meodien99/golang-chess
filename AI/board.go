/*
	Package ai implement the rules for playing chess
*/

package ai

import (
	"fmt"
	"strconv"
	"strings"
	"errors"
)

type Board struct {
	Board []*Piece // all of piece on board
	Turn int // 1: white, -1: black
}


/*
	Return the color of the piece that occupies a given square
	If the square is empty, returns 0
	If the square is outside of the bounds of the board, returns -2
*/
func (b *Board) Occupied(s *Square) (int, byte) {
	var capture byte
	if !(1 <= s.X && s.X <= 8 && 1 <= s.Y && s.Y <= 8){
		return -2, capture
	}
	
	for _, p := range b.Board {
		if p.Position.X == s.X && p.Position.Y == s.Y && !p.Captured {
			return p.Color, p.Name
		}
	}
	
	return 0, capture
}


//Convert board to array of string, ready for printing or conversion to FEN
func (b *Board) ToArray() (boardArray [8][8]string) {
	boardArray = [8][8]string{}
	
	for _, piece := range boardArray {
		if !piece.Captured {
			if piece.Color == 1 {
				boardArray[piece.Position.Y-1][piece.Position.X-1] = strings.ToUpper(string(piece.Name))
			} else {
				boardArray[piece.Position.Y-1][piece.Position.X-1] = string(piece.Name)
			}
		}
	}
	return
}


func (b *Board) PrintBoard(){
	boardArray := b.ToArray()
	
	for y:= 7; y >=0; y-- {
		for x:=0; x < 8; x ++ {
			if boardArray[y][x] == "" {
				fmt.Printf(". ")
			} else {
				fmt.Printf("%s", boardArray[y][x])
			}
		}
	}
	fmt.Println()
}

// Converts the position to FEN, only including position and turn.
// See: http://en.wikipedia.org/wiki/Forsyth%E2%80%93Edwards_Notation
func (b *Board) ToFen() (fen string) {
	boardArray := b.ToArray()
	fen = ""
	empty := 0
	
	for y:=7; y>=0; y-- {
		for x:=0; x < 8; x++ {
			if boardArray[y][x] == "" {
				empty += 1
			} else {
				if empty != 0 {
					fen += strconv.Itoa(empty)
					empty = 0
				}
				
				fen += boardArray[y][x]
			}
		}
		
		if empty != 0 {
			fen += strconv.Itoa(empty)
			empty = 0
		}
		
		if y != 0 {
			fen += "/"
		}
	}
	
	if b.Turn == 1 {
		fen += " w"
	} else {
		fen += " b"
	}
	
	return
}

/*
	Check if king is in check
	Pass the color of the king that you want to check
	Return true if king in check / false if not
*/
func (b *Board) IsCheck(color int) bool {
	var kingsquare Square
	
	if color = 1 {
		kingsquare = b.Board[0].Position
	} else {
		kingsquare = b.Board[1].Position
	}
	
	for _, piece := range b.Board {
		if piece.Color == color*-1 {
			for _, move := range piece.legalMoves(b, false) {
				if move.End == kingsquare {
					return true
				}
			}
		}
	}
	return false
}

// Determines if castling to a particular side is legal.
// Side should be 1 or 8 depending to indicate kingside or queenside castling.
func (b *Board) can_castle(side int) bool {
	var rookIndex int
	var kingIndex int
	
	if b.Turn == 1 {
		kingIndex = 0
	} else {
		kingIndex = 1
	}
	
	if !b.Board[kingIndex].Can_castle {
		return false
	}
	
	for i, p := range b.Board {
		if p.Name == 'r' && side == p.Position.X {
			rookIndex = i
			break
		}
	}
	
	if rookIndex == 0 {
		return false
	}
	
	if !b.Board[rookIndex].Can_castle {
		return false
	}
	
	if b.Board[rookIndex].Position.Y != b.Board[kingIndex].Position.Y {
		return false
	}
	
	for i := minInt(b.Board[rookindex].Position.X, b.Board[kingindex].Position.X) + 1; i < maxInt(b.Board[rookindex].Position.X, b.Board[kingindex].Position.X); i++ {
		s := &Square{
			X: i,
			Y: b.Board[kingIndex].Position.Y,
		}
		
		if o,_ := b.Occupied(s); o != 0 {
			return false
		}
		
		if i != 2{
			for _, p := range b.Board {
				if p.Color == b.Turn*-1 && p.Attacking(s, b){
					return false
				}
			}
		}
	}
		
	return true
}

func (b *Board) castleHandler(m *Move, side int) error {
	var rookIndex int
	var kingIndex int
	
	if b.Turn == 1 {
		kingIndex = 0
	} else {
		kingIndex = 1
	}
	
	for i, p := range b.Board {
		if p.Name == 'r' && side == p.Position.X {
			rookIndex = i
			break
		}
	}
	
	if rookIndex == 0 {
		return errors.New("func castleHandler: should have found rook")
	}
	
	b.Board[kingIndex].Position = m.End
	if side == 8 {
		b.Board[rookIndex].Position.X = 6
	}
	
	if side == 1 {
		b.Board[rookIndex].Position.X = 4
	}
	
	return nil
}


// Modifies a board in-place to undo a given move
func (b *Board) UndoMove(m *Move) {
	var pieceAdded bool
	var pieceMoved bool
	
	for i,p := range b.Board {
		if p.Position == m.End {
			if p.Color == b.Turn*-1{
				if !pieceMoved && !p.Captured {
					b.Board[i].Position = m.Begin
					pieceMoved = true
					
					if m.Piece == 'p' && b.Board[i].Name != 'p' {
						//undo pawn promotion
						b.Board[i].Name = 'p'
						b.Board[i].Infinite_direction = false
						b.Board[i].Directions = [][2]int{
							{0, 1 * p.Color},
						}
					}
				} else if m.Piece == 'k' {
					// undo castle
					if m.Begin.X == 5 && (m.End.X == 3 || m.End.X == 7) && ((p.Color == 1 && m.Begin.Y == 1) || (p.Color == -1 && m.Begin.Y == 8)){
						
						for i, p := range b.Board {
							if p.Name == 'r' && p.Color == b.Turn*-1 && !p.Captured && p.Position.Y == m.Begin.Y {
								if m.End.X == 3 && p.Position.X == 4 {
									b.Board[i].Position.X = 1
									break
								} else if m.End.X == 7 && p.Position.X == 6 {
									b.Board[i].Position.X = 8
									break
								}
							}
						}
					}
				}
			} else {
				if p.Captured && p.Name == m.Capture && !pieceAdded {
					b.Board[i].Captured = false
					pieceAdded = true
				}
			}
		}
	}
	
	b.Turn*= -1
}


// Modifies a bord in-place.
// Forces a piece to a given square without checking move legality.
func (b *Board) ForceMove(m *Move) {
	for i, p := range b.Board {
		if !p.Captured {
			if m.Begin == p.Position {
				b.Board[i].Position.X, b.Board[i].Position.Y = m.End.X, m.End.Y
				
				if m.Piece == 'p' {
					if (p.Color == 1 && m.End.Y == 8 ) || (p.Color == -1 && m.End.Y == 1) {
						if promotion := m.Promotion; promotion == 'q' {
							b.Board[i].Name = promotion
							b.Board[i].Directions = [][2]int{
								{1, 1},
								{1, 0},
								{1, -1},
								{0, 1},
								{0, -1},
								{-1, 1},
								{-1, 0},
								{-1, -1},
							}
							b.Board[i].Infinite_direction = true
						} else if promotion == 'r' {
							b.Board[i].Name = promotion
							b.Board[i].Directions = [][2]int {
								{1, 0},
								{-1, 0},
								{0, 1},
								{0, -1},
							}
							b.Board[i].Infinite_direction = true
						} else if promotion == 'n' {
							b.Board[i].Name = promotion
							b.Board[i].Directions = [][2]int {
								{1, 2},
								{-1, 2},
								{1, -2},
								{-1, -2},
								{2, 1},
								{-2, 1},
								{2, -1},
								{-2, -1},
							}
						} else if promotion == 'b' {
							b.Board[i].Name = promotion
							b.Board[i].Directions = [][2]int {
								{1, 1},
								{1, -1},
								{-1, 1},
								{-1, -1},
							}
							b.Board[i].Infinite_direction = true
						}
					}
				} else if m.Piece == 'k' {
					if m.Begin.X == 5 && (m.End.X == 3 || m.End.X == 7) && ((p.Color == 1 && m.Begin.Y == 1) || (p.Color == -1 && m.Begin.Y == 8)) {
						// if king trying to castle
						for i, p := range b.Board {
							if p.Name == 'r' && p.Color == b.Turn && !p.Captured && p.Position.Y == m.Begin.Y {
								if m.End.X == 3 && p.Position.X == 1 {
									b.Board[i].Position.X = 4
									break
								} else if m.End.X == 7 && p.Position.X == 8 {
									b.Board[i].Position.X = 6
									break
								}
							}
						}
					}
				}
			} else if p.Position.X == m.End.X && p.Position.Y == m.End.Y {
				b.Board[i].Captured = true
			}
		}
	}
	b.Turn *= -1
}










