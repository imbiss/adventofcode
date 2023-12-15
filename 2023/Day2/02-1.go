package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	content, err := readFromFile()

	if err != nil {
		log.Fatal(err)
		return
	}

	processLines(content)

}

func readFromFile() (string, error) {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return "", err
	}

	return string(content), nil
}

func processLines(input string) {
	lines := strings.Split(input, "\n")
	totalSum := 0

	for _, line := range lines {

		gameRound, groupResults, err := processLine(line)

		if err != nil {
			fmt.Printf("Error %v\n", err)
			return
		}

		result := getLineResult(groupResults)

		fmt.Printf("%s, result=%d\n", gameRound, result)
		totalSum += result

	}

	fmt.Printf("Total: %d\n\n", totalSum)
}

func processLine(input string) (string, string, error) {
	parts := strings.Split(input, ":")

	if len(parts) != 2 {
		return "", "", fmt.Errorf("Can not split %s", input)
	}

	return parts[0], parts[1], nil
}

func splitBySemicolon(input string) ([]string, error) {
	parts := strings.Split(input, ";")

	if len(parts) < 1 {
		return nil, fmt.Errorf("can not split by semicolon %s", input)
	}

	return parts, nil
}

// innput: blue, 4 red; 1 red, 2 green, 6 blue; 2 green
func getLineResult(groupResults string) int {

	groups, err := splitBySemicolon(groupResults)

	if err != nil {
		fmt.Printf("Error %v\n", err)
		return 0
	}

	if len(groups) < 1 {
		fmt.Printf("no gourp found\n")
		return 0
	}

	redMax, blueMax, greenMax := 0, 0, 0

	for _, group := range groups {
		red, blue, green := parseGroup(group)
		if red > redMax {
			redMax = red
		}

		if blue > blueMax {
			blueMax = blue
		}

		if green > greenMax {
			greenMax = green
		}
	}

	return redMax * blueMax * greenMax

}

/* input like "3 blue, 4 red"  or "3 green, 15 blue, 14 red" */
func parseGroup(input string) (int, int, int) {
	fmt.Printf("parse group: %s\n", input)
	red, blue, green := 0, 0, 0
	colors := strings.Split(input, ",")

	for _, color := range colors {

		color, num := parseEachColor(color)

		if color == "red" {
			red = num
		}

		if color == "blue" {
			blue = num
		}

		if color == "green" {
			green = num
		}
	}
	fmt.Printf("red=%d, blue=%d, green=%d\n", red, blue, green)
	return red, blue, green
}

// input like "3 blue" or "4 red" or "3 green"
func parseEachColor(color string) (string, int) {
	color = strings.TrimSpace(color)
	parts := strings.Split(color, " ")

	num := parts[0]
	colorName := parts[1]

	count, err := strconv.Atoi(num)

	if err != nil {
		fmt.Println("Conv failed", err)
		return "", 0
	}

	return colorName, count
}

func getLastDigit(input string) int {
	var lastDigit string

	lastDigit = input[5:]

	digit, err := strconv.Atoi(lastDigit)

	if err != nil {
		fmt.Println("can not convert to int!")

		return 0
	}

	fmt.Printf("+ %d\n", digit)

	return digit
}
