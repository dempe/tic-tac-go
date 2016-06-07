package main

import "fmt"

type Board struct {
	tiles [3][3]int
}

func main() {
	b := constructInitialBoard()
	playing := true

	for playing {
		displayPrompt()
		printBoard(b)
	}
}

func displayPrompt() {
	fmt.Println("Please input your move in the form:  row,col type.  Example:  0,2 X")
	fmt.Println("Rows and columns go from 0 - 2, while types can be one of 'X' or 'O'")
}

func constructInitialBoard() Board {
	return Board{[3][3]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0}}}
}

func printBoard(b Board) {
	fmt.Println("Here is the current board:")
	for i := 0; i < 3; i++ {
		fmt.Println()
		printRow(b.tiles[i])
	}
}

func printRow(row [3]int) {
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			fmt.Print("|")
		}

		switch i {
		case 1:
			printPlace(row[0])
		case 4:
			printPlace(row[1])
		case 7:
			printPlace(row[2])
		default:
			fmt.Print(" _ ")
		}
	}

	fmt.Println("|")
}

func printPlace(value int) {
	switch value {
	case 0:
		fmt.Print("")
	case 1:
		fmt.Print("O")
	case 2:
		fmt.Print("X")
	}
}
