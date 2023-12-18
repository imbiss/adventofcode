package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type ScratchCard struct {
	Card           int
	WinningNumbers []int
	YourNumbers    []int
	Points         int
}

func main() {
	part1("input.txt")
}

func part1(fileName string) {
	fmt.Printf("Reading from file: %s.\n", fileName)
	lines := loadSchema(fileName)

	allCards := []ScratchCard{}

	for _, line := range lines {
		isOk, card := parseLine(line)
		if false == isOk {
			log.Fatal("Parse line failed")
		}
		allCards = append(allCards, card)
	}

	sum := totalPoints(allCards)

	fmt.Println("==========")
	fmt.Printf("sum: %d\n", sum)
	fmt.Println("==========")
}

func parseLine(input []rune) (bool, ScratchCard) {
	line := string(input)
	parts := strings.Split(line, "|")

	if len(parts) != 2 {
		log.Fatal("Failed to split header")
		return false, ScratchCard{}
	}

	header := strings.Split(parts[0], ":")
	if len(header) != 2 {
		log.Fatal("Faild to split card number")
		return false, ScratchCard{}
	}

	card := ScratchCard{}
	card.Card = parseCardNumber(header[0])
	card.WinningNumbers = parseWinningNumbers(header[1])
	card.YourNumbers = parseYourNumbers(parts[1])
	card.Points = CountingPoints(card.WinningNumbers, card.YourNumbers)

	return true, card
}

func parseCardNumber(header string) int {
	cardNumberStr := header[len("Card "):]
	cardNumberStr = strings.TrimSpace(cardNumberStr)

	cardNumber, err := strconv.Atoi(cardNumberStr)
	if err != nil {
		fmt.Println(cardNumberStr)
		log.Fatal("Failed to parse card number to integer")
		return 0
	}

	return cardNumber
}

func parseWinningNumbers(heaer string) []int {
	return parseNumbers(heaer)
}

func parseNumbers(header string) []int {
	header = strings.TrimSpace(header)
	numStrings := strings.Split(header, " ")

	var numbers []int

	for _, numStr := range numStrings {
		if "" != numStr {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println(numStr)
				log.Fatal("Failed to parse the last nubmers")
				return []int{}
			}
			numbers = append(numbers, num)
		}

	}

	return numbers
}

func parseYourNumbers(nums string) []int {
	return parseNumbers(nums)
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
		fmt.Println(line)
	}

	return engineSchematic
}

func CountingPoints(winNum, yourNum []int) int {
	count := 0

	for _, win := range winNum {
		for _, your := range yourNum {
			if your == win {
				count++
			}
		}
	}

	point := int(math.Pow(float64(2), float64(count-1)))
	return point
}

func totalPoints(cards []ScratchCard) int {

	//fmt.Println(cards)
	sum := 0
	for _, card := range cards {
		sum += card.Points
	}

	return sum
}
