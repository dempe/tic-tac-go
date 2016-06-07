package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	tiles [3][3]int
}

func main() {
	b := constructInitialBoard()
	playing := true

	for playing {
		readInput(b)
		printBoard(b)
		playing = false
	}
}

func readInput(b Board) {
	fmt.Println("Please input your move in the form:  row,col,type.  Example:  0,2,X")
	fmt.Println("Rows and columns go from 0 - 2, while types can be one of 'X' or 'O'")

	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	position, mark, err := parseInput(input)

	if err != nil {
		fmt.Println(err)
	}
}

func parseInput(input string) ([]int, string, error) {
	arr := strings.Split(input, ",")

	if len(arr) < 3 {
		return nil, "", errors.New("Unrecognized input")
	}

	row, rowErr := strconv.Atoi(arr[0])
	col, colErr := strconv.Atoi(arr[1])
	mark := strings.Replace(arr[2], "\n", "", 1)

	if rowErr != nil || colErr != nil {
		return nil, "", errors.New("Unrecognized input")
	}

	if row < 0 || row > 2 || col < 0 || col > 2 {
		return nil, "", errors.New("Unrecognized input")
	}

	if mark != "X" && mark != "O" {
		return nil, "", errors.New("Unrecognized input" + mark)
	}

	return []int{row, col}, mark, nil
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
