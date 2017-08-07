package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMove1(t *testing.T) {
	fakeRover := roverCoords {
		x: 0,
		y: 0,
		direction: 0,
	}

	if err := fakeRover.move('L', 5, 5); err != nil {
		t.Errorf("failed: %v\n", err)
	}
	assert.Equal(t, fakeRover.direction, 3, "Direction should be W")
}

func TestMove2(t *testing.T) {
        fakeRover := roverCoords {
                x: 0,
                y: 0,
                direction: 0,
        }

        if err := fakeRover.move('N', 5, 5); err != nil {
                assert.EqualErrorf(t,  err, "invalid input", "failed: %v\n", err)
        }
}

func TestMove3(t *testing.T) {
        fakeRover := roverCoords {
                x: 1,
                y: 1,
                direction: 0,
        }       
        
        if err := fakeRover.move('M', 0, 0); err != nil {
                assert.EqualErrorf(t, err, "exceeding upper Y boundary", "failed: %v\n", err)
        }
}

func TestMove4(t *testing.T) {
        fakeRover := roverCoords {
                x: 0,
                y: 1,
                direction: 3,
        }

        if err := fakeRover.move('M', 5, 5); err != nil {
                assert.EqualErrorf(t, err, "moving beyond lowerX on x-axis", "failed: %v\n", err)
        }
}

func TestMove5(t *testing.T) {
        fakeRover := roverCoords {
                x: 1,
                y: 2,
                direction: 0,
        }

        if err := fakeRover.move('M', 5, 5); err != nil {
                t.Errorf("failed: %v\n", err)
        }
	assert.Equal(t, fakeRover.y,  3, "y should be 3")
}

func TestExecuteCommand1(t *testing.T) {
	fakeRover := roverCoords {
		x: 1,
		y: 2,
		direction: 0,
	}
	fakeRover.setMoves("LMLMLMLMM")

	if err := fakeRover.executeCommand(5, 5); err != nil {
                t.Errorf("failed: %v\n", err)
        }
        assert.Equal(t, fakeRover.x, 1, "x should be 1")
	assert.Equal(t, fakeRover.y, 3, "y should be 3")
	assert.Equal(t, fakeRover.direction, 0, "direction should be N")
}


func TestExecuteCommand2(t *testing.T) {
        fakeRover := roverCoords {
                x: 3,
                y: 3,
                direction: 1,
        }
        fakeRover.setMoves("MMRMMRMRRM")

        if err := fakeRover.executeCommand(5, 5); err != nil {
                t.Errorf("failed: %v\n", err)
        }
        assert.Equal(t, fakeRover.x, 5, "x should be 1")
        assert.Equal(t, fakeRover.y, 1, "y should be 3")
        assert.Equal(t, fakeRover.direction, 1, "direction should be E")
}

