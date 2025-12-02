package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	dialPoint := 50
	zeroCounter := 0

	data, _ := os.ReadFile("data.txt")
	lines := strings.SplitSeq(string(data), "\n")

	for line := range lines {
		if line == "" {
			continue
		}

		direction := line[:1]
		clicks, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(-1)
		}

		switch direction {
		case "R":
			if dialPoint == 0 {
				dialPoint++
				clicks--
			}
			for range clicks {
				dialPoint++
				if dialPoint == 100 {
					dialPoint = 0
					zeroCounter++
				}
			}
		case "L":
			if dialPoint == 0 {
				dialPoint = 99
				clicks--
			}

			wrap := false
			for range clicks {
				if wrap {
					dialPoint = 100
					wrap = false
				}
				dialPoint--
				if dialPoint == 0 {
					wrap = true
					zeroCounter++
				}
			}
		default:
			panic(-1)
		}
	}

	fmt.Println(zeroCounter)
}
