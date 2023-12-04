package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	count := 1
	var points float64 = 0
	for scanner.Scan() {
		initSlice := strings.Split(scanner.Text(), ":")
		// fmt.Println(initSlice[1])
		numSlices := strings.Split(initSlice[1], "|")
		winNums := strings.Split(numSlices[0], ",")
		myNums := strings.Split(numSlices[1], ",")
		winNums = strings.Split(winNums[0], " ")
		myNums = strings.Split(myNums[0], " ")
		numWins := 0
		for i, v := range winNums {
			winNums[i] = strings.Trim(v, " ")
		}
		for i, v := range myNums {
			myNums[i] = strings.Trim(v, " ")
		}
		for _, v := range winNums {
			for _, w := range myNums {
				if v == w && v != "" && w != "" {
					fmt.Println("Match found: ", v, " ", w)
					numWins++
				}
			}
		}
		if numWins > 0 {
			points += math.Pow(float64(2), float64(numWins-1))
		}
		count++
	}
	fmt.Println("Total points: ", points)

}
