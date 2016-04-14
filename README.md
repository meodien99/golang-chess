# golang-chess
This project is written for improving my knowlegde in Go

# How to install
After installing [GO](https://golang.org/doc/install), clone this repository then in root of directory run :
```
go install
```
If you get NOTCE such as "no install location outside GOPATH", make sure you setted your GOPATH enviroment variable (in OSX and Linux, make sure you setted $GOBIN : 
```
export GOBIN=$GOPATH/bin
```

# Project Structure
Main project organization is inside `src/` directory, include :
### ai/
  * Handles everything related to the AI .
  * I'm useing [alphabeta search and board evaluation](https://en.wikipedia.org/wiki/Alpha%E2%80%93beta_pruning)

### chess/
  * Handles the game engine (Eg: game state storage or piece movement, ...).
  * Contains helper functions that are entirely reliant on the rules of Chess, such as whether a given square on a board is occupied.


Folder `assets` contains all of files for displaying chess on browser (Thankfully [chessboardjs](http://chessboardjs.com/) for beautiful chess board, and [chessjs](https://github.com/jhlywa/chess.js) for validating piece movement in client-side).

I'm used [mux](https://github.com/gorilla/mux) for implementing requests router and dispatchers as well as handling AJAX request from Go, make sure that it's installed, if not, run cmd `go get` from your root directory, [see more details](https://golang.org/cmd/go/).

# Testing 
run `go test` inside root directory to ensure everything works

After all, run :
```
go build
```

brower will be opened immediately (Linux & Mac) , if you are using others, simply execute the output file has compiled.

# Resources
  * [GoLang Documentation](http://golang.org/)
  * [Chess Programming](http://chessprogramming.wikispaces.com/)
  * [AI Search Theory](http://www.frayn.net/beowulf/theory.html)
  * [Forsyth Edwards Notation Books](http://www.365chess.com/)


