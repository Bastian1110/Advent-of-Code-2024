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

func parseInput(lines []string) ([]int, [][]int) {
	var result []int
	var equation [][]int
	for _, line := range lines {
		parts := strings.Split(line, ":")
		numbers := strings.Split(parts[1], " ")
		var temp []int
		for _, n := range numbers[1:] {
			r, _ := strconv.Atoi(n)
			temp = append(temp, r)
		}
		equation = append(equation, temp)
		y, _ := strconv.Atoi(parts[0])
		result = append(result, y)
	}
	return result, equation
}

func generateTernaryConfigurations(n int) [][]int {
	if n <= 0 {
		return nil
	}

	totalConfigs := 1
	for i := 0; i < n; i++ {
		totalConfigs *= 3
	}

	result := make([][]int, totalConfigs)

	for i := 0; i < totalConfigs; i++ {
		config := make([]int, n)
		num := i
		for j := 0; j < n; j++ {
			config[j] = num % 3
			num /= 3
		}
		result[i] = config
	}

	return result
}

func generateBooleanConfigurations(n int) [][]bool {
	totalConfigs := 1 << n
	result := make([][]bool, totalConfigs)

	for i := 0; i < totalConfigs; i++ {
		config := make([]bool, n)
		for j := 0; j < n; j++ {
			config[j] = (i & (1 << j)) != 0
		}
		result[i] = config
	}

	return result
}

func execute(equation []int, ops []bool) int {
	res := equation[0]
	for i := 1; i < len(equation); i++ {
		if ops[i-1] {
			res *= equation[i]
		} else {
			res += equation[i]
		}
	}
	return res
}

func isSolvableOne(y int, equation []int) bool {
	configs := generateBooleanConfigurations(len(equation) - 1)
	for _, v := range configs {
		if execute(equation, v) == y {
			return true
		}
	}
	return false
}

func one(lines []string) {
	ys, equations := parseInput(lines)
	count := 0
	for i := range equations {
		if isSolvableOne(ys[i], equations[i]) {
			count += ys[i]
		}
	}
	fmt.Println(count)

}

func executeTwo(equation []int, ops []int) int {
	res := equation[0]
	for i := 1; i < len(equation); i++ {
		if ops[i-1] == 0 {
			res *= equation[i]
		} else if ops[i-1] == 1 {
			res += equation[i]
		} else {
			merged := strconv.Itoa(res) + strconv.Itoa(equation[i])
			mergedNum, _ := strconv.Atoi(merged)
			res = mergedNum
		}
	}
	return res
}

func isSolvableTwo(y int, equation []int) bool {
	configs := generateTernaryConfigurations(len(equation) - 1)
	for _, v := range configs {
		if executeTwo(equation, v) == y {
			return true
		}
	}
	return false
}

func two(lines []string) {
	ys, equations := parseInput(lines)
	count := 0
	for i := range equations {
		if isSolvableTwo(ys[i], equations[i]) {
			count += ys[i]
		}
	}
	fmt.Println(count)

}
func main() {
	lines := ReadFile("./test.txt")
	one(lines)
	two(lines)

}
