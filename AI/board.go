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
	
}

// Converts the position to FEN, only including position and turn.
// See: http://en.wikipedia.org/wiki/Forsyth%E2%80%93Edwards_Notation
func (b *Board) ToFen() string {
	return
}