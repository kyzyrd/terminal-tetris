package rendering

import (
	"fmt"
)

// --------------------------------------------------------------------------------------|

// CreateImgScore создает изображение для счета и уровня
func CreateImgScore(score int, level int, delLines int) Img {
	scoreImage := make(Img, 0, 6)

	// заполняем срез значениями
	scoreImage = append(scoreImage, fmt.Sprintf("ПОЛНЫХ СТРОК: %d", delLines))
	scoreImage = append(scoreImage, fmt.Sprintf("УРОВЕНЬ:      %d", level))
	scoreImage = append(scoreImage, fmt.Sprintf("  СЧЕТ: %d", score%1000))

	// добавляем звездочки для тысяч
	stars := score / 1000
	thousands := ""
	for i := 0; i < stars; i++ {
		if i > 0 && i%5 == 0 {
			scoreImage = append(scoreImage, thousands)
			thousands = ""
		}
		thousands += " *"
	}
	if thousands != "" {
		scoreImage = append(scoreImage, thousands)
	}

	return scoreImage
}
