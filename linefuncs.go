package mechs

import (
	"terminal-tetris2/shape"
	"terminal-tetris2/utils"
)

// --------------------------------------------------------------------------------------|

// FindFullLines находит полные линии в игровом поле
func FindFullLines(bricks []shape.Brick) []int {

	sample := -1
	count := 0
	fullLines := make([]int, 0)

	for i := 0; i < len(bricks); i++ {
		y := bricks[i].Y

		if y != sample {
			sample = y
			count = 0
		}
		count++

		if count == utils.WIDTH {
			fullLines = append(fullLines, i-utils.WIDTH+1) // начальный индекс включая
			fullLines = append(fullLines, i+1)             // конечный индекс не включая
		}
	}
	return fullLines
}

// --------------------------------------------------------------------------------------|

// ClearLines удаляет полные линии из массива bricks
func ClearLines(bricks *[]shape.Brick, fullLines []int) {
	start := 0
	end := 0

	// идем с конца, чтобы не выходить за пределы bricks, так как он уменьшается
	for i := len(fullLines) - 1; i > -1; i -= 2 {
		start = fullLines[i-1]
		end = fullLines[i]
		*bricks = append((*bricks)[:start], (*bricks)[end:]...)
	}
}

// --------------------------------------------------------------------------------------|

// BlinkLines мигает полными линиями, чтобы показать их удаление
func BlinkLines(bricks []shape.Brick, fullLines []int) {
	start := 0
	end := 0

	// идем с конца, чтобы не выходить за пределы bricks, так как он уменьшается
	for i := len(fullLines) - 1; i > -1; i -= 2 {
		start = fullLines[i-1]
		end = fullLines[i]
		for i := start; i < end; i++ {
			bricks[i].Visible = !bricks[i].Visible
		}
	}
}

// --------------------------------------------------------------------------------------|

// DropLines опускает линии вниз, чтобы заполнить пустоты после удаления полных линий
func DropLines(bricks []shape.Brick) {

	sample := utils.HEIGHT
	lastLine := utils.HEIGHT - 1

	// идем с конца, чтобы не выходить за пределы bricks
	for i := len(bricks) - 1; i > -1; i-- {
		y := bricks[i].Y

		if y != sample {
			for i := 0; i < len(bricks); i++ {
				if bricks[i].Y == y {
					bricks[i].Y = lastLine
				}
			}
			sample = bricks[i].Y
			lastLine--
		}
	}
}
