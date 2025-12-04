package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	total := 0

	data, _ := os.ReadFile("data.txt")
	banks := strings.SplitSeq(string(data), "\n")

	for bank := range banks {

		slotsToFill := 12
		battery := 0
		index := 0
		maxIndex := 0
		maxDigit := 10

		for {

			battery = 0
			for i := len(bank) - 1; i >= maxIndex; i-- {
				digit, _ := strconv.Atoi(string(bank[i]))

				if digit >= maxDigit {
					continue
				}

				if digit >= battery {
					battery = digit
					index = i
				}
			}

			if slotsToFill > (len(bank) - index) {
				maxDigit = battery
				continue
			}

			slotsToFill--
			maxIndex = index + 1
			maxDigit = 10

			total += (battery * int(math.Pow(float64(10), float64(slotsToFill))))

			if slotsToFill == 0 {
				break
			}
		}
	}
	fmt.Println(total)
}
