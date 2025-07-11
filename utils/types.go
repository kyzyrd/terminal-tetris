package utils

import (
	"sync"
	"terminal-tetris2/shape"

	"github.com/nsf/termbox-go"
)

//--------------------------------------------------------------------------------// TYPES

// GameVar - структура для хранения глобальных переменных игры
type GameVar struct {
	ShowNS   bool // показывать следующую фигуру
	HideHint bool // скрывать подсказки
	Quit     bool // условие завершения программы

	CurrentTime int64 // текущее время в миллисекундах
	Timer       Timers
	Stat        Stats
	Object      GameObjts
}

type Timers struct {
	Fell  int64 // таймер для падения фигуры
	Step  int64 // таймер для шага игры
	Clear int64 // таймер для очистки линий
}

type Stats struct {
	FullLines []int // индексы полных линий
	LastLine  int   // количество очков за последнюю линию
	DelLines  int   // количество удаленных линий
	Score     int   // игровые очки
	Level     int   // игровой уровень
}

type GameObjts struct {
	Shape     *shape.Shape  // текущая фигура
	NextShape *shape.Shape  // следующая фигура
	OldBricks []shape.Brick // старые кирпичики, которые уже упали на игровое поле
}

type Controls struct {
	Ev         termbox.Event // событие клавиатуры
	NewEvent   bool          // новое событие клавиатуры
	MutexEvent sync.Mutex    // мьютекс для синхронизации событий клавиатуры
}

// ScoreEntry represents a single score record
type ScoreEntry struct {
	Name  EntryParam
	Level EntryParam
	Score EntryParam
}

// Parameters of every entry (metadata)
type EntryParam struct {
	Type  byte
	Len   byte
	Num   byte
	Value any
}
