package main

import (
	"ai"
	"chess"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"time"

	"runtime"

	"github.com/gorilla/mux"
)

const (
	INDEX = `
		<html>
		<head>
			<title>Play Chess</title>
			<link rel="stylesheet" type="text/css" href="assets/chess/chessboardjs/css/chessboard-0.3.0.min.css">
			<script src="http://ajax.googleapis.com/ajax/libs/jquery/1.11.0/jquery.min.js"></script>
			<script src="assets/chess/chessjs/chess.min.js"></script>
		</head>
		<body>
			<div id="board" style="width: 400px"></div>
			<script src="assets/chess/chessboardjs/js/chessboard-0.3.0.js"></script>
			<script src="assets/chess/highlight-legals.js"></script>
		</body>
		</html>
	`

	PORT = ":6789"
	LOG  = true
)

var (
	inCMoves   = make(chan *chess.Move, 1)
	outCMoves  = make(chan *chess.Move, 1)
	quitCMoves = make(chan int, 1)
)

// Accepts a string such as "e4'"and converts it to the Square struct.
func stringToSquare(s string) chess.Square {
	var square chess.Square

	for i, b := range chess.Files {
		if b == s[0] {
			square.X = i + 1
			break
		}
	}

	for i, b := range chess.Ranks {
		if b == s[1] {
			square.Y = i + 1
			break
		}
	}

	return square
}

// Accepts a string such as "pe2-e4" and converts it to the Move struct.
func stringToMove(s string) *chess.Move {
	var move chess.Move
	move.Begin = stringToSquare(s[1:3])
	move.End = stringToSquare(s[4:])
	move.Piece = s[0]
	return &move
}

// Intended to run as a goroutine.
// Keeps track of the state of a single game, recieving and sending moves through the appropriate channel.
func game() {
	board := &chess.Board{Turn: 1}
	board.SetUpPieces()
	url := fmt.Sprintf("http://localhost%s", PORT)

	var cmdString string
	switch os := runtime.GOOS; os {
	case "darwin": // OS X
		cmdString = "open"
	case "linux": // Linux
		cmdString = "xdg-open"
	default: //windows, freebsd, openbds ...
		cmdString = ""
	}

	if cmdString == "open" || cmdString == "xdg-open" {
		cmd := exec.Command(cmdString, url)
		
		if _, err := cmd.Output(); err != nil {
			panic(err)
		}
	}

	rand.Seed(time.Now().UTC().UnixNano())

	for {
		select {
		case oppMove := <-inCMoves:
			for _, p := range board.Board {
				if p.Position.X == oppMove.Begin.X && p.Position.Y == oppMove.Begin.Y {
					oppMove.Piece = p.Name
					break
				}
			}
			board.ForceMove(oppMove)
			if LOG {
				fmt.Println("user move: ", oppMove.ToString())
				board.PrintBoard()
			}
			var myMove *chess.Move
			if moves, ok := ai.Book[board.ToFen()]; ok {
				myMove = stringToMove(moves[rand.Intn(len(moves))])
			} else {
				if m := ai.AlphaBeta(board, 4, ai.BLACKWIN, ai.WHITEWIN); m != nil {
					myMove = m
				} else {
					quitCMoves <- 1
					break
				}
			}

			board.ForceMove(myMove)
			outCMoves <- myMove
			if LOG {
				fmt.Println("ai move:", myMove.ToString())
				board.PrintBoard()
			}
		case <-quitCMoves:
			board.SetUpPieces()
			board.Turn = 1
		}
	}
}

// Serves the index, including relevant JS files
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, INDEX)
}

// Gets a move form from an AJAX request and sends it to the chest program
// Waits for a response from the chess program and sends that back to client
func chessHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("chess handler")
	w.Header().Set("Content-Type", "application/json")
	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	var promotion byte = 'q'

	if p, ok := r.Form["promotion"]; ok {
		promotion = p[0][0]
	}

	oppMove := &chess.Move{
		Begin:     stringToSquare(r.Form["from"][0]),
		End:       stringToSquare(r.Form["to"][0]),
		Promotion: promotion,
	}

	inCMoves <- oppMove
	myMove := <-outCMoves
	myMoveD := map[string]interface{}{"from": myMove.Begin.ToString(), "to": myMove.End.ToString(), "promotion": "q"}
	myMoveB, _ := json.Marshal(myMoveD)
	fmt.Fprint(w, string(myMoveB))
}

type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}
	return neuteredReaddirFile{f}, nil
}

type neuteredReaddirFile struct {
	http.File
}

func (f neuteredReaddirFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func main() {
	go game()
	
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/move", chessHandler)

	http.Handle("/", r)

	fs := justFilesFilesystem{http.Dir("assets/")}
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(fs)))

	http.ListenAndServe(PORT, nil)
}
