// +build unitTest

package robot

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRobot(t *testing.T) {
	t.Run("RobotHappyPath", TestStartRobot_HappyPath)
	t.Run("MoveRobotToEarlierPath", TestMoveRobotToEarlierPath)
	t.Run("MoveRobotOutOfPlane", TestMoveRobotOutOfPlane)
	t.Run("ChangeDirection", TestChangeDirection)
}

func TestStartRobot_HappyPath(t *testing.T) {
	robot := NewRobot(4, 4, 0, 0, North)
	coordinate1, coordinate2, direction, err := robot.StartRobot("MMMRMMLM")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 2, coordinate1, "Robot coordinate1 is not correct")
	assert.Equal(t, 4, coordinate2, "Robot coordinate2 is not correct")
	assert.Equal(t, North, direction, "Robot direction is not correct")
}

func TestMoveRobotToEarlierPath(t *testing.T) {
	robot := NewRobot(4, 4, 0, 0, North)
	_, _, _, err := robot.StartRobot("MMMRRM")
	assert.Errorf(t, err, "Not returning error")
	{
		assert.Equal(t, "robot is trying to walk where it already travelled", err.Error())
	}
}

func TestChangeDirection(t *testing.T) {
	robot := NewRobot(4, 4, 0, 0, North)
	robot.direction = West
	robot.changeDirection(Right)
	assert.Equal(t, North, robot.direction, "Robot direction is not correct")
}

func TestMoveRobotOutOfPlane(t *testing.T) {
	robot := NewRobot(4, 4, 0, 0, South)
	err := robot.moveRobot()
	assert.Errorf(t, err, "Not returning error")
	{
		assert.Equal(t, "robot is trying to walk away from the rectangular plane", err.Error())
	}
	robot.xCoordinate = 4
	robot.direction = North
	err = robot.moveRobot()
	assert.Errorf(t, err, "Not returning error")
	{
		assert.Equal(t, "robot is trying to walk away from the rectangular plane", err.Error())
	}
}
