package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	rolls := 0

	data, _ := os.ReadFile("data.txt")
	lines := strings.SplitSeq(string(data), "\n")

	grid := [][]byte{}
	for line := range lines {
		grid = append(grid, []byte(line))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			adjacent := 0
			char := grid[i][j]

			if char == '.' {
				continue
			}

			firstline := i == 0
			lastline := i == len(grid)-1

			firstrow := j == 0
			lastrow := j == len(grid[i])-1

			if !firstline {
				if !firstrow {
					if grid[i-1][j-1] == '@' {
						adjacent++
					}
				}
				if grid[i-1][j] == '@' {
					adjacent++
				}
				if !lastrow {
					if grid[i-1][j+1] == '@' {
						adjacent++
					}
				}
			}

			if !firstrow {
				if grid[i][j-1] == '@' {
					adjacent++
				}
			}
			if !lastrow {
				if grid[i][j+1] == '@' {
					adjacent++
				}
			}

			if !lastline {
				if !firstrow {
					if grid[i+1][j-1] == '@' {
						adjacent++
					}
				}
				if grid[i+1][j] == '@' {
					adjacent++
				}
				if !lastrow {
					if grid[i+1][j+1] == '@' {
						adjacent++
					}
				}
			}

			if adjacent < 4 {
				rolls++
			}
		}
	}

	fmt.Println(rolls)
}
