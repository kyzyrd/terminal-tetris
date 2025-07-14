package utils

import (
	"math"

	"github.com/nsf/termbox-go"
)

//------------------------------------------------------------------------------------// CONST

const (
	TYPE_NIL = byte(iota)
	TYPE_SCORE
	TYPE_LEVEL
	TYPE_STRING
)

// константы для игры
const (
	T_QUANT             = 1             // квант времени
	TIMER_OFF           = math.MaxInt64 // значение для выключения таймера
	TICK                = 60            // время между шагами игры в миллисекундах
	BASE_SPEED          = 600           // базовая скорость падения фигуры в миллисекундах
	CLEAR_TIME          = 600           // время очистки линий в миллисекундах
	FELL_TIME           = 270           // время фиксации падения фигуры в миллисекундах
	SPEED_INCREASER     = 60            // множитель на уровень в миллисекундах
	FELL_TIME_INCREASER = 15            // множитель на уровень в миллисекундах
)

const (
	BASE_SCORE         = 19 // базовое количество очков за линию
	LINES_LEVEL_UP     = 10 // количество линий для повышения уровня
	DEL_LINES_LEVEL_UP = 4  // количество строк при одновременном удалении
	LEVEL_BONUS        = 3  // бонус за уровень
	FORESIGHT_BONUS    = 5  // бонус за не использование подсказки следующей фигуры
	MAX_LEVEL          = 9  // максимальный уровень
)

// размеры игрового окна
const (
	WINDOW_HEIGHT = 24
	WINDOW_WIDTH  = 67
)

// размеры игрового поля
const (
	HEIGHT = 20
	WIDTH  = 10
)

// клавиши управления
const (
	KEY_LEFT      = termbox.KeyArrowLeft  // клавиша влево
	KEY_RIGHT     = termbox.KeyArrowRight // клавиша вправо
	KEY_ROTATE    = termbox.KeyArrowUp    // клавиша вверх (вращение фигуры)
	KEY_DROP      = termbox.KeySpace      // клавиша пробел (падение фигуры)
	KEY_SHOW_NEXT = 'n'                   // клавиша для показа следующей фигуры
	KEY_SPEEDUP   = termbox.KeyArrowDown  // клавиша вниз (ускорение падения фигуры)
	KEY_HIDE_HINT = 'h'                   // клавиша для скрытия подсказок
	KEY_LEVEL_UP  = 'u'                   // клавиша для повышения уровня
	KEY_QUIT      = termbox.KeyEsc        // клавиша для выхода из игры
)

const (
	SCOREBOARD_FILE = "scoreboard"
	TOP_SCORES      = 15
)
