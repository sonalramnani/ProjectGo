package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


const (
	lowerX = 0
	lowerY = 0
)


func readLn(r *bufio.Reader) (string, error) {
	var line []byte
	var err error

	if line, _, err = r.ReadLine(); err != nil {
		return "", err
	}
	return string(line), nil
}


func readInput(reader *bufio.Reader) (int, int, string, string, error) {
 	var startX, startY int
        var err error

        line, _ := readLn(reader)
        s := strings.Split(line, " ")

        if startX, err = strconv.Atoi(s[0]); err != nil {
                return 0, 0, "", "", fmt.Errorf("error in converting x coords\n")
        }
        if startY, err = strconv.Atoi(s[1]); err != nil {
                return 0, 0, "", "",  fmt.Errorf("error in converting y coords\n")
        }

        ln, _ := readLn(reader)
        moves := string(ln)
	return startX, startY,s[2], moves, nil
}


func main() {
	var upperX, upperY, startX, startY int
	var err error
	var dir, moves string
	var r *roverCoords
	var roverArr []*roverCoords = make([]*roverCoords, 2)
	var index = 0

	reader := bufio.NewReader(os.Stdin)
	line, _ := readLn(reader)
	s := strings.Split(line, " ")
	upperX, _ = strconv.Atoi(s[0])
	upperY, _ = strconv.Atoi(s[1])

	if startX, startY, dir, moves, err = readInput(reader); err != nil {
 		fmt.Printf("Invalid input: %v", err)
		os.Exit(1)
	}

	if r, err = NewRover(startX, startY, dir); err != nil {
		fmt.Printf("Invalid input: %v", err)
		os.Exit(1)
	}
	r.setMoves(moves)
	roverArr[index] = r
	index++

	if startX, startY, dir, moves, err = readInput(reader); err != nil {
		fmt.Printf("Invalid input: %v", err)
                os.Exit(1)
        }

        if r, err = NewRover(startX, startY, dir); err != nil {
		fmt.Printf("Invalid input: %v", err)
                os.Exit(1)
	}
        r.setMoves(moves)
	roverArr[index] = r

	for i := 0; i < 2; i++ {	
		roverArr[i].executeCommand(upperX, upperY)
	}
}	
