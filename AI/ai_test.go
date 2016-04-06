package ai

import "testing"

func TestToFen(t *testing.T){
	board := &Board{Turn: 1}
	board.SetUpPieces()
	
	if fen := board.ToFen(); fen != "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w" {
		t.Errorf("Initial position expected fen:\n%s\nInstead got:\n%s\n", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w", fen)
	}
	
	m := &Move{
		Piece:'p',
		Begin: Square{
			X: 5,
			Y: 2,
		},
		End: Square{
			X: 5,
			Y: 4,
		},
	}
	
	board.ForceMove(m)
	if fen := board.ToFen(); fen != "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b" {
		t.Errorf("After 1...e4 expected fen:\n%s\nInstead got:\n%s\n", "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b", fen)
	}

}

func TestCanCastle(t *testing.T){
	board := &Board{Turn:1}
	
	board.PlacePiece('k', 1, 5, 1)
	board.Board[0].Can_castle = true
	
	board.PlacePiece('r', 1, 8, 1)
	board.Board[1].Can_castle = true
	
	board.PlacePiece('b', 1, 6, 1)
	
	if board.can_castle(8) {
		t.Error("Castle allowed through blocking piece")
	}
	
	board.Board[2].Color = -1
	board.Board[2].Position.Y = 2
	if board.can_castle(8) {
		t.Error("Castle allowed when king in check")
	}
	
	board.Board[2].Position.X = 5
	board.Board[2].Position.Y = 3
	if board.can_castle(8) {
		t.Error("Castle allowed when king placed in check")
	}
	
	board.Board[2].Color = 1
	board.Board[0].Can_castle = false
	if board.can_castle(8) {
		t.Error("Castle allowed after king moved")
	}
	
	board.Board[0].Can_castle = true
	board.Board[1].Can_castle = false
	if board.can_castle(8) {
		t.Error("Castle allowed after rook move")
	}
	
	board.Board[1].Can_castle = true
	board.Board[1].Position.Y = 2
	if board.can_castle(8) {
		t.Error("Castle allowed when rook out of position")
	}
	
	board.Board[1].Position.Y = 1
	if !board.can_castle(8) {
		t.Error("Error when making a legal castle")
	}
}

func TestIsCheck(t *testing.T) {
	board := &Board{Turn:1}
	board.PlacePiece('k', 1, 1, 1)
	board.PlacePiece('k', -1, 8, 8)
	board.PlacePiece('r', 1, 8, 1)
	if check := board.IsCheck(1); check == true {
		t.Error("False positive when dertermining check")
	}
	
	if check := board.IsCheck(-1); check == false {
		t.Error("False negative when dertermining check")
	}
	
	if king := board.Board[0]; king.Position.X != 1 || king.Position.Y != 1 {
		t.Errorf("isCheck modified board, king moves from {X: 1, Y:1} to %+v", king.Position)
	}
}

func TestMoveIsCheck(t *testing.T){
	board := &Board{Turn: 1}
	
}








