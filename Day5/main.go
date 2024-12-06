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

func parseInput(lines []string) ([][]int, [][]int) {
	space := false
	var rules [][]int
	var updates [][]int
	for _, v := range lines {
		if v == "" {
			space = true
			continue
		}
		if !space {
			parts := strings.Split(v, "|")
			left, _ := strconv.Atoi(parts[0])
			right, _ := strconv.Atoi(parts[1])
			rules = append(rules, []int{left, right})
		} else {
			parts := strings.Split(v, ",")
			var temp []int
			for _, n := range parts {
				num, _ := strconv.Atoi(n)
				temp = append(temp, num)
			}
			updates = append(updates, temp)
		}
	}
	return rules, updates
}

func findIndex(slice []int, target int) int {
	for i, num := range slice {
		if num == target {
			return i
		}
	}
	return -1
}

func checkCorrectOrder(rules [][]int, update []int) bool {
	for _, rule := range rules {
		left := findIndex(update, rule[0])
		right := findIndex(update, rule[1])
		if left > right && left != -1 && right != -1 {
			return false
		}
	}
	return true
}

func one(rules [][]int, updates [][]int) {
	res := 0
	for _, u := range updates {
		if checkCorrectOrder(rules, u) {
			middleIndex := len(u) / 2
			res += u[middleIndex]
		}
	}
	fmt.Println("One : ", res)

}

func overloadWithRule(a, b int, rules [][]int) bool {
	for _, r := range rules {
		if r[0] == a && r[1] == b {
			return false
		}
	}
	return true
}

func orderUpdate(rules [][]int, update []int) []int {
	n := len(update)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if overloadWithRule(update[j], update[j+1], rules) {
				update[j], update[j+1] = update[j+1], update[j]
			}
		}
	}
	return update
}

func two(rules [][]int, updates [][]int) {
	res := 0
	for _, u := range updates {
		if !checkCorrectOrder(rules, u) {
			fmt.Println("Update : ", u)
			update := orderUpdate(rules, u)
			fmt.Println("Updated : ", update)
			middleIndex := len(update) / 2
			res += update[middleIndex]
		}
	}
	fmt.Println("Two : ", res)

}

func main() {
	lines := ReadFile("./test.txt")
	rules, updates := parseInput(lines)
	//one(rules, updates)
	two(rules, updates)

}
