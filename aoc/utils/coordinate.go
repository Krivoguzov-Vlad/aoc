package utils

var (
	None  = Coordinate{}
	Up    = Coordinate{Y: -1}
	Down  = Coordinate{Y: 1}
	Left  = Coordinate{X: -1}
	Right = Coordinate{X: 1}
)

type Coordinate struct {
	X, Y int
}

func (c Coordinate) ManhattanDistance() int {
	return Abs(c.X) + Abs(c.Y)
}

func (c Coordinate) ToDirection() Coordinate {
	if c.X != 0 {
		c.X /= Abs(c.X)
	}
	if c.Y != 0 {
		c.Y /= Abs(c.Y)
	}
	return c
}

func (c Coordinate) Add(o Coordinate) Coordinate {
	c.X += o.X
	c.Y += o.Y
	return c
}

func (c Coordinate) Sub(o Coordinate) Coordinate {
	c.X -= o.X
	c.Y -= o.Y
	return c
}

func (c Coordinate) Neighbours() []Coordinate {
	return []Coordinate{
		c.Add(Up),
		c.Add(Down),
		c.Add(Left),
		c.Add(Right),
	}
}

func AllDirections() [4]Coordinate {
	return [4]Coordinate{Up, Down, Left, Right}
}
