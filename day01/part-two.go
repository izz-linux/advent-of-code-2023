package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	_ "strings"
)

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}

func convertToNum(s string) int {
	i := 10
	switch s {
	case "one", "1":
		i = 1
	case "two", "2":
		i = 2
	case "three", "3":
		i = 3
	case "four", "4":
		i = 4
	case "five", "5":
		i = 5
	case "six", "6":
		i = 6
	case "seven", "7":
		i = 7
	case "eight", "8":
		i = 8
	case "nine", "9":
		i = 9
	case "zero", "0":
		i = 0
	}
	return i
}

func main() {
	file, err := os.Open("input2.txt")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	re := regexp.MustCompile("[0-9]|zero|one|two|three|four|five|six|seven|eight|nine")
	sum := 0
	for scanner.Scan() {
		s := strings.Join(re.FindAllString(scanner.Text(), 1), "")
		x := convertToNum(s)

		count := 1
		y := 10
		for true {
			fromRight := strings.Join(re.FindAllString(scanner.Text()[len(scanner.Text())-count:], 1), "")
			y = convertToNum(fromRight)
			if y != 10 {
				break
			} else {
				count++
			}
		}
		//
		//t := strings.Join(re.FindAllString(reverse(scanner.Text()), 1), "")
		//y := convertToNum(t)

		thisSum := x*10 + y
		fmt.Println(x)
		fmt.Println(y)
		sum += thisSum
	}
	fmt.Println(sum)
}
