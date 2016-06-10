package ai

import (
	"container/list"
	"errors"
	"fmt"

	"github.com/dempe/tictacgo/gamelogic"
)

type Score struct {
	Score         int
	NotDetermined bool
}

type GameState struct {
	board     gamelogic.Board
	turn      int
	SubStates *list.List
}

func NewGameState(b gamelogic.Board) (*GameState, error) {
	turn := whoseTurn(b)
	moves := CalculatePossibleMoves(b)
	subStates := list.New()

	for e := moves.Front(); e != nil; e = e.Next() {
		newBoard := b.Copy()

		value, ok := e.Value.([]int)

		if !ok {
			return nil, errors.New("Expected type []int")
		}

		//fmt.Printf("Calculating new state for position %d,%d:\n", value[0], value[1])

		newBoard.PlaceMove(value, gamelogic.DecodeValue(turn))
		//newBoard.PrintBoard()
		subState, _ := NewGameState(*newBoard)
		subStates.PushBack(subState)
	}

	return &GameState{b, turn, subStates}, nil
}

func (g *GameState) Sum() int {
	return g.sumHelper(0)
}

func (g *GameState) sumHelper(sum int) int {
	if g.SubStates.Len() == 0 {
		return 1
	}

	for e := g.SubStates.Front(); e != nil; e = e.Next() {
		value, ok := e.Value.(*GameState)

		if !ok {
			fmt.Println("Unexpected value")
		}

		sum += value.sumHelper(0)
	}

	return sum
}

func whoseTurn(b gamelogic.Board) int {
	var xcount, ycount int
	tiles := b.GetTiles()

	for i := 0; i < 3; i++ {
		for j := 0; j < len(tiles[i]); j++ {
			if tiles[i][j] == 1 {
				ycount++
			} else if tiles[i][j] == 2 {
				xcount++
			}
		}
	}

	// X always goes first
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

func CalculatePossibleMoves(b gamelogic.Board) *list.List {
	positions := list.New()
	tiles := b.GetTiles()

	if !b.GetWinningPlayer().Undetermined {
		return positions
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < len(tiles[i]); j++ {
			if tiles[i][j] == 0 {
				positions.PushBack([]int{i, j})
			}
		}
	}

	return positions
}
