package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	content, err := readFromFile()

	if err != nil {
		log.Fatal(err)
	} else {
		processLines(content)
	}

}

func printFirstAndLastNumbers(lineWithEnglishNumber string) (int, int) {

	line := replaceEnglishNumbers(lineWithEnglishNumber)

	// 过滤掉非数字字符
	var numbers []string
	for _, r := range line {
		if unicode.IsDigit(r) {
			numbers = append(numbers, string(r))
		}
	}

	if len(numbers) > 0 {
		firstNumber, _ := strconv.Atoi(numbers[0])
		lastNumber, _ := strconv.Atoi(numbers[len(numbers)-1])

		fmt.Printf("%d%d\n", firstNumber, lastNumber)

		return firstNumber, lastNumber
	}

	return 0, 0
}

func processLines(input string) {
	lines := strings.Split(input, "\n")
	totalSum := 0

	for _, line := range lines {
		first, last := printFirstAndLastNumbers(line)
		totalSum += first*10 + last
	}

	fmt.Printf("Total: %d\n\n", totalSum)
}

func readFromFile() (string, error) {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return "", err
	}

	return string(content), nil
}

var englishToArabic = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func replaceEnglishNumbers(input string) string {
	// 使用正则表达式匹配英语数字
	re := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine)`)
	result := re.ReplaceAllStringFunc(input, func(match string) string {
		// 获取匹配到的英语数字对应的阿拉伯数字
		englishNumber := match
		if arabicNumber, found := englishToArabic[match]; found {
			englishNumber = arabicNumber
		}

		// 返回阿拉伯数字的字符串形式
		return englishNumber
	})

	return result
}
