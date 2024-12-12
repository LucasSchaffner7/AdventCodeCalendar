package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func read_input(file_name string) [][]byte {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines [][]byte

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, append([]byte(nil), scanner.Bytes()...))
	}

	return lines
}

func detect_XMAS(lines [][]byte) int {
	directions := [][2]int{
		{1, 0},   // EAST
		{1, -1},  // SOUTH-EAST
		{0, -1},  // SOUTH
		{-1, -1}, // SOUTH-WEST
		{-1, 0},  // WEST
		{-1, 1},  // NORTH-WEST
		{0, 1},   // NORTH
		{1, 1},   // NORTH-EAST
	}

	nb_XMAS := 0

	for y, line := range lines {
		for x, character := range line {
			if character != 'X' {
				continue
			}

			for _, dir := range directions {
				nb_XMAS += check_XMAS(lines, x, y, dir)
			}
		}
	}

	return nb_XMAS
}

func check_XMAS(lines [][]byte, x int, y int, dir [2]int) int {
	dx, dy := dir[0], dir[1]

	max_x, max_y := x+dx*3, y+dy*3
	if verify_bounds(max_x, max_y, len(lines[0]), len(lines)) {
		return 0
	}

	for i := 1; i < 4; i++ {
		nx, ny := x+dx*i, y+dy*i

		switch i {
		case 1:
			if lines[ny][nx] != 'M' {
				return 0
			}
		case 2:
			if lines[ny][nx] != 'A' {
				return 0
			}
		case 3:
			if lines[ny][nx] != 'S' {
				return 0
			}
		}
	}

	return 1
}

func detect_X_MAS(lines [][]byte) int {
	directions := [][2]int{
		{1, 1},   // NORTH-EAST
		{1, -1},  // SOUTH-EAST
		{-1, -1}, // SOUTH-WEST
		{-1, 1},  // NORTH-WEST
	}

	nb_X_MAS := 0

	for y, line := range lines {
		for x, character := range line {
			if character != 'A' {
				continue
			}

			nb_X_MAS += check_X_MAS(lines, x, y, directions)
		}
	}

	return nb_X_MAS
}

func check_X_MAS(lines [][]byte, x int, y int, directions [][2]int) int {
	for i := 0; i < 2; i++ {
		dx1, dy1 := directions[i][0], directions[i][1]
		dx2, dy2 := directions[i+2][0], directions[i+2][1]

		max_x1, max_y1 := x+dx1, y+dy1
		max_x2, max_y2 := x+dx2, y+dy2
		if verify_bounds(max_x1, max_y1, len(lines[0]), len(lines)) || verify_bounds(max_x2, max_y2, len(lines[0]), len(lines)) {
			return 0
		}

		if !((lines[y+dy1][x+dx1] == 'M' && lines[y+dy2][x+dx2] == 'S') || (lines[y+dy1][x+dx1] == 'S' && lines[y+dy2][x+dx2] == 'M')) {
			return 0
		}
	}

	return 1
}

func verify_bounds(x int, y int, dim_x int, dim_y int) bool {
	return x < 0 || x >= dim_x || y < 0 || y >= dim_y
}

func main() {
	fmt.Println("Merry day 4 in code calendar !")
	lines := read_input("input.txt")
	fmt.Printf("The number of XMAS is : %d\n", detect_XMAS(lines))
	fmt.Printf("The number of X-MAS is : %d\n", detect_X_MAS(lines))
}
