package day03

import (
	"fmt"

	"github.com/jenspfeifle/adventofcode-2015/utils"
)

const (
	N = '^'
	S = 'v'
	E = '>'
	W = '<'
)

type Coord struct {
	x int
	y int
}

func Main() error {
	fmt.Println("Day 03:")
	lines, err := utils.ReadInputLines("./inputs/03.in")
	if err != nil {
		return err
	}
	// single line input
	directions := lines[0]
	part1(directions)
	part2(directions)
	return nil
}

func part1(directions string) {
	houses := map[Coord]int{}
	pos := Coord{0, 0}
	houses[pos] = 1
	for _, direction := range directions {
		pos = nextPos(pos, direction)
		houses[pos] += 1
	}
	total := countHousesWithPresents(houses)
	fmt.Println("Part1: Houses with at least one present:", total)
}

func part2(directions string) {
	houses := map[Coord]int{}
	posSanta := Coord{0, 0}
	posRoboSanta := Coord{0, 0}
	houses[posSanta] = 2

	for n, direction := range directions {
		if n%2 == 0 {
			// Robo-Santa's turn
			posRoboSanta = nextPos(posRoboSanta, direction)
			houses[posRoboSanta] += 1
		} else {
			// Santa's turn
			posSanta = nextPos(posSanta, direction)
			houses[posSanta] += 1
		}
	}
	total := countHousesWithPresents(houses)
	fmt.Println("Part2: Houses with at least one present:", total)
}

// count houses with at least one present
func countHousesWithPresents(houses map[Coord]int) int {
	total := 0
	for _, v := range houses {
		if v > 0 {
			total += 1
		}
	}
	return total
}

// given a position and direction, return new position
func nextPos(pos Coord, direction rune) Coord {
	x := pos.x
	y := pos.y
	switch direction {
	case N:
		y += 1
	case S:
		y -= 1
	case E:
		x += 1
	case W:
		x -= 1
	default:
		// FIXME: err?
	}
	return Coord{x, y}
}
