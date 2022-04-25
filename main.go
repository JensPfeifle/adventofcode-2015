package main

import (
	"fmt"
	"github.com/jenspfeifle/adventofcode-2015/day02"
	"github.com/jenspfeifle/adventofcode-2015/day03"
	"os"
)

func main() {

	days := []func() error{day02.Main, day03.Main}

	for _, fn := range days {
		err := fn()
		if err != nil {
			fmt.Fprintln(os.Stderr, "reading file:", err)
		}
	}
}
