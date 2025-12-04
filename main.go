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
	banks := strings.SplitSeq(string(data), "\n")

	for bank := range banks {

		firstBattery := 0
		firstBatteryIndex := 0
		for i := len(bank) - 1; i >= 0; i-- {
			digit, _ := strconv.Atoi(string(bank[i]))

			if digit >= firstBattery {
				firstBattery = digit
				firstBatteryIndex = i
			}
		}

		secondBattery := 0
		for j := len(bank) - 1; j > firstBatteryIndex; j-- {
			digit, _ := strconv.Atoi(string(bank[j]))

			if digit >= secondBattery {
				secondBattery = digit
			}
		}

		if secondBattery == 0 {
			secondBattery = firstBattery
			firstBattery = 0

			for k := 0; k < firstBatteryIndex; k++ {
				digit, _ := strconv.Atoi(string(bank[k]))

				if digit >= firstBattery {
					firstBattery = digit
				}
			}
		}

		total += firstBattery*10 + secondBattery
	}

	fmt.Println(total)
}
