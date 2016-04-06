package rules

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
	board.PlacePiece('k', 1, 1, 1)
	board.PlacePiece('b', 1, 2, 2)
	board.PlacePiece('q', -1, 4, 4)
	checkmove := board.Board[1].makeMoveTo(3,1)

	if check := moveIsCheck(board, checkmove); !check {
		t.Error("Check not recognized")
	}
	
	okmove := board.Board[1].makeMoveTo(3, 3)
	if check := moveIsCheck(board, okmove); check {
		t.Error("False positive with ok move")
	}
	
	captureMove := board.Board[1].makeMoveTo(4, 4)
	if check := moveIsCheck(board, captureMove); check {
		t.Error("Capturing pinning piece with pinned piece places user in check")
	}
	
	board = &Board{Turn: 1}
	board.PlacePiece('k', 1, 1, 1)
	board.PlacePiece('r', -1, 8, 1)
	board.PlacePiece('b', 1, 7, 2)
	m := board.Board[2].makeMoveTo(8, 1)
	
	if check := moveIsCheck(board, m); check {
		t.Error("Capturing the attacking piece still places user in check")
	}
}

func TestOccupied(t *testing.T) {
	b := &Board{}
	b.SetUpPieces()
	whiteSquare := &Square{
		X: 1,
		Y: 1,
	}
	blackSquare := &Square{
		X: 8,
		Y: 8,
	}
	emptySquare := &Square{
		X: 5,
		Y: 5,
	}
	nonSquare := &Square{
		X: 10,
		Y: 10,
	}
	
	if out, _ := b.Occupied(whiteSquare); out != 1 {
		t.Errorf("Expected 1, got %d", out)
	}
	if out, _ := b.Occupied(blackSquare); out != -1 {
		t.Errorf("expected -1, got %d", out)
	}
	if out, _ := b.Occupied(emptySquare); out != 0 {
		t.Errorf("expected 0, got %d", out)
	}
	if out, _ := b.Occupied(nonSquare); out != -2 {
		t.Errorf("expected -2, got %d", out)
	}
}

func TestIsOver(t *testing.T) {
	board := &Board{Turn: 1}
	board.PlacePiece('k', 1, 1, 1)
	board.PlacePiece('q', -1, 2, 2)
	board.PlacePiece('r', -1, 8, 2)
	if result := board.IsOver(); result != -2 {
		t.Errorf("Expected black wins, got a result of %d", result)
	}
	board.Board[1].Position.Y = 3
	if result := board.IsOver(); result != 1 {
		t.Errorf("Expected stalemate, got a result of %d", result)
	}
	board = &Board{Turn: -1}
	board.PlacePiece('k', 1, 1, 1)
	board.PlacePiece('k', -1, 8, 8)
	board.PlacePiece('b', 1, 6, 6)
	board.PlacePiece('r', -1, 8, 7)
	board.PlacePiece('r', -1, 7, 8)
	if over := board.IsOver(); over != 0 {
		t.Errorf("Black is in check but can block, IsOver still returned %d", over)
	}
}

func TestAllLegalMoves(t *testing.T) {
	board := &Board{Turn: -1}
	board.PlacePiece('k', -1, 1, 1)
	board.PlacePiece('k', 1, 8, 8)
	board.PlacePiece('p', -1, 4, 3)
	moves := board.AllLegalMoves()

	if moveslen := len(moves); moveslen != 4 {
		t.Errorf("Too many possible moves on the board. 4 moves expected, %d moves recieved", moveslen)
	}
	for i, m1 := range moves {
		for j, m2 := range moves {
			if m2 == m1 && i != j {
				t.Error("Duplicate moves returned, ", moves)
			}
		}
	}
}

func TestAttacking(t *testing.T) {
	board := &Board{Turn: 1}
	board.PlacePiece('k', 1, 1, 1)
	board.PlacePiece('r', 1, 2, 2)
	board.PlacePiece('p', 1, 2, 3)
	rook := board.Board[1]
	s := &Square{
		X: 4,
		Y: 2,
	}
	if !rook.Attacking(s, board) {
		t.Errorf("Rook not attacking on open line, should be attacking %+v from %+v", s, rook.Position)
	}
	s.X, s.Y = 2, 3
	if !rook.Attacking(s, board) {
		t.Errorf("Rook not attacking own piece, should be attacking %+v from %+v", s, rook.Position)
	}
	s.Y = 5
	if rook.Attacking(s, board) {
		t.Errorf("Rook attacking through own piece, should not be attacking %+v from %+v", s, rook.Position)
	}
}

func TestMakeMoveTo(t *testing.T) {
	board := &Board{Turn: 1}
	board.PlacePiece('k', 1, 1, 1)
	m := board.Board[0].makeMoveTo(2, 2)
	if m.Piece != 'k' {
		t.Error("Warped piece name from ", 'k', " to ", m.Piece)
	}
	if m.Begin != board.Board[0].Position {
		t.Errorf("Piece originated at 1,1. Current location: %+v, move begin: %+v", board.Board[0].Position, m.Begin)
	}
	if m.End.X != 2 || m.End.Y != 2 {
		t.Errorf("Incorrect ending square. Should be 2, 2, ended up at %+v", m.End)
	}
}



