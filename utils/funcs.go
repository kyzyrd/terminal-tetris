package utils

import (
	"terminal-tetris2/shape"
)

// --------------------------------------------------------------------------------------------------|

// CheckCollision проверяет коллизию между фигурами и игровым полем
func CheckCollision(RawShapes ...[]shape.Brick) bool {

	existBrick := make(map[shape.Brick]bool)

	for _, currentShape := range RawShapes {
		for _, brick := range currentShape {
			if _, ok := existBrick[brick]; ok { // если кирпичик уже существует в мапе
				return true // значит есть коллизия
			}

			// если кирпичик выходит за границы игрового поля по X, то есть коллизия
			if (brick.X < 0) || (brick.X >= WIDTH) {
				return true
			}
			// если кирпичик выходит за границы игрового поля по Y, то есть коллизия
			if (brick.Y < 0) || (brick.Y >= HEIGHT) {
				return true
			}

			// добавляем кирпичик в мапу, чтобы проверить его на существование в следующей итерации
			existBrick[brick] = false
		}
	}
	return false // если ни один кирпичик не пересекается с другими, то коллизии нет
}
