package marsrover

import "testing"

func TestNewPlateau(t *testing.T) {
	x, y := 5, 5
	plateau, _ := NewPlateau(x, y)

	if plateau.x != 5 || plateau.y != 5 {
		t.Errorf("Expected (5, 5), got (%d, %d)", plateau.x, plateau.y)
	}
}

func TestCannotCreatePlateauWithxLessThan1(t *testing.T) {
	x, y := -1, 5
	_, err := NewPlateau(x, y)
	if err == nil {
		t.Errorf("Cannot create Plateau With negative x")
	}

}
