/*
	Package ai implement the rules for playing chess
*/

package ai

import (
	"fmt"
	"strconv"
	"strings"
)

type Board struct {
	Board []*Piece // all of piece on board
	Turn int // 1: white, -1: black
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














