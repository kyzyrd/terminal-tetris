package shape

import (
	"math/rand/v2"
)

//---------------------------------------------------------------------------------------//CONST

const (
	SHAPE_0 = iota
	SHAPE_I
	SHAPE_L
	SHAPE_J
	SHAPE_S
	SHAPE_Z
	SHAPE_T
	NUM_SHAPES
)

const (
	ROTATE_FORWARD   = RotDir(true)
	ROTATE_BACKWARDS = RotDir(false)
)

const (
	HIDE = true
	SHOW = false
)

//---------------------------------------------------------------------------------------//TYPE

// структура Brick хранит координаты кирпичика и его видимость
type Brick struct {
	X       int
	Y       int
	Visible bool
}

// структура Vec2 хранит координаты точки в двумерном пространстве
type Vec2 struct {
	X int
	Y int
}

// структура _Orientation хранит ориентацию фигуры
type _Orientation struct {
	Bricks []Brick
	Next   *_Orientation
	Prev   *_Orientation
}

// псевдоним для типа _ShapeType, который используется для обозначения типа фигуры
type _ShapeType = uint8

// структура Shape представляет фигуру в игре, содержит ориентацию, тип фигуры и позицию
type Shape struct {
	_orientation *_Orientation
	_shapeType   _ShapeType
	_pos         Vec2
}

// тип RotDir используется для обозначения направления вращения фигуры
type RotDir = bool

// глобальный массив для хранения ориентаций фигур
var _shapesOrientation [7]*_Orientation = [7]*_Orientation{}

//---------------------------------------------------------------------------------------//FUNC

// метод GetNewShape создает новую фигуру случайного типа
func GetNewShape() *Shape {
	newShape := new(Shape)
	newShape._shapeType = _ShapeType(rand.Int() % NUM_SHAPES)
	newShape._orientation = _shapesOrientation[newShape._shapeType]

	return newShape
}

//---------------------------------------------------------------------------------------//

// метод SortBricksByY сортирует массив кирпичиков по координате Y
func SortBricksByY(bricks []Brick) {
	for i := 0; i < len(bricks); i++ {
		for j := i + 1; j < len(bricks); j++ {
			if bricks[i].Y > bricks[j].Y {
				bricks[i], bricks[j] = bricks[j], bricks[i]
			}
		}
	}
}

//---------------------------------------------------------------------------------------//METHOD

// метод GetBrick
func (S *Shape) GetBricks() []Brick {
	bricks := S._orientation.Bricks
	bricksCopy := make([]Brick, len(bricks))

	// смещаем координаты кирпичиков относительно позиции фигуры
	for i := 0; i < len(bricksCopy); i++ {
		bricksCopy[i].X = bricks[i].X + S._pos.X
		bricksCopy[i].Y = bricks[i].Y + S._pos.Y
	}

	return bricksCopy
}

//---------------------------------------------------------------------------------------//

func (S *Shape) GetPosition() Vec2 {
	return S._pos
}

//---------------------------------------------------------------------------------------//

func (S *Shape) Move(dx, dy int) Vec2 {
	S._pos.X += dx
	S._pos.Y += dy
	return Vec2{X: dx, Y: dy}
}

//---------------------------------------------------------------------------------------//

// вращение фигуры против часовой стрелки
func (S *Shape) Rotate(rotDir RotDir) {
	if rotDir == ROTATE_FORWARD {
		S._orientation = S._orientation.Next
	} else {
		S._orientation = S._orientation.Prev
	}
}

//---------------------------------------------------------------------------------------//

func init() {
	// Инициализация фигуры O
	_O1.Next = &_O1
	_O1.Prev = &_O1

	// Инициализация фигуры I
	_I1.Next = &_I2
	_I2.Next = &_I1

	_I1.Prev = &_I2
	_I2.Prev = &_I1

	// Инициализация фигуры S
	_S1.Next = &_S2
	_S2.Next = &_S1

	_S1.Prev = &_S2
	_S2.Prev = &_S1

	// Инициализация фигуры Z
	_Z1.Next = &_Z2
	_Z2.Next = &_Z1

	_Z1.Prev = &_Z2
	_Z2.Prev = &_Z1

	// Инициализация фигуры L
	_L1.Next = &_L2
	_L2.Next = &_L3
	_L3.Next = &_L4
	_L4.Next = &_L1

	_L1.Prev = &_L4
	_L2.Prev = &_L1
	_L3.Prev = &_L2
	_L4.Prev = &_L3

	// Инициализация фигуры J
	_J1.Next = &_J2
	_J2.Next = &_J3
	_J3.Next = &_J4
	_J4.Next = &_J1

	_J1.Prev = &_J4
	_J2.Prev = &_J1
	_J3.Prev = &_J2
	_J4.Prev = &_J3

	// Инициализация фигуры T
	_T1.Next = &_T2
	_T2.Next = &_T3
	_T3.Next = &_T4
	_T4.Next = &_T1

	_T1.Prev = &_T4
	_T2.Prev = &_T1
	_T3.Prev = &_T2
	_T4.Prev = &_T3

	// Начальная позиция фигуры в linkedList
	_shapesOrientation[SHAPE_0] = &_O1
	_shapesOrientation[SHAPE_I] = &_I1
	_shapesOrientation[SHAPE_L] = &_L1
	_shapesOrientation[SHAPE_J] = &_J1
	_shapesOrientation[SHAPE_S] = &_S1
	_shapesOrientation[SHAPE_Z] = &_Z1
	_shapesOrientation[SHAPE_T] = &_T1
}

//---------------------------------------------------------------------------------------//

var _O1 _Orientation = _Orientation{
	Bricks: []Brick{
		{X: 0, Y: 0}, {X: 1, Y: 0},
		{X: 0, Y: 1}, {X: 1, Y: 1},
	},
}

//---------------------------------------------------------------------------------------//

var _I1 _Orientation = _Orientation{
	Bricks: []Brick{
		{X: -1, Y: 0}, {X: 2, Y: 0},
		{X: 1, Y: 0}, {X: 0, Y: 0},
	},
}

var _I2 _Orientation = _Orientation{
	Bricks: []Brick{
		{X: 1, Y: 0}, {X: 1, Y: -1},
		{X: 1, Y: 1}, {X: 1, Y: 2},
	},
}

//---------------------------------------------------------------------------------------//

var _S1 _Orientation = _Orientation{
	Bricks: []Brick{
		{X: 2, Y: 0}, {X: 1, Y: 0},
		{X: 1, Y: 1}, {X: 0, Y: 1},
	},
}
var _S2 _Orientation = _Orientation{
	Bricks: []Brick{
		{X: 1, Y: -1}, {X: 2, Y: 0},
		{X: 1, Y: 0}, {X: 2, Y: 1},
	},
}

//---------------------------------------------------------------------------------------//

var _Z1 _Orientation = _Orientation{
	Bricks: []Brick{
		{X: 0, Y: 0}, {X: 1, Y: 0},
		{X: 1, Y: 1}, {X: 2, Y: 1},
	},
}

var _Z2 _Orientation = _Orientation{
	Bricks: []Brick{
		{X: 1, Y: 0}, {X: 2, Y: 0},
		{X: 1, Y: 1}, {X: 2, Y: -1},
	},
}

//---------------------------------------------------------------------------------------//

var _L1 _Orientation = _Orientation{
	Bricks: []Brick{
		{X: 0, Y: 1}, {X: 0, Y: 0},
		{X: 2, Y: 0}, {X: 1, Y: 0},
	},
}

var _L2 _Orientation = _Orientation{
	Bricks: []Brick{
		{X: 1, Y: -1}, {X: 2, Y: 1},
		{X: 1, Y: 1}, {X: 1, Y: 0},
	},
}

var _L3 _Orientation = _Orientation{
	Bricks: []Brick{
		{X: 2, Y: -1}, {X: 0, Y: 0},
		{X: 2, Y: 0}, {X: 1, Y: 0},
	},
}

var _L4 _Orientation = _Orientation{
	Bricks: []Brick{
		{X: 1, Y: -1}, {X: 0, Y: -1},
		{X: 1, Y: 1}, {X: 1, Y: 0},
	},
}

//---------------------------------------------------------------------------------------//

var _J1 _Orientation = _Orientation{
	Bricks: []Brick{
		{X: 2, Y: 1}, {X: 0, Y: 0},
		{X: 2, Y: 0}, {X: 1, Y: 0},
	},
}

var _J2 _Orientation = _Orientation{
	Bricks: []Brick{
		{X: 1, Y: -1}, {X: 2, Y: -1},
		{X: 1, Y: 1}, {X: 1, Y: 0},
	},
}

var _J3 _Orientation = _Orientation{
	Bricks: []Brick{
		{X: 0, Y: -1}, {X: 0, Y: 0},
		{X: 2, Y: 0}, {X: 1, Y: 0},
	},
}

var _J4 _Orientation = _Orientation{
	Bricks: []Brick{
		{X: 1, Y: -1}, {X: 0, Y: 1},
		{X: 1, Y: 1}, {X: 1, Y: 0},
	},
}

//---------------------------------------------------------------------------------------//

var _T1 _Orientation = _Orientation{
	Bricks: []Brick{
		{X: 0, Y: -1 + 1}, {X: 1, Y: -1 + 2},
		{X: 2, Y: -1 + 1}, {X: 1, Y: -1 + 1},
	},
}

var _T2 _Orientation = _Orientation{
	Bricks: []Brick{
		{X: 1, Y: -1 + 0}, {X: 2, Y: -1 + 1},
		{X: 1, Y: -1 + 1}, {X: 1, Y: -1 + 2},
	},
}

var _T3 _Orientation = _Orientation{
	Bricks: []Brick{
		{X: 0, Y: -1 + 1}, {X: 1, Y: -1 + 0},
		{X: 2, Y: -1 + 1}, {X: 1, Y: -1 + 1},
	},
}

var _T4 _Orientation = _Orientation{
	Bricks: []Brick{
		{X: 0, Y: -1 + 1}, {X: 1, Y: -1 + 2},
		{X: 1, Y: -1 + 0}, {X: 1, Y: -1 + 1},
	},
}
