package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	sum := 0

	data, _ := os.ReadFile("data.txt")
	full := strings.Split(string(data), "\n")

	row := full[0]
	width := len(strings.Fields(row))

	for counter := range width {
		operator := strings.Fields(full[4])[counter]

		first, _ := strconv.Atoi(strings.Fields(full[0])[counter])
		second, _ := strconv.Atoi(strings.Fields(full[1])[counter])
		third, _ := strconv.Atoi(strings.Fields(full[2])[counter])
		fourth, _ := strconv.Atoi(strings.Fields(full[3])[counter])

		switch operator {
		case "+":
			sum += first + second + third + fourth
		case "*":
			sum += first * second * third * fourth
		}
	}

	fmt.Println(sum)
}
