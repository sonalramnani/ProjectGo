package main

import (
	"fmt"
	"errors"
)

var dirMap map[string]int = map[string]int{"N": 0, "E": 1, "S": 2, "W": 3}

type roverCoords struct {
	x         int
	y         int
	direction int
	moves     string
}

func NewRover(inX int, inY int, dir string) (*roverCoords, error) {
	d, ok := dirMap[dir]
	if !ok {
		return nil, fmt.Errorf("invalid direction\n")
	}

	r := &roverCoords{
		x:         inX,
		y:         inY,
		direction: d,
	}
	return r, nil
}

func (r *roverCoords) setMoves(m string) {
	r.moves = m
}

func (r *roverCoords) getDirection() (string, error) {
        for k, v := range dirMap {
                if v == r.direction {
                        return k, nil
                }
        }
        return "0", errors.New("invalid direction")
}

func (r *roverCoords) move(movement byte, upperX, upperY int) error {
	switch movement {
	case 'R':
		if r.direction == 3 {
			r.direction = 0
		} else {
			r.direction++
		}

	case 'L':
		if r.direction == 0 {
			r.direction = 3
		} else {
			r.direction--
		}

	case 'M':
		switch r.direction {
                case 0:
			//direction = North
                        if r.y >= upperY {
                                return fmt.Errorf("exceeding upper Y boundary")
                        }
                        r.y++
		case 1:
			//direction = East
			if r.x >= upperX {
				return fmt.Errorf("exceeding upper X boundary")
			}
			r.x++
		case 2:
			//direction = South
			if r.y <= lowerY {
				return fmt.Errorf("moving beyond lowerY on y-axis")
			}
			r.y--
		case 3:
			//direction = West
			if r.x <= lowerX {
				return fmt.Errorf("moving beyond lowerX on x-axis")
			}
			r.x--
		default:
			return fmt.Errorf("invalid direction")
		}
	default:
		return fmt.Errorf("invalid input")
	}
	return nil
}

func (r *roverCoords) executeCommand(upperX, upperY int) error {
		var dir string
		var err error 

                moves := r.moves
                for i := 0; i < len(moves); i++ {
                        if err := r.move(moves[i], upperX, upperY); err != nil {
                        	return err
			}
                }

                if dir, err = r.getDirection(); err != nil {
			return err
		}

		fmt.Printf("After move: x:%d, y: %d, dir: %s\n", r.x, r.y, dir)
		return nil
}
