package marsrover

import "testing"

func TestNewPlateau(t *testing.T) {
	x, y := 5, 5
	plateau := NewPlateau(x, y)

	if plateau.x != 5 || plateau.y != 5 {
		t.Errorf("Expected (5, 5), got (%d, %d)", plateau.x, plateau.y)
	}
}
