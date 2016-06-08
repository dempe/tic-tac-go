package ai

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dempe/tictacgo/gamelogic"
)

func ComputerMove(b gamelogic.Board, aiType string) []int {
	fmt.Println("Computer's turn!")

	switch aiType {
	case "random":
		return moveRandomly(b)
	case "minimax":
		return moveUsingMiniMax(b)
	default:
		return moveRandomly(b)
	}
}

func moveUsingMiniMax(b gamelogic.Board) []int {
	source := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(source)
	row := rand.Intn(3)
	col := rand.Intn(3)

	tiles := b.GetTiles()

	for tiles[row][col] != 0 {
		row = rand.Intn(3)
		col = rand.Intn(3)
	}

	return []int{row, col}
}

func moveRandomly(b gamelogic.Board) []int {
	source := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(source)
	row := rand.Intn(3)
	col := rand.Intn(3)

	tiles := b.GetTiles()

	for tiles[row][col] != 0 {
		row = rand.Intn(3)
		col = rand.Intn(3)
	}

	return []int{row, col}
}
