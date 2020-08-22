package main

import "fmt"

func displayBoard(board [8][8]rune){
	for _,row := range board{
		for _, col := range row{
			fmt.Printf("| %c ", col)
		}
		println("|\n---------------------------------")
	} 
}


// black -> top -> lowercase
//white -> bottom -> uppercase
func main(){
	var chessBoard [8][8]rune

	//putting black pieces
	chessBoard[0][0] = 'r'
	chessBoard[0][1] = 'n'
	chessBoard[0][2] = 'b'
	chessBoard[0][3] = 'q'
	chessBoard[0][4] = 'k'
	chessBoard[0][5] = 'n'
	chessBoard[0][6] = 'b'
	chessBoard[0][7] = 'r'

	//putting white pieces
	chessBoard[7][0] = 'R'
	chessBoard[7][1] = 'N'
	chessBoard[7][2] = 'B'
	chessBoard[7][3] = 'Q'
	chessBoard[7][4] = 'K'
	chessBoard[7][5] = 'N'
	chessBoard[7][6] = 'B'
	chessBoard[7][7] = 'R'



	//putting Pawns
	for column := range chessBoard{
		chessBoard[1][column] = 'p'
		chessBoard[6][column] = 'P'
	}

	displayBoard(chessBoard)
}

