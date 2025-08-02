package marsrover

import "errors"

type Plateau struct {
	x, y int
}

func NewPlateau(x, y int) (*Plateau, error) {
	if x < 1 || y < 1 {
		return nil, errors.New("Cannot create plateau with x<1")
	}
	return &Plateau{x: x, y: y}, nil
}
