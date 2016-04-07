package ai

import (
	"fmt"
	"chess"
)

const (
	LOG = true
)

// Ref: http://web.cs.swarthmore.edu/~meeden/cs63/f05/minimax.html

// Standard minmax search with alpha beta pruning.
// Initial call: alpha set to lowest value, beta set to highest.
// Top level returns a move.
func AlphaBeta(b *chess.Board, depth int, alpha, beta float64) *chess.Move {
	if b.IsOver() != 0 || depth == 0 {
		return nil
	}
	
	var bestMove *chess.Move = nil
	var result float64
	
	moveList := orderedMove(b, false)
	
	if b.Turn == 1 {
		for _, move := range moveList {
			b.ForceMove(move)
			
			if move.Capture != 0 || b.IsCheck(b.Turn) {
				result = AlphaBetaChild(b, depth - 1, alpha, beta, true)
			} else {
				result = AlphaBetaChild(b, depth - 1, alpha, bet, false)
			}
			b.UndoMove(move)
			if result > alpha {
				alpha = result
				bestMove = move
				bestMove.Score = alpha
			}
			
			if alpha >= beta {
				bestMove = move
				bestMove.Score = alpha
				return bestMove
			}
		}
		
		if bestMove == nil {
			return b.AllLegalMoves()[0]
		}
		
		return bestMove
	} else {
		for _, move := range moveList {
			b.ForceMove(move)
			
			if move.Capture != 0 || b.IsCheck(b.Turn) {
				result = AlphaBetaChild(b, depth - 1, alpha, beta, true)
			} else {
				result = AlphaBetaChild(b, depth - 1, alpha, bet, false)
			}
			
			if LOG {
				fmt.Println(move.ToString(), result)
			}
			
			b.UndoMove(move)
			if result < alpha {
				beta = result
				bestMove = move
				bestMove.Score = beta
			}
			
			if beta <= alpha {
				bestMove = move
				bestMove.Score = beta
				return bestMove
			}
		}
		
		if bestMove == nil {
			return b.AllLegalMoves()[0]
		}
		
		return bestMove
	}
	
	if bestMove == nil {
		return b.AllLegalMoves()[0]
	}
	
	return bestMove
}

// Child level return an evaluation
func AlphaBetaChild(b *chess.Board, depth int, alpha, beta float64, volatile bool) float64 {
	var moveList []*chess.Move
	
	if b.IsOver() != 0 {
		return EvalBoard(b)
	} else if depth == 0 {
		if !volatile {
			return EvalBoard(b)
		}
		depth += 1
		moveList = orderedMove(b, true)
	} else {
		moveList = orderedMove(b, false)
	}
	
	var score float64
	
	if b.Turn == 1 {
		for _, move := range moveList {
			b.ForceMove(move)
			if !volatile && (move.Capture != 0 || b.IsCheck(b.Turn)) {
				score = AlphaBetaChild(b, depth -1, alpha, beta, true)
			} else {
				score = AlphaBetaChild(b, depth -1, alpha, beta, false)
			}
			
			b.UndoMove(move)
			if score > alpha {
				alpha = score
			}
			
			if alpha >= beta {
				return alpha
			}
		}
		
		return alpha
	} else {
		for _, move := range moveList {
			b.ForceMove(move)
			if !volatile && (move.Capture != 0 || b.IsCheck(b.Turn)) {
				score = AlphaBetaChild(b, depth -1, alpha, beta, true)
			} else {
				score = AlphaBetaChild(b, depth -1, alpha, beta, false)
			}
			
			b.UndoMove(move)
			if score < beta {
				alpha = score
			}
			
			if beta <=  alpha {
				return beta
			}
		}
		
		return beta
	}
	
	return 0
}








