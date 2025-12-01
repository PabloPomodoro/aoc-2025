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

		if clicks > 100 {
			clicks = clicks % 100
		}

		switch direction {
		case "R":
			dialPoint = (dialPoint + clicks) % 100
		case "L":
			dialPoint = (dialPoint - clicks) % 100
		default:
			panic(-1)
		}

		if dialPoint == 0 {
			zeroCounter++
		}
	}

	fmt.Println(zeroCounter)
}
