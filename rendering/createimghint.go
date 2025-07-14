package rendering

// --------------------------------------------------------------------------------------|

// CreateImgHint создает изображение с подсказками для управления
func CreateImgHint() Img {
	hintImg := make(Img, 0, 6)

	hintImg = append(hintImg, "<: НАЛЕВО >: НАПРАВО")
	hintImg = append(hintImg, "    ^: ПОВОРОТ")
	hintImg = append(hintImg, "\\/: УСКОРИТЬ")
	hintImg = append(hintImg, "n: ПОКАЗАТЬ СЛЕДУЮЩУЮ")
	hintImg = append(hintImg, "h: СТЕРЕТЬ ЭТОТ ТЕКСТ")
	hintImg = append(hintImg, "  ПРОБЕЛ - СБРОСИТЬ")

	return hintImg
}
