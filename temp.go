package main
import "fmt"

var boardWidth = 4
var pieces = []rune(`ABCD`)
var directions = []int{-7, 7, 1, -1}
var runeBoard = []rune(
	`XXXXXX
X..C.X
XAACCX
X.BD.X
X.BDDX
XXXXXX
`)
var container = [][]rune{}
var moves int

func main() {
	run(runeBoard, 0)
	// var temp = move(pieces[0], directions[0], runeBoard)
	// fmt.Println(string(temp))
}

func run(currentBoard []rune, moves int) {
	var m = moves
	var c = copyBoard(currentBoard)
	//check if board exists and if 
	if(len(container) != 0 || boardExists(c)){
		fmt.Println(m)	
	}
	container = append(container, c)
	fmt.Println(string(c))
	//make a new board and recurse
	for i:=0;i<len(pieces);i++{
		for j:=0;j<len(directions);j++{
			if(movable(pieces[i], directions[j], c)){
				var newBoard = copyBoard(c)
				var temp = move(pieces[i], directions[j], newBoard)
				fmt.Println(string(pieces[i]), directions[j])
				m++
				run(temp, m)	
			}
		}
	}
}

func boardExists(board []rune) bool {
	for i:=0;i<len(container);i++{
		if(equalBoards(container[i], board)){
			return true
		}
	}
	return false
}

func move(piece rune, direction int, board []rune) []rune {
	var temp = copyBoard(board);
	// delete piece
	for i:=0;i<len(temp);i++ {
		if (temp[i] == piece) {
			temp[i] = '.'
		}
	}
	//attempt to move piece
	for i:=0;i<len(board);i++ {
		if (board[i] == piece) {
			temp[i + direction] = board[i]
		}
	}
	return temp
}

func movable(piece rune, direction int, board []rune) bool {
	var temp = copyBoard(board);
	// delete piece
	for i:=0;i<len(temp);i++ {
		if (temp[i] == piece) {
			temp[i] = '.'
		}
	}
	//attempt to move piece
	for i:=0;i<len(board);i++ {
		if (board[i] == piece && temp[i + direction] != '.') {
			return false
		}
	}
	// fmt.Println(string(temp))
	return true
}

func equalBoards(a, b []rune) bool {
	for i := 0; i < len(a); i++{
		if(a[i] != b[i]){
			return false
		}
	}
	return true
}


func printBoard(board string){
	for i:=0; i<len(board); i++ {
		fmt.Print(string(board[i]))
		if(i % (boardWidth + 2) == boardWidth + 1){
			fmt.Print("\n");
		}
	}
}

func copyBoard(board []rune) []rune{
	var newBoard = make([]rune, len(runeBoard))
	copy(newBoard, board)
	return newBoard
}

