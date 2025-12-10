package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	total := 0

	data, _ := os.ReadFile("data.txt")
	full := strings.Split(string(data), "\n")

	lastRow := full[len(full)-1]
	for index, operator := range lastRow {
		if operator == ' ' {
			continue
		}

		firstDigitIndex := 0
		last := true
		for next := index + 1; next < len(lastRow); next++ {
			last = true
			if lastRow[next] != ' ' {
				firstDigitIndex = next - 2
				last = false
				break
			}
		}

		if last {
			width := 0
			for _, line := range full {
				if len(line) > width {
					width = len(line)
				}
			}
			firstDigitIndex = width - 1
		}

		sum := 0
		for corrective := 0; corrective < len(full)-1; corrective++ {
			builder := []rune{}
			for i, line := range full {
				if i == len(full)-1 {
					break
				}
				if (firstDigitIndex - corrective) > -1 {
					digit, _ := strconv.Atoi(string(line[firstDigitIndex-corrective]))
					if digit != 0 {
						builder = append(builder, rune('0'+digit))
					}
				}
			}
			number, _ := strconv.Atoi(string(builder))
			fmt.Println(number)

			switch operator {
			case '+':
				sum += number
			case '*':
				if sum == 0 {
					sum = 1
				}
				sum *= number
			}
		}
		fmt.Println(sum)
		fmt.Println()

		total += sum
	}

	fmt.Println(total)
}
