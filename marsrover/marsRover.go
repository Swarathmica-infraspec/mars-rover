package marsrover

import "errors"

type Plateau struct {
	X, Y int
}

type Rover struct {
	X, Y    int
	Dir     Direction
	Plateau Plateau
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

var directions = [...]string{"N", "E", "S", "W"}

func (d Direction) String() string {
	switch d {
	case North:
		return "N"
	case East:
		return "E"
	case South:
		return "S"
	case West:
		return "W"
	default:
		return "?"
	}
}

func NewPlateau(x, y int) (*Plateau, error) {
	if x < 1 || y < 1 {
		return nil, errors.New("Cannot create plateau with x<1 or y<1")
	}
	return &Plateau{X: x, Y: y}, nil
}

func NewRover(x, y int, dir Direction, plateau Plateau) (*Rover, error) {
	if x < 0 || y < 0 {
		return nil, errors.New("Rover cannot have starting position with negative coordinates")
	}
	return &Rover{X: x, Y: y, Dir: dir, Plateau: plateau}, nil
}

func (r *Rover) TurnLeft() {
	r.Dir = (r.Dir + 3) % 4
}

func (r *Rover) TurnRight() {
	r.Dir = (r.Dir + 1) % 4
}

func (r *Rover) Move() {
	newX, newY := r.X, r.Y

	switch r.Dir {
	case North:
		newY++
	case East:
		newX++
	case South:
		newY--
	case West:
		newX--
	}

	if newX >= 0 && newY >= 0 && newX <= r.Plateau.X && newY <= r.Plateau.Y {
		r.X, r.Y = newX, newY
	}
}

func (r *Rover) Execute(commands string) error {
	for _, c := range commands {
		switch c {
		case 'L':
			r.TurnLeft()
		case 'R':
			r.TurnRight()
		case 'M':
			r.Move()
		default:
			return errors.New("invalid command")
		}
	}
	return nil
}
