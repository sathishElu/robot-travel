package robot

import (
	"errors"
	"log"
	"strings"
)

type Direction string
type Command string

const (
	North Direction = "N"
	South           = "S"
	East            = "E"
	West            = "W"
)

const (
	Move  Command = "M"
	Left          = "L"
	Right         = "R"
)

const mark = "*"

type Robot struct {
	plane       [][]string
	direction   Direction
	xCoordinate int
	yCoordinate int
}

//NewRobot builds a Robot with valid config
func NewRobot(hSize, vSize, xCoordinate, yCoordinate int, direction Direction) Robot {
	//Align with the requirement. Increasing by 1.

	if len(direction) != 1 || strings.ContainsAny(string(direction), "NSEW") == false {
		log.Fatal("robot's starting direction is not valid")
	}

	plane := make([][]string, hSize)
	for i := range plane {
		plane[i] = make([]string, vSize)
	}

	if xCoordinate >= len(plane) || yCoordinate >= len(plane[xCoordinate]) {
		log.Fatal("robot's starting position is not within the rectangular plane")
	}

	robot := Robot{
		plane:       plane,
		direction:   direction,
		xCoordinate: xCoordinate,
		yCoordinate: yCoordinate,
	}
	return robot
}

//StartRobot starts the robot with the commands. Robot would stop, if any command is not valid.
func (c *Robot) StartRobot(commands string) (int, int, Direction, error) {
	cmd := []rune(commands)
	for i := 0; i < len(cmd); i++ {
		nextStep := Command(cmd[i])
		if nextStep == Right || nextStep == Left {
			c.changeDirection(nextStep)
			continue
		} else if nextStep == Move {
			isMoved := c.moveRobot()
			if isMoved == false {
				return c.xCoordinate, c.yCoordinate, c.direction, nil
			}
		} else {
			return 0, 0, North, errors.New("robot's command is not valid")
		}
	}
	return c.xCoordinate, c.yCoordinate, c.direction, nil
}

func (c *Robot) changeDirection(command Command) {
	if command == Right {
		switch {
		case c.direction == North:
			c.direction = East
		case c.direction == South:
			c.direction = West
		case c.direction == East:
			c.direction = South
		case c.direction == West:
			c.direction = North
		}
	} else if command == Left {
		switch {
		case c.direction == North:
			c.direction = West
		case c.direction == South:
			c.direction = East
		case c.direction == East:
			c.direction = North
		case c.direction == West:
			c.direction = South
		}
	}
}

func (c *Robot) moveRobot() bool {
	xAxis, yAxis := c.xCoordinate, c.yCoordinate
	switch {
	case c.direction == North:
		xAxis = c.xCoordinate + 1
	case c.direction == South:
		xAxis = c.xCoordinate - 1
	case c.direction == East:
		yAxis = c.yCoordinate + 1
	case c.direction == West:
		yAxis = c.yCoordinate - 1
	}

	if xAxis >= len(c.plane[c.xCoordinate]) || xAxis < 0 || yAxis >= len(c.plane) || yAxis < 0 {
		return false
	} else if c.plane[xAxis][yAxis] == mark {
		return false
	} else if c.plane[xAxis][yAxis] != "" {
		return false
	} else {
		c.xCoordinate = xAxis
		c.yCoordinate = yAxis
		c.plane[c.xCoordinate][c.yCoordinate] = mark
	}
	return true
}
