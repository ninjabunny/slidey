package main
import "fmt"

var boardWidth = 4
var runeBoard = []rune(
	`XXXXXX
X..C.X
XAACCX
X.BD.X
X.BDDX
XXXXXX
`)

func main() {
	// c := board
	var temp = make([]rune, len(runeBoard))
	copy(temp, runeBoard)

	// temp[0] = 'B'
	if(compareBoards(temp, runeBoard)){
		fmt.Println(string(runeBoard))
	} else {
		fmt.Println("tis false")	
	}
}

func compareBoards(a, b []rune) bool {
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