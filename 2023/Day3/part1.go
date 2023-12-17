package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	part1("input.txt")
	part2("input.txt")
}

type Result struct {
	Numbers   []string
	Positions [][]int
}

type ResultRatio struct {
	Numbers   []string
	Positions [][]int
	Ratio     []int
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

	a.searchRowEnd = myResult.Positions[i][0] + 1
	if a.searchRowEnd > rowNum-1 {
		a.searchRowEnd = rowNum - 1
	}

	a.searchColStart = myResult.Positions[i][1] - 1
	if a.searchColStart < 0 {
		a.searchColStart = 0
	}

	a.searchColEnd = myResult.Positions[i][1] + len(myResult.Numbers[i])
	if a.searchColEnd > colNum-1 {
		a.searchColEnd = colNum - 1
	}
	fmt.Printf("Searching area: (%d, %d), (%d, %d)\n", a.searchRowStart, a.searchColStart, a.searchRowEnd, a.searchColEnd)

	return a
}

func defineGearSearchArea(i, j, maxRow, maxCol int) SearchArea {
	a := SearchArea{}

	a.searchRowStart = i - 1
	if a.searchRowStart < 0 {
		a.searchRowStart = 0
	}

	a.searchRowEnd = i + 1
	if a.searchRowEnd > maxRow {
		a.searchRowEnd = maxRow
	}

	a.searchColStart = j - 1
	if a.searchColStart < 0 {
		a.searchColStart = 0
	}

	a.searchColEnd = j + 1
	if a.searchColEnd > maxCol {
		a.searchColEnd = maxCol
	}

	return a
}

func areaSearch(a SearchArea, maxtrix [][]rune, toCheck string) bool {
	foundParts := false

	for row := a.searchRowStart; row <= a.searchRowEnd && false == foundParts; row++ {
		fmt.Printf("row %d/%d start.", row, a.searchRowEnd)

		for colum := a.searchColStart; colum <= a.searchColEnd && false == foundParts; colum++ {
			if isSymbol(maxtrix[row][colum]) {
				if !foundParts {
					foundParts = true
					fmt.Print("Found!\n")
				}
			} else {
				fmt.Print(".")
			}
		}
	}
	return foundParts
}

func printResult(result Result) {
	fmt.Println("Numbers:", result.Numbers)
	fmt.Println("Pos", result.Positions)
}

func printResultRatio(result ResultRatio) {
	fmt.Println("Numbers:", result.Numbers)
	fmt.Println("Pos", result.Positions)
	fmt.Println("Ratio", result.Ratio)
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

func findAllGears(matrix [][]rune, allNumbers Result) ResultRatio {

	myResult := ResultRatio{}
	isGearResult := false
	ratio := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			//fmt.Println("scaning (", i, ",", j, ")")
			isGearResult, ratio = isGearAndRatio(matrix, i, j, allNumbers)
			if isGearResult {
				// found *
				pos := []int{i, j}
				myResult.Numbers = append(myResult.Numbers, string(matrix[i][j]))
				myResult.Positions = append(myResult.Positions, pos)
				myResult.Ratio = append(myResult.Ratio, ratio)
				fmt.Println("Ratio saved")
			} else {
				// fmt.Println("Not *")
			}
		}
	}
	fmt.Println("Found ", len(myResult.Numbers), "x * totaly")
	return myResult
}

func isGearAndRatio(matrix [][]rune, i, j int, allNumbers Result) (bool, int) {
	char := matrix[i][j]
	maxRow := len(matrix)
	maxCol := len(matrix[i])

	if char == '*' {
		searchArea := defineGearSearchArea(i, j, maxRow, maxCol)
		isFound, ratio := findGearAdjacentNum(searchArea, matrix, allNumbers)
		if isFound {
			fmt.Println("Found gear, ratio=", ratio)
		}
		return isFound, ratio
	} else {
		return false, 0
	}

}

func containsValue(slice []int, toCheck int) bool {
	for _, v := range slice {
		if v == toCheck {
			return true
		}
	}

	return false
}

func findGearAdjacentNum(area SearchArea, matrix [][]rune, allNumbers Result) (bool, int) {

	buffer := []int{}
	index := []int{}

	for i := area.searchRowStart; i <= area.searchRowEnd; i++ {
		for j := area.searchColStart; j <= area.searchColEnd; j++ {

			if unicode.IsDigit(matrix[i][j]) {
				number, idx := scanAdjacentNumber(i, j, matrix, allNumbers)
				if !containsValue(index, idx) {
					buffer = append(buffer, number)
					index = append(index, idx)
				} else {
					fmt.Printf("The number %d is already in buffer\n", number)
				}

			}
		}
	}

	if len(buffer) != 2 {
		fmt.Println("not exactly 2 parts. not gear.")
		return false, 0
	}

	gearRatio := 1
	for _, num := range buffer {
		gearRatio *= num
	}

	return true, gearRatio
}

func scanAdjacentNumber(i, j int, matrix [][]rune, allNumbers Result) (int, int) {

	fmt.Printf("scan (%d,%d) for number\n", i, j)
	count := 0
	for _, pos := range allNumbers.Positions {
		row := pos[0]
		col := pos[1]
		number := allNumbers.Numbers[count]
		{
		}
		if row == i && j >= col && j <= col+len(number) {
			num, err := strconv.Atoi(number)
			if err != nil {
				fmt.Println("Can not convert to int", err)
			}
			fmt.Println("Found matched number", num)
			return num, count
		}
		count++
	}
	return 0, 0
}

func algoPart1(matrix [][]rune) {
	allNumbers := findAllNumbers(matrix)

	printResult(allNumbers)

	sum := sumAdjacent(allNumbers, len(matrix), len(matrix[0]), matrix)

	fmt.Println("===================")
	fmt.Println("Part1 Sum:", sum)
	fmt.Println("===================")
}

func algoPart2(matrix [][]rune) {
	allNumbers := findAllNumbers(matrix)
	allGears := findAllGears(matrix, allNumbers)

	fmt.Println("all Gears info:")
	printResultRatio(allGears)

	sum := sumGearRatios(allGears.Ratio)

	fmt.Println("===================")
	fmt.Println("Part2 Sum:", sum)
	fmt.Println("===================")
}

func sumAdjacent(allNumbers Result, rowNum int, colNum int, maxtrix [][]rune) int64 {

	sum := int64(0)
	fmt.Println("rowNum:", rowNum, "colNum:", colNum)

	for i := 0; i < len(allNumbers.Numbers); i++ {
		toCheck := allNumbers.Numbers[i]
		fmt.Printf("\nChecking number: %s(length:%d)\n", toCheck, len(toCheck))

		area := defineSearchArea(allNumbers, i, rowNum, colNum)
		found := areaSearch(area, maxtrix, toCheck)

		if found {
			sum += toInt(toCheck)
		}
	}
	return sum
}

func sumGearRatios(ratios []int) int {
	sum := 0
	for _, r := range ratios {
		sum += r
	}
	return sum
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

func part1(fileName string) {
	fmt.Printf("Reading from file: %s.\n", fileName)
	algoPart1(loadSchema(fileName))
}

func loadSchema(fileName string) [][]rune {
	content, err := os.ReadFile(fileName)
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

	return engineSchematic
}

func part2(fileName string) {
	fmt.Printf("Reading from file: %s.\n", fileName)
	algoPart2(loadSchema(fileName))
}

func printTextContent(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
}
