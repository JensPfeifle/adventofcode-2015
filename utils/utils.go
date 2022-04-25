package utils

import (
	"bufio"
	"os"
)

func ReadInputLines(fileName string) ([]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(f)
	out := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		out = append(out, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return out, nil
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
