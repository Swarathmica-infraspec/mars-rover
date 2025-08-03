package main

import (
	"fmt"
	"strconv"
	"strings"

	"marsrover/marsrover"
)

func main() {
	input := `5 5
			1 2 N
			LMLMLMLMM
			3 3 E
			MMRMMRMRRM`
	lines := strings.Split(strings.TrimSpace(input), "\n")

	if len(lines) < 1 {
		fmt.Println("Missing plateau input")
		return
	}

	plateauInputs := strings.Fields(lines[0])
	if len(plateauInputs) != 2 {
		fmt.Println("Invalid plateau input")
		return
	}
	maxX, _ := strconv.Atoi(plateauInputs[0])
	maxY, _ := strconv.Atoi(plateauInputs[1])

	plateau, err := marsrover.NewPlateau(maxX, maxY)
	if err != nil {
		fmt.Println("Invalid plateau:", err)
		return
	}

	for i := 1; i < len(lines); i += 2 {
		if i+1 >= len(lines) {
			fmt.Println("Incomplete rover input at line", i+1)
			break
		}

		roverPosition := strings.Fields(strings.TrimSpace(lines[i]))
		if len(roverPosition) != 3 {
			fmt.Println("Invalid rover position:", lines[i])
			continue
		}
		x, _ := strconv.Atoi(roverPosition[0])
		y, _ := strconv.Atoi(roverPosition[1])

		var dir marsrover.Direction
		switch roverPosition[2] {
		case "N":
			dir = marsrover.North
		case "E":
			dir = marsrover.East
		case "S":
			dir = marsrover.South
		case "W":
			dir = marsrover.West
		default:
			fmt.Println("Invalid direction:", roverPosition[2])
			continue
		}

		instructions := strings.TrimSpace(lines[i+1])

		rover, err := marsrover.NewRover(x, y, dir, *plateau)
		if err != nil {
			fmt.Println("Error creating rover:", err)
			continue
		}

		if err := rover.Execute(instructions); err != nil {
			fmt.Println("Error executing commands:", err)
			continue
		}

		fmt.Printf("%d %d %s\n", rover.X, rover.Y, rover.Dir)
	}
}
