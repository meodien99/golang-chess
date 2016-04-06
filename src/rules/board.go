package rules

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Board struct {
	Board []*Piece // all of piece on board
	Turn  int      // 1: white, -1: black
}

/*
	Return the color of the piece that occupies a given square
	If the square is empty, returns 0
	If the square is outside of the bounds of the board, returns -2
*/
func (b *Board) Occupied(s *Square) (int, byte) {
	var capture byte
	if !(1 <= s.X && s.X <= 8 && 1 <= s.Y && s.Y <= 8) {
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
func (b *Board) ToArray() [8][8]string {
	boardArray := [8][8]string{}

	for _, piece := range b.Board {
		if !piece.Captured {
			if piece.Color == 1 {
				boardArray[piece.Position.Y-1][piece.Position.X-1] = strings.ToUpper(string(piece.Name))
			} else {
				boardArray[piece.Position.Y-1][piece.Position.X-1] = string(piece.Name)
			}
		}
	}
	return boardArray
}

func (b *Board) PrintBoard() {
	boardArray := b.ToArray()

	for y := 7; y >= 0; y-- {
		for x := 0; x < 8; x++ {
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

	for y := 7; y >= 0; y-- {
		for x := 0; x < 8; x++ {
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
	var kingSquare Square

	if color == 1 {
		kingSquare = b.Board[0].Position
	} else {
		kingSquare = b.Board[1].Position
	}

	for _, piece := range b.Board {
		if piece.Color == color*-1 {
			for _, move := range piece.legalMoves(b, false) {
				if move.End == kingSquare {
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

	for i := minInt(b.Board[rookIndex].Position.X, b.Board[kingIndex].Position.X) + 1; i < maxInt(b.Board[rookIndex].Position.X, b.Board[kingIndex].Position.X); i++ {
		s := &Square{
			X: i,
			Y: b.Board[kingIndex].Position.Y,
		}

		if o, _ := b.Occupied(s); o != 0 {
			return false
		}

		if i != 2 {
			for _, p := range b.Board {
				if p.Color == b.Turn*-1 && p.Attacking(s, b) {
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

	for i, p := range b.Board {
		if p.Position == m.End {
			if p.Color == b.Turn*-1 {
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
					} else if m.Piece == 'k' {
						// undo castle
						if m.Begin.X == 5 && (m.End.X == 3 || m.End.X == 7) && ((p.Color == 1 && m.Begin.Y == 1) || (p.Color == -1 && m.Begin.Y == 8)) {

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
				}
			} else {
				if p.Captured && p.Name == m.Capture && !pieceAdded {
					b.Board[i].Captured = false
					pieceAdded = true
				}
			}
		}
	}

	b.Turn *= -1
}

// Modifies a bord in-place.
// Forces a piece to a given square without checking move legality.
func (b *Board) ForceMove(m *Move) {
	for i, p := range b.Board {
		if !p.Captured {
			if m.Begin == p.Position {
				b.Board[i].Position.X, b.Board[i].Position.Y = m.End.X, m.End.Y

				if m.Piece == 'p' {
					if (p.Color == 1 && m.End.Y == 8) || (p.Color == -1 && m.End.Y == 1) {
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
							b.Board[i].Directions = [][2]int{
								{1, 0},
								{-1, 0},
								{0, 1},
								{0, -1},
							}
							b.Board[i].Infinite_direction = true
						} else if promotion == 'n' {
							b.Board[i].Name = promotion
							b.Board[i].Directions = [][2]int{
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
							b.Board[i].Directions = [][2]int{
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

// Modifies a board in-place
// Return an error without modifying board if illegal move
// Set a captured piece's location to (0,0)
// Change the turn of the board once move is sucessfully completed
func (b *Board) Move(m *Move) error {
	if m.Piece == 'k' && m.Begin.X-m.End.X != 1 && m.End.X-m.Begin.X != 1 {
		if (b.Turn == 1 && m.End.Y != 1) || (b.Turn == -1 && m.End.Y != 8) {
			return errors.New("func Move: illegal move")
		}

		var side int
		if m.End.X == 7 {
			side = 8
		} else if m.End.X == 3 {
			side = 1
		} else {
			return errors.New("func Move: invalid castle destination")
		}

		if !b.can_castle(side) {
			return errors.New("func can_castle: cannot castle")
		}

		err := b.castleHandler(m, side)

		if err == nil {
			b.Turn *= -1
		}

		return err
	}

	var pieceFound bool
	var pieceIndex int
	var capture bool
	var capturedPiece int

	for i, p := range b.Board {
		if m.Begin == p.Position && m.Piece == p.Name && b.Turn == p.Color && !p.Captured {
			pieceIndex = i
			pieceFound = true
		} else if m.End == p.Position && p.Color == b.Turn*-1 && !p.Captured {
			capture = true
			capturedPiece = i
		}

		if pieceFound && capture {
			break
		}
	}

	if !pieceFound {
		return errors.New("func Move: invalid piece")
	}

	var legal bool
	legals := b.Board[pieceIndex].legalMoves(b, true)

	for _, move := range legals {
		if m.Begin == move.Begin && m.End == move.End && m.Piece == move.Piece {
			legal = true
			b.Board[pieceIndex].Position = move.End
			break
		}
	}

	if !legal {
		return errors.New("func Move: illegal move")
	}

	//en passant
	if !capture && m.Piece == 'p' && (m.Begin.X-m.End.X == 1 || m.End.X-m.Begin.X == 1) {
		capture = true

		for i, p := range b.Board {
			if p.Position.X == m.End.X && p.Position.Y == m.Begin.Y {
				capturedPiece = i
				break
			}
		}
	}

	if capture {
		b.Board[capturedPiece].Captured = true
	}

	if m.Piece == 'k' || m.Piece == 'r' {
		b.Board[pieceIndex].Can_castle = false
	}

	for i, _ := range b.Board {
		b.Board[i].Can_en_passant = false
	}

	if m.Piece == 'p' {
		if m.Begin.Y-m.End.Y == 2*-b.Board[pieceIndex].Color {
			b.Board[pieceIndex].Can_en_passant = true
		} else if (b.Turn == 1 && m.End.Y == 8) || (b.Turn == -1 && m.End.Y == 1) {
			if promotion := m.Promotion; promotion == 'q' {
				b.Board[pieceIndex].Name = promotion
				b.Board[pieceIndex].Directions = [][2]int{
					{1, 1},
					{1, 0},
					{1, -1},
					{0, 1},
					{0, -1},
					{-1, 1},
					{-1, 0},
					{-1, -1},
				}
				b.Board[pieceIndex].Infinite_direction = true
			} else if promotion == 'r' {
				b.Board[pieceIndex].Name = promotion
				b.Board[pieceIndex].Directions = [][2]int{
					{1, 0},
					{-1, 0},
					{0, 1},
					{0, -1},
				}
				b.Board[pieceIndex].Infinite_direction = true
			} else if promotion == 'n' {
				b.Board[pieceIndex].Name = promotion
				b.Board[pieceIndex].Directions = [][2]int{
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
				b.Board[pieceIndex].Name = promotion
				b.Board[pieceIndex].Directions = [][2]int{
					{1, 1},
					{1, -1},
					{-1, 1},
					{-1, -1},
				}
				b.Board[pieceIndex].Infinite_direction = true
			}
		}
	}

	b.Turn *= -1

	return nil
}

// Return all legal moves available to the player whose turn it is.
func (b *Board) AllLegalMoves() []*Move {
	legals := make([]*Move, 0)

	for _, p := range b.Board {
		if p.Color == b.Turn {
			for _, m := range p.legalMoves(b, true) {
				legals = append(legals, m)
			}
		}
	}

	return legals
}

// Check if the game has ended
// Return 2 if white wins, -2 if black wins, 1 if it's stalemate, 0 if it still going
func (b *Board) IsOver() int {
	if len(b.AllLegalMoves()) == 0 {
		if b.IsCheck(b.Turn) {
			return -2 * b.Turn
		}
		return 1
	}
	return 0
}

// Give a name, color, and coordinates, place the appropriate piece on the board
// Does not add flags such as Can_castle, must be done manually
func (b *Board) PlacePiece(name byte, color, x, y int) {
	p := &Piece{
		Name:  name,
		Color: color,
		Position: Square{
			X: x,
			Y: y,
		},
	}

	if name == 'b' || name == 'r' || name == 'q' {
		p.Infinite_direction = true
	}

	if name == 'p' {
		p.Directions = [][2]int{
			{0, 1 * color},
		}
	} else if name == 'b' {
		p.Directions = [][2]int{
			{1, 1},
			{1, -1},
			{-1, 1},
			{-1, -1},
		}
	} else if name == 'n' {
		p.Directions = [][2]int{
			{1, 2},
			{-1, 2},
			{1, -2},
			{-1, -2},
			{2, 1},
			{-2, 1},
			{2, -1},
			{-2, -1},
		}
	} else if name == 'r' {
		p.Directions = [][2]int{
			{1, 0},
			{-1, 0},
			{0, 1},
			{0, -1},
		}
	} else if name == 'q' {
		p.Directions = [][2]int{
			{1, 1},
			{1, 0},
			{1, -1},
			{0, 1},
			{0, -1},
			{-1, 1},
			{-1, 0},
			{-1, -1},
		}
	} else if name == 'k' {
		p.Directions = [][2]int{
			{1, 1},
			{1, 0},
			{1, -1},
			{0, 1},
			{0, -1},
			{-1, 1},
			{-1, 0},
			{-1, -1},
		}
	}

	b.Board = append(b.Board, p)
}

func (b *Board) SetUpPieces() {
	b.Board = make([]*Piece, 0)

	pawnRows := [2]int{2, 7}
	pieceRows := [2]int{1, 8}
	rookFiles := [2]int{1, 8}
	knightFiles := [2]int{2, 7}
	bishopFiles := [2]int{3, 6}
	queenFile := 4
	kingFile := 5

	for _, rank := range pieceRows {
		//put the king first
		var color int
		if rank == 1 {
			color = 1
		} else {
			color = -1
		}

		b.PlacePiece('k', color, kingFile, rank)
		b.Board[len(b.Board)-1].Can_castle = true
	}

	for _, rank := range pieceRows {
		var color int
		if rank == 1 {
			color = 1
		} else {
			color = -1
		}

		for _, file := range rookFiles {
			b.PlacePiece('r', color, file, rank)
			b.Board[len(b.Board)-1].Can_castle = true
		}

		for _, file := range knightFiles {
			b.PlacePiece('n', color, file, rank)
		}

		for _, file := range bishopFiles {
			b.PlacePiece('b', color, file, rank)
		}

		//queen at last
		b.PlacePiece('q', color, queenFile, rank)
	}

	for _, rank := range pawnRows {
		var color int
		if rank == 2 {
			color = 1
		} else {
			color = -1
		}

		for file := 1; file <= 8; file++ {
			b.PlacePiece('p', color, file, rank)
		}
	}
}
