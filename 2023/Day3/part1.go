package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	solvePuzzle("example.txt")
}

type Result struct {
	Numbers   []string
	Positions [][]int
}

type SearchArea struct {
	searchRowStart int
	searchRowEnd   int
	searchColStart int
	searchColEnd   int
}

func defineSearchArea(myResult Result, i, rowNum, colNum int) SearchArea {
	a := SearchArea{}

	a.searchRowStart = myResult.Positions[i][0] - 1
	if a.searchRowStart < 0 {
		a.searchRowStart = 0
	}
	fmt.Println("searchRowStart:", a.searchRowStart)

	a.searchRowEnd = myResult.Positions[i][0] + 1
	if a.searchRowEnd > rowNum-1 {
		a.searchRowEnd = rowNum - 1
	}
	fmt.Println("searchRowEnd:", a.searchRowEnd)

	a.searchColStart = myResult.Positions[i][1] - 1
	if a.searchColStart < 0 {
		a.searchColStart = 0
	}
	fmt.Println("searchColStart:", a.searchColStart)

	a.searchColEnd = myResult.Positions[i][1] + len(myResult.Numbers[i])
	if a.searchColEnd > colNum-1 {
		a.searchColEnd = colNum - 1
	}
	fmt.Println("searchColEnd:", a.searchColEnd)

	return a
}

func areaSearch(a SearchArea, maxtrix [][]rune, toCheck string) bool {
	foundParts := false

	for row := a.searchRowStart; row <= a.searchRowEnd && false == foundParts; row++ {
		fmt.Printf("row %d/%d start.\n", row, a.searchRowEnd)

		for colum := a.searchColStart; colum <= a.searchColEnd && false == foundParts; colum++ {
			fmt.Printf("col %d start.\n", colum)
			if isSymbol(maxtrix[row][colum]) {
				if !foundParts {
					foundParts = true
				}
			}
			fmt.Printf("colum %d finished.\n", colum)
		}

		fmt.Printf("row %d finished.\n", row)
		fmt.Println("")
	}

	if !foundParts {
		fmt.Printf("%s未找到符号\n", toCheck)
		return false
	} else {
		fmt.Printf("%s找到了相邻符\n", toCheck)
		return true
	}
}

func printResult(result Result) {
	fmt.Println("Numbers:", result.Numbers)
	fmt.Println("Pos", result.Positions)
}

func scanHorizontalContinuousNumbers(matix [][]rune, row, col int) (string, []int) {
	// the first digit
	num := string(matix[row][col])
	pos := []int{row, col}

	for j := col + 1; j < len(matix[row]) && unicode.IsDigit(matix[row][j]); j++ {
		num += string(matix[row][j])
	}

	return num, pos
}

func findAllNumbers(matrix [][]rune) Result {
	myResult := Result{}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if unicode.IsDigit(matrix[i][j]) {
				num, pos := scanHorizontalContinuousNumbers(matrix, i, j)
				if num != "" {
					// found something
					myResult.Numbers = append(myResult.Numbers, num)
					myResult.Positions = append(myResult.Positions, pos)
					j = j + len(num) - 1
				}
			}

		}
	}

	return myResult
}

func algoTwo(matrix [][]rune) {
	allNumbers := findAllNumbers(matrix)

	printResult(allNumbers)

	sumAdjacent := sumAdjacent(allNumbers, len(matrix), len(matrix[0]), matrix)
	sum := sumAdjacent

	fmt.Println("Adjacent: ", sumAdjacent)
	fmt.Println("===================")
	fmt.Println("Sum:", sum)
}

func sumAdjacent(allNumbers Result, rowNum int, colNum int, maxtrix [][]rune) int64 {

	sum := int64(0)
	fmt.Println("rowNum:", rowNum, "colNum:", colNum)

	for i := 0; i < len(allNumbers.Numbers); i++ {
		toCheck := allNumbers.Numbers[i]
		fmt.Printf("Checking result %s(length:%d)\n", toCheck, len(toCheck))

		area := defineSearchArea(allNumbers, i, rowNum, colNum)
		found := areaSearch(area, maxtrix, toCheck)

		if found {
			sum += toInt(toCheck)
		}

		if false {
			x := allNumbers.Positions[i][0]
			y := allNumbers.Positions[i][1]
			if isDiagonally(x, y, len(maxtrix)) {
				sum += toInt(toCheck)
			}
		}
	}
	return sum
}

func isDiagonally(x, y, size int) bool {

	if x == y {
		return true
	}

	if x+y == size {
		return true
	}

	return false
}

func isSymbol(x rune) bool {

	if unicode.IsDigit(x) || x == '.' {
		return false
	}

	return true
}

func toInt(a string) int64 {
	num, err := strconv.ParseInt(a, 10, 64)
	if err != nil {
		fmt.Println("failed to cast to int")
		return 0
	}
	return num
}

func solvePuzzle(fileName string) {

	fmt.Printf("Reading from file: %s.\n", fileName)
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	fmt.Printf("Line: %d\n", len(lines))
	fmt.Println()

	engineSchematic := make([][]rune, len(lines))
	for i, line := range lines {
		engineSchematic[i] = []rune(line)
	}

	printTextContent(lines)

	algoTwo(engineSchematic)
}

func printTextContent(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
}
