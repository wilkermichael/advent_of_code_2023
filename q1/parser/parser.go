package parser

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func Parse(filepath string, parseFunc func([]string) (int, error)) (int, error) {
	// Open a file handler
	f, err := os.Open(filepath)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	s := make([]string, 0)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		s = append(s, scanner.Text())
	}

	return parseFunc(s)
}

func ParseSum(lines []string) (int, error) {
	return parseLine(lines, parseStringToInt)
}

func ParseSumWithStrings(lines []string) (int, error) {
	return parseLine(lines, parseStringstoIntsWithStrings)
}

func parseLine(lines []string, parseFunc func(s string) (int, error)) (int, error) {
	nums := make([]int, 0)
	for _, l := range lines {
		n, err := parseFunc(l)
		if err != nil {
			return 0, err
		}
		nums = append(nums, n)
	}

	sum := 0
	for _, n := range nums {
		sum = sum + n
	}

	return sum, nil
}

type numPos struct {
	num string
	pos int
}

func parseStringstoIntsWithStrings(s string) (int, error) {
	// Create a list of possible string numbers
	stringNumbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	// Build our list of numbers to position
	np := make([]numPos, 0)
	for i, sn := range stringNumbers {
		index := strings.Index(s, sn)
		if index != -1 {
			np = append(np, numPos{num: fmt.Sprintf("%d", i+1), pos: index})
		}
		index = strings.LastIndex(s, sn)
		if index != -1 {
			np = append(np, numPos{num: fmt.Sprintf("%d", i+1), pos: index})
		}
	}

	// Add position of digits
	for i, r := range s {
		if unicode.IsDigit(r) {
			np = append(np, numPos{num: string(r), pos: i})
		}
	}

	// Sort by position
	sort.SliceStable(np, func(i, j int) bool {
		return np[j].pos > np[i].pos
	})

	firstDigit := np[0].num
	secondDigit := np[len(np)-1].num
	return strconv.Atoi(firstDigit + secondDigit)
}

func parseStringToInt(s string) (int, error) {
	firstDigit := ""
	secondDigit := ""
	for _, r := range s {
		if unicode.IsDigit(r) {
			if firstDigit == "" {
				firstDigit = string(r)
			} else {
				secondDigit = string(r)
			}
		}
	}
	// If second digit is still empty, then both fhe first and last digit are the same
	if secondDigit == "" {
		secondDigit = firstDigit
	}
	return strconv.Atoi(firstDigit + secondDigit)
}
