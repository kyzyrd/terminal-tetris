package rendering

//---------------------------------------------------------------------------------------//

// CreateImgLogo создает изображение для счета и уровня
func CreateImgLogo() Img {
	gameLogo := make(Img, 0, 3)

	gameLogo = append(gameLogo, "[]    ")
	gameLogo = append(gameLogo, "ТЕТРИС")
	gameLogo = append(gameLogo, "    []")

	return gameLogo
}

//---------------------------------------------------------------------------------------//

func CreateImgLevel() Img {
	level := make(Img, 0, 1)
	level = append(level, "ВАШ УРОВЕНЬ? (0-9) -")
	return level
}
