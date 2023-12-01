package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	re := regexp.MustCompile("[0-9]")
	sum := 0
	for scanner.Scan() {
		s := strings.Join(re.FindAllString(scanner.Text(), 1), "")
		x, _ := strconv.Atoi(s)
		t := strings.Join(re.FindAllString(reverse(scanner.Text()), 1), "")
		y, _ := strconv.Atoi(t)

		thisSum := x*10 + y
		fmt.Println(x)
		fmt.Println(y)
		sum += thisSum
	}
	fmt.Println(sum)
}
