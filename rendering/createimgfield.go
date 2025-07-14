package rendering

import (
	"fmt"
	"strings"
	"terminal-tetris2/shape"
	"terminal-tetris2/utils"
)

// --------------------------------------------------------------------------------------------|

// CreateImgField создает изображение игрового поля
func CreateImgField(oldBricks, shapeBricks []shape.Brick) Img {
	fieldImage := make([]string, 0, utils.HEIGHT)

	field := createTmpField(oldBricks, shapeBricks)

	for i := range field {
		row := createRowString(field, i)
		fieldImage = append(fieldImage, fmt.Sprintf("<!%s!>", row))
	}
	// добавляем нижние границы игрового поля
	fieldImage = append(fieldImage, fmt.Sprintf("<!%s!>", strings.Repeat("==", utils.WIDTH)))
	fieldImage = append(fieldImage, fmt.Sprintf("  %s  ", strings.Repeat("\\/", utils.WIDTH)))

	return Img(fieldImage)
}

// --------------------------------------------------------------------------------------------|

// createTmpField создает временное игровое поле для отображения
func createTmpField(oldBricks, shapeBricks []shape.Brick) [][]shape.Brick {
	field := make([][]shape.Brick, utils.HEIGHT) // создаем двумерный срез для игрового поля

	// инициализируем каждую строку игрового поля
	for i := 0; i < len(field); i++ {
		field[i] = make([]shape.Brick, utils.WIDTH)
	}

	// инициализируем каждую ячейку игрового поля как невидимую
	for y := 0; y < utils.HEIGHT; y++ {
		for x := 0; x < utils.WIDTH; x++ {
			field[y][x].Visible = shape.HIDE
		}
	}

	// устанавливаем видимость старых кирпичиков
	for _, brick := range oldBricks {
		if brick.Y > -1 {
			field[brick.Y][brick.X].Visible = brick.Visible
		}
	}

	// устанавливаем видимость текущей фигуры
	for _, brick := range shapeBricks {
		if brick.Y > -1 {
			field[brick.Y][brick.X].Visible = brick.Visible
		}
	}

	return field
}

// --------------------------------------------------------------------------------------------|

// createRowString создает строку для отображения игрового поля
func createRowString(field [][]shape.Brick, row int) string {
	result := ""
	for _, cell := range field[row] {
		if cell.Visible == shape.SHOW {
			result += "[]"
		} else {
			result += " ."
		}
	}
	return result
}
