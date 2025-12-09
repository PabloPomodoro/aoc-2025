package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	total := 0

	data, _ := os.ReadFile("test_data.txt")
	full := strings.Split(string(data), "\n")

	row := full[0]
	width := len(strings.Fields(row))

	for counter := width - 1; counter >= 0; counter-- {

		firstString := strings.Fields(full[0])[counter]
		secondString := strings.Fields(full[1])[counter]
		thirdString := strings.Fields(full[2])[counter]

		max := len(firstString)
		if len(secondString) > max {
			max = len(secondString)
		}
		if len(thirdString) > max {
			max = len(thirdString)
		}

		var sum int
		for k := max - 1; k >= 0; k-- {

			thirdDigit := 0
			secondDigit := 0
			firstDigit := 0

			if k < len(firstString) {
				firstDigit, _ = strconv.Atoi(strings.Split(firstString, "")[k])
			}
			if k < len(secondString) {
				secondDigit, _ = strconv.Atoi(strings.Split(secondString, "")[k])
			}
			if k < len(thirdString) {
				thirdDigit, _ = strconv.Atoi(strings.Split(thirdString, "")[k])
			}

			number := 100*firstDigit + 10*secondDigit + thirdDigit

			fmt.Println(number)

			operator := strings.Fields(full[3])[counter]
			switch operator {
			case "+":
				sum = 0
				sum += number
			case "*":
				sum = 1
				sum *= number
			}
		}

		total += sum
	}

	fmt.Println(total)
}
