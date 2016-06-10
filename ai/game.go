package ai

import (
	"container/list"
	"fmt"

	"github.com/dempe/tictacgo/gamelogic"
)

type Score struct {
	Score         int
	NotDetermined bool
}

type GameState struct {
	position  [2]int
	subStates *list.List
}

func whoseTurn(b Board) int {
	var xcount, ycount int

	for i := 0; i < 3; i++ {
		for j := 0; j < len(tiles[i]); j++ {
			if tiles[i][j] == 1 {
				ycount++
			} else if tiles[i][j] == 2 {
				xcount++
			}
		}
	}

	if xcount == ycount {
		return 2
	}

	return 1
}

func CalculateScore(b gamelogic.Board, mark string) (Score, error) {
	if mark != "X" && mark != "O" {
		return Score{0, true}, fmt.Errorf("Unrecognized mark, %s.  Must be X or O", mark)
	}

	winner := b.GetWinningPlayer()

	if winner.Undetermined {
		return Score{0, false}, nil
	} else if mark == winner.Mark {
		return Score{1, false}, nil
	}

	return Score{-1, false}, nil
}

func CalculatePossibleMoves(b gamelogic.Board, mark string) (*list.List, error) {
	if mark != "X" && mark != "O" {
		return list.New(), fmt.Errorf("Unrecognized mark, %s.  Must be X or O", mark)
	}

	positions := list.New()
	tiles := b.GetTiles()

	for i := 0; i < 3; i++ {
		for j := 0; j < len(tiles[i]); j++ {
			if tiles[i][j] == 0 {
				positions.PushBack([2]int{i, j})
			}
		}
	}

	return positions, nil
}
