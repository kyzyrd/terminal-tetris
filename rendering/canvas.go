package rendering

import (
	"fmt"
)

// -----------------------------------------------------------------------------------------------------------|CONST

const (
	DEFAULT         = Theme("\033[40;38;2;0;255;0;1m")     // аутентичный зеленый
	BLACK_AND_WHITE = Theme("\033[40;38;2;255;255;255;1m") // черно-белый
)

// -----------------------------------------------------------------------------------------------------------|TYPE

type Theme = string // псевдоним для ясности применения

type Img = []string // псевдоним, представляющих изображение игрового поля

type Canvas struct {
	_width, _height int
	_posX, _posY    int
	_theme          string
	_window         [][]rune
}

// -----------------------------------------------------------------------------------------------------------|FUNC

func GetNewCanvas(w, h int, theme Theme) *Canvas {
	// инициализируем атрибуты объекта
	newCanvas := new(Canvas)
	newCanvas._height = h
	newCanvas._width = w
	newCanvas._theme = theme

	newCanvas._posX = 0
	newCanvas._posY = 0

	// инициализация окна для отрисовки
	newCanvas._window = make([][]rune, newCanvas._height)
	for i := 0; i < len(newCanvas._window); i++ {
		newCanvas._window[i] = make([]rune, newCanvas._width)
	}
	newCanvas.Clear()

	return newCanvas
}

// -----------------------------------------------------------------------------------------------------------|METHOD

func (C *Canvas) SetImage(img Img, x, y int) {
	var runeRow []rune
	for rowId, row := range img {
		canvasY := rowId + y

		if canvasY < 0 || canvasY >= len(C._window) {
			continue
		}

		runeRow = []rune(row)
		for columnId, column := range runeRow {
			canvasX := columnId + x
			if canvasY < 0 || canvasX >= len(C._window[canvasY]) {
				continue
			}
			C._window[canvasY][canvasX] = column
		}
	}
}

// -----------------------------------------------------------------------------------------------------------|

func (C *Canvas) Print() {
	fmt.Printf("\033[H") // сохраняем текущее положение курсора
	fmt.Print(C._theme)
	for _, row := range C._window {
		fmt.Println(string(row))
	}

	fmt.Printf("\033[0m") // сбрасываем цвет и стилизацию
}

// -----------------------------------------------------------------------------------------------------------|

func (C *Canvas) Clear() {
	for y := 0; y < len(C._window); y++ {
		for x := 0; x < len(C._window[y]); x++ {
			C._window[y][x] = ' '
		}
	}
}
