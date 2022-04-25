package day02

import (
	"errors"
	"fmt"
	"github.com/jenspfeifle/adventofcode-2015/utils"
	"os"
	"strconv"
	"strings"
)

func Main() error {
	fmt.Println("Day 02:")
	lines, err := utils.ReadInputLines("./inputs/02.in")
	if err != nil {
		fmt.Fprintln(os.Stderr, "reading file:", err)
	}

	var totalPaperArea = 0
	var totalRibbonLength = 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		d, err := splitLineIntoDimensions(line)
		if err != nil {
			return err
		}

		totalPaperArea += wrappingPaperRequired(d)
		totalRibbonLength += ribbonLengthRequired(d)
	}
	fmt.Println("Part1: Total paper area =", totalPaperArea)
	fmt.Println("Part2: Total riboon length =", totalRibbonLength)
	return nil
}

type dimensions struct {
	l int
	w int
	h int
}

func splitLineIntoDimensions(line string) (dimensions, error) {
	if len(line) == 0 {
		return dimensions{}, errors.New("Given line is empty")
	}
	var segments []string = strings.Split(line, "x")
	numbers := make([]int, len(segments))
	if len(segments) != 3 {
		return dimensions{}, errors.New("Wrong number of dimensions: " + line)
	}
	for i, s := range segments {
		n, err := strconv.Atoi(s)
		if err != nil {
			return dimensions{}, err
		}
		numbers[i] = n
	}

	return dimensions{numbers[0], numbers[1], numbers[2]}, nil

}
func ribbonLengthRequired(d dimensions) int {
	return smallestPerimeter(d) + totalVolume(d)
}

func wrappingPaperRequired(d dimensions) int {
	return totalSurfaceArea(d) + areaOfSmallestSide(d)
}

func smallestPerimeter(d dimensions) int {
	perimeters := []int{2 * (d.l + d.w), 2 * (d.w + d.h), 2 * (d.h + d.l)}
	min, _ := utils.MinMax(perimeters)
	return min
}

func totalVolume(d dimensions) int {
	return d.l * d.w * d.h
}

func areaOfSmallestSide(d dimensions) int {
	sides := []int{d.l * d.w, d.w * d.h, d.h * d.l}
	min, _ := utils.MinMax(sides)
	return min
}

func totalSurfaceArea(d dimensions) int {
	return 2*d.l*d.w + 2*d.w*d.h + 2*d.h*d.l
}
