package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

func one(line string) {
	pattern := `mul\(\d{1,3},\d{1,3}\)`
	re := regexp.MustCompile(pattern)
	operations := re.FindAllString(line, -1)
	res := 0
	for _, v := range operations {
		parts := strings.Split(v[4:len(v)-1], ",")
		first, _ := strconv.Atoi(parts[0])
		second, _ := strconv.Atoi(parts[1])
		res += first * second
	}
	fmt.Println(res)

}

func two(line string) {
	pattern := `(mul\(\d{1,3},\d{1,3}\))|(do\(\))|(don't\(\))`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(line, -1)
	fmt.Println(matches)
	res := 0
	do := true
	for _, v := range matches {
		if v[0] == 'm' && do {
			parts := strings.Split(v[4:len(v)-1], ",")
			first, _ := strconv.Atoi(parts[0])
			second, _ := strconv.Atoi(parts[1])
			res += first * second
		}
		if v == "do()" {
			do = true
		}
		if v == "don't()" {
			do = false
		}
	}
	fmt.Println(res)

}

func main() {
	lines := ReadFile("./test.txt")
	fmt.Println(lines[0])
	one(lines[0])
	two(lines[0])

}
