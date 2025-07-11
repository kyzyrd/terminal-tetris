package rendering

import (
	"strings"
	"terminal-tetris2/shape"
)

// --------------------------------------------------------------------------------------|

// CreateImgNS создает изображение для следующей фигуры
func CreateImgNS(nextShape shape.Shape) Img {
	bricks := nextShape.GetBricks()

	minX, minY, maxX, maxY := getBounds(bricks)

	width := (maxX - minX + 1) * 2
	height := maxY - minY + 1

	image := createEmptyImage(width, height)

	placeBricks(image, bricks, minX, minY)

	return image
}

// --------------------------------------------------------------------------------------|

// getBounds находит минимальные и максимальные координаты для массива кирпичей
func getBounds(bricks []shape.Brick) (minX, minY, maxX, maxY int) {
	if len(bricks) == 0 {
		return 0, 0, 0, 0
	}

	minX, minY = bricks[0].X, bricks[0].Y
	maxX, maxY = minX, minY

	for _, brick := range bricks[1:] {
		minX = min(minX, brick.X)
		minY = min(minY, brick.Y)
		maxX = max(maxX, brick.X)
		maxY = max(maxY, brick.Y)
	}

	return minX, minY, maxX, maxY
}

// --------------------------------------------------------------------------------------|

// createEmptyImage создает пустое изображение заданного размера
func createEmptyImage(width, height int) Img {
	image := make(Img, height)
	for i := range image {
		image[i] = strings.Repeat(" ", width)
	}
	return image
}

// --------------------------------------------------------------------------------------|

// placeBrickOnImage размещает один кирпич на изображении
func placeBrickOnImage(image Img, brick shape.Brick, offsetX, offsetY int) {
	x := (brick.X - offsetX) * 2
	y := brick.Y - offsetY

	// Заменить символы в строке
	row := []rune(image[y])
	row[x] = '['
	row[x+1] = ']'
	image[y] = string(row)
}

// --------------------------------------------------------------------------------------|

// placeBricks размещает все кирпичи на изображении
func placeBricks(image Img, bricks []shape.Brick, offsetX, offsetY int) {
	for _, brick := range bricks {
		placeBrickOnImage(image, brick, offsetX, offsetY)
	}
}
