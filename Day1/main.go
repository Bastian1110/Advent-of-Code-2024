package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func parseInput(lines []string) ([]int, []int) {
	var left []int
	var right []int
	for _, line := range lines {
		parts := strings.Split(line, " ")
		numL, _ := strconv.Atoi(parts[0])
		numR, _ := strconv.Atoi(parts[3])
		left = append(left, numL)
		right = append(right, numR)
	}
	return left, right
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func sumSlice(values []int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}

func one(lines []string) {
	left, right := parseInput(lines)
	sort.Ints(left)
	sort.Ints(right)
	var distances []int
	for i, _ := range left {
		distances = append(distances, absInt(left[i]-right[i]))
	}
	fmt.Println(sumSlice(distances))

}

func two(lines []string) {
	left, right := parseInput(lines)
	appearances := make(map[int]int)
	for _, v := range right {
		_, exists := appearances[v]
		if exists {
			appearances[v] += 1
		} else {
			appearances[v] = 1
		}
	}
	var results []int
	for _, v := range left {
		_, exists := appearances[v]
		if exists {
			results = append(results, appearances[v]*v)
		}
	}
	fmt.Println(sumSlice(results))

}

func main() {
	lines := ReadFile("./test.txt")
	one(lines)
	two(lines)

}
