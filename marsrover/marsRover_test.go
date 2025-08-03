package marsrover

import "testing"

func TestNewPlateau(t *testing.T) {
	x, y := 5, 5
	plateau, _ := NewPlateau(x, y)

	if plateau.X != 5 || plateau.Y != 5 {
		t.Errorf("Expected (5, 5), got (%d, %d)", plateau.X, plateau.Y)
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

func TestCreateRover(t *testing.T) {
	plateau := Plateau{X: 5, Y: 5}

	rover, _ := NewRover(1, 2, North, plateau)

	if rover.x != 1 || rover.y != 2 {
		t.Errorf("Expected position (1, 2), got (%d, %d)", rover.x, rover.y)
	}
	if rover.dir != North {
		t.Errorf("Expected direction North, got %s", rover.dir.String())
	}
}

func TestRoverCannotHavexLessThan0(t *testing.T) {
	plateau := Plateau{X: 5, Y: 5}

	_, err := NewRover(-1, 2, North, plateau)

	if err == nil {
		t.Errorf("Rover Cannot have x<0")
	}

}

func TestRoverCannotHaveyLessThan0(t *testing.T) {
	plateau := Plateau{X: 5, Y: 5}

	_, err := NewRover(1, -2, North, plateau)

	if err == nil {
		t.Errorf("Rover Cannot have y<0")
	}

}

func TestRoverTurnsLeft(t *testing.T) {
	plateau, _ := NewPlateau(5, 5)
	rover, _ := NewRover(1, 2, North, *plateau)

	rover.TurnLeft()
	if rover.dir != West {
		t.Errorf("Expected direction West after left turn, got %s", rover.dir)
	}

}

func TestRoverTurnsRight(t *testing.T) {
	plateau, _ := NewPlateau(5, 5)
	rover, _ := NewRover(1, 2, North, *plateau)

	rover.TurnRight()
	if rover.dir != East {
		t.Errorf("Expected direction West after left turn, got %s", rover.dir)
	}

}

func TestRoverMove(t *testing.T) {
	plateau, _ := NewPlateau(5, 5)
	rover, _ := NewRover(1, 2, North, *plateau)

	rover.Move()

	if rover.x != 1 || rover.y != 3 {
		t.Errorf("Expected position (1, 3), got (%d, %d)", rover.x, rover.y)
	}
}

func TestExecute(t *testing.T) {
	plateau, _ := NewPlateau(5, 5)
	rover, _ := NewRover(1, 2, North, *plateau)

	err := rover.Execute("LMLMLMLMM")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if rover.x != 1 || rover.y != 3 || rover.dir != North {
		t.Errorf("Expected position (1, 3, N), got (%d, %d, %s)", rover.x, rover.y, rover.dir)
	}
}
