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

func TestCannotCreatePlateauWithyLessThan1(t *testing.T) {
	x, y := 5, -5
	_, err := NewPlateau(x, y)
	if err == nil {
		t.Errorf("Cannot create Plateau With negative y")
	}

}

func TestDirectionString(t *testing.T) {
	if North.String() != "N" {
		t.Errorf("Expected 'N', got '%s'", North.String())
	}
	if West.String() != "W" {
		t.Errorf("Expected 'W', got '%s'", West.String())
	}
}
