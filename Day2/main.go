package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFile(path string) []string {
	var res []string
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Cant open the file 😬")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res
}

func parseInput(lines []string) [][]int {
	var result [][]int
	for _, line := range lines {
		parts := strings.Split(line, " ")
		var temp []int
		for _, value := range parts {
			num, _ := strconv.Atoi(value)
			temp = append(temp, num)
		}
		result = append(result, temp)
	}
	return result
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func checkSafe(levels []int) bool {
	for i := 0; i < len(levels)-1; i++ {
		distance := absInt(levels[i] - levels[i+1])
		if distance < 1 || distance > 3 {
			return false
		}
	}
	return true
}

func checkContinues(levels []int) bool {
	isIncreasing := levels[0] < levels[len(levels)-1]
	for i := 0; i < len(levels)-1; i++ {
		if isIncreasing && levels[i] >= levels[i+1] {
			return false
		}
		if !isIncreasing && levels[i] <= levels[i+1] {
			return false
		}
	}
	return true
}

func isSafeWithOneRemoval(levels []int) bool {
	if checkSafe(levels) && checkContinues(levels) {
		return true
	}

	for i := 0; i < len(levels); i++ {
		modified := append([]int{}, levels[:i]...)
		modified = append(modified, levels[i+1:]...)
		if checkSafe(modified) && checkContinues(modified) {
			return true
		}
	}
	return false
}

func one(lines []string) {
	input := parseInput(lines)
	res := 0
	for _, line := range input {
		if checkContinues(line) && checkSafe(line) {
			res += 1
		}
	}
	fmt.Println(res)
}

func two(lines []string) {
	input := parseInput(lines)
	res := 0
	for _, line := range input {
		if isSafeWithOneRemoval(line) {
			fmt.Println(line)
			res++
		}
	}
	fmt.Println(res)
}

func main() {
	lines := ReadFile("./test.txt")
	one(lines)
	two(lines)

}
