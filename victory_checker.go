package tictacgo

func isGameOver(b Board) bool {
	return checkRows(b) || checkDiagonals(b) || checkCols(b)
}

func checkCols(b Board) bool {
	for i := 0; i < gridSize; i++ {
		firstCell := b.tiles[0][i]

		if firstCell == 0 {
			continue
		}

		if firstCell == b.tiles[1][i] && b.tiles[1][i] == b.tiles[2][i] {
			return true
		}
	}

	return false
}

func checkDiagonals(b Board) bool {
	topLeft := b.tiles[0][0]
	topRight := b.tiles[0][2]
	middle := b.tiles[1][1]
	bottomLeft := b.tiles[2][0]
	bottomRight := b.tiles[2][2]

	if middle == 0 {
		return false
	}

	return (topLeft == middle && middle == bottomRight) || (topRight == middle && middle == bottomLeft)
}

func checkRows(b Board) bool {
	for i := 0; i < gridSize; i++ {
		firstCell := b.tiles[i][0]

		if firstCell == 0 {
			continue
		}

		if firstCell == b.tiles[i][1] && b.tiles[i][1] == b.tiles[i][2] {
			return true
		}
	}

	return false
}
