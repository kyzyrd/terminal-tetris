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

// --------------------------------------------------------------------------------------------------|

func CreateNewEntry(name string, level, score int) *ScoreEntry {
	res := ScoreEntry{
		Score: EntryParam{
			Type:  TYPE_SCORE,
			Len:   4,
			Num:   1,
			Value: score,
		},
		Level: EntryParam{
			Type:  TYPE_LEVEL,
			Len:   4,
			Num:   1,
			Value: level,
		},
		Name: EntryParam{
			Type:  TYPE_STRING,
			Len:   byte(len(name)),
			Num:   byte(len([]rune(name))),
			Value: name,
		},
	}

	return &res
}

// -----------------------------------------------------------------------------------------|

func SortAppend(allEntries []*ScoreEntry, newScoreEntry *ScoreEntry) ([]*ScoreEntry, int) {
	allEntries = append(allEntries, newScoreEntry)
	Sort(allEntries)

	for i := 0; i < len(allEntries); i++ {
		if allEntries[i] == newScoreEntry {
			return allEntries, i
		}
	}

	return allEntries, -1
}

// -----------------------------------------------------------------------------------------|

func Sort(allEntries []*ScoreEntry) []*ScoreEntry {
	leftScore, rightScore := 0, 0
	leftLevel, rightLevel := 0, 0
	leftName, rightName := "", ""

	for i := 0; i < len(allEntries); i++ {
		for j := i + 1; j < len(allEntries); j++ {
			leftScore = allEntries[i].Score.Value.(int)
			rightScore = allEntries[j].Score.Value.(int)
			leftLevel = allEntries[i].Level.Value.(int)
			rightLevel = allEntries[j].Level.Value.(int)
			leftName = allEntries[i].Name.Value.(string)
			rightName = allEntries[j].Name.Value.(string)

			if leftScore < rightScore {
				allEntries[i], allEntries[j] = allEntries[j], allEntries[i]
			} else if leftScore == rightScore && leftLevel < rightLevel {
				allEntries[i], allEntries[j] = allEntries[j], allEntries[i]
			} else if leftScore == rightScore && leftLevel == rightLevel && leftName < rightName {
				allEntries[i], allEntries[j] = allEntries[j], allEntries[i]
			}
		}
	}

	return allEntries
}
