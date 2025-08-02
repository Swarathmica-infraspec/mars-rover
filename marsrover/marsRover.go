package marsrover

import "errors"

type Plateau struct {
	x, y int
}

type Rover struct {
	x, y    int
	dir     Direction
	plateau Plateau
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
		return nil, errors.New("Cannot create plateau with x<1")
	}
	return &Plateau{x: x, y: y}, nil
}

func NewRover(x, y int, dir Direction, plateau Plateau) *Rover {

	return &Rover{x: x, y: y, dir: dir, plateau: plateau}
}
