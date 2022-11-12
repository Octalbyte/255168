package main

import (
	"fmt"
)

type State [][]int

func main(){
var board State
var row []int
row = append(row, 0)
row = append(row, 0)
row = append(row, 0)

board = append(board, row)
board = append(board, row)
board = append(board, row)

/*for i := 0; i < len(board); i++ {
	for j := 0; j < len(board[i]); j++ {
		board[i][j] = 0 
	}
}*/	
var bd []State

bd = append(bd, board)
//fmt.Println(bd)
jdi(bd, 0, 0, 0)

}

func jdi(boards []State, totones int, totfours int, totdraws int) {
	
	draws := 0
	ones := 0
	fours := 0
	var nexts []State
	var unfinished []State
	for b := 0; b < len(boards); b++ {
		//fmt.Println(boards[b])		
		st := evalPos(boards[b])
		if st == 1 {ones++}
		if st == 4 {fours++}
		if st == 2 {draws++}
		if st == 3 { unfinished = append(unfinished, boards[b])}
	}
	//fmt.Print("Layer: %d X-wins, %d O-wins, %d draws and %d unfinished games ", ones, fours, draws, len(unfinished))
	totones += ones
	totfours += fours
	totdraws += draws
	//fmt.Print(unfinished)
	for un := 0; un < len(unfinished); un++ {
		nx := getNext(unfinished[un])
		fmt.Print(nx, "\n") //You may want to comment this out 
		nexts = append(nexts, nx...)
	}
	if len(nexts) != 0 {
			fmt.Print("\n TOTAL STATS")
			fmt.Print(totones, " X wins")
			fmt.Print(totfours, " O wins")
			fmt.Print(totdraws, " draws")
			jdi(nexts, totones, totfours, totdraws)
	} else {
			fmt.Print("\nTOTAL STATS ")
			fmt.Print(totones, " X wins ")
			fmt.Print(totfours, " O wins ")
			fmt.Print(totdraws, " draws ")
	}
}

func evalPos(board State) int {
	sumRow0 := board[0][0] + board[0][1] + board[0][2]
	sumRow1 := board[1][0] + board[1][1] + board[1][2]
	sumRow2 := board[2][0] + board[2][1] + board[2][2]
	sumCol0 := board[0][0] + board[1][0] + board[2][0]
	sumCol1 := board[0][1] + board[1][1] + board[2][1]
	sumCol2 := board[0][2] + board[1][2] + board[2][2]
	sumDg1 := board[0][0] + board[1][1] + board[2][2]
	sumDg2 := board[0][2] + board[1][1] + board[2][0]
	sumBoard := sumRow0 + sumRow1 + sumRow2
	if sumRow0 == 3 || sumRow1 == 3 || sumRow2 == 3 || sumCol0 == 3 || sumCol1 == 3 || sumCol2 == 3 || sumDg1 == 3 || sumDg2 == 3 {
		return 1
	}
	if sumRow0 == 12 || sumRow1 == 12 || sumRow2 == 12 || sumCol0 == 12 || sumCol1 == 12 || sumCol2 == 12 || sumDg1 == 12 || sumDg2 == 12 {
		return 4
	}
	if sumBoard == 21 {
		return 2	
	} 
	return 3
}

func getNext(board State) []State {
	ones := 0
	fours := 0
	for row := 0; row < len(board); row++ {
		for place := 0; place < len(board[row]); place++ {
			if board[row][place] == 1 { ones++}
			if board[row][place] == 4 { fours++}
		}
	}
	play := 0
	if ones > fours {
			play = 4
		} else {
			play = 1 	
		}
	var next []State
	for row := 0; row < len(board); row++ {
		for place := 0; place < len(board[row]); place++ {
			if board[row][place] == 0 {
				
				npos := make(State, len(board))
				for row := 0; row < len(board); row++ {
					for place := 0; place < len(board[row]); place++ {
						npos[row] = append(npos[row], board[row][place])
					}
				}
				//fmt.Print(board)
				//fmt.Print(npos)
				npos[row][place] = play
					
				next = append(next, npos)
				
				//fmt.Print(next[len(next)-1], "\n")
			}
		}
	}
	return next
}
