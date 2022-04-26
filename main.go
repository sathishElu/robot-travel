package main

import (
	"fmt"
	"robot-travel/robot"
	"log"
)

func main() {

	var hSize, vSize, xCoordinate, yCoordinate int
	var position, commands string

	fmt.Println("Test Input")
	_, err := fmt.Scanf("%d%d\n", &hSize, &vSize)
	if err != nil {
		log.Fatal("coordinates of the rectangular plane is not valid")
	}

	_, err = fmt.Scanf("%d%d%s\n", &xCoordinate, &yCoordinate, &position)
	if err != nil {
		log.Fatal("robot's starting position is not valid")
	}

	_, err = fmt.Scanf("%s\n", &commands)
	if err != nil {
		log.Fatal("robot's command is not valid")
	}

	newRobot := robot.NewRobot(hSize, vSize, xCoordinate, yCoordinate, robot.Direction(position))
	coordinate1, coordinate2, direction, err := newRobot.StartRobot(commands)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Test Output")
	fmt.Println(coordinate1, coordinate2, direction)
}
