package main

import (
	"bufio"
	"fmt"
	"os"
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

func parseInput(lines []string) ([][]rune, []int) {
	var mapp [][]rune
	var guard []int
	for i, line := range lines {
		runeSlice := []rune(line)
		mapp = append(mapp, runeSlice)
		x := strings.IndexRune(line, '^')
		if x != -1 {
			guard = []int{x, i}
		}

	}
	return mapp, guard
}

func rotateClockwise(x []int) (int, int) {
	return x[1], -x[0]
}

func prettyPrint(matrix [][]rune) {
	for _, row := range matrix {
		fmt.Println(string(row))
	}
}

func one(lines []string) {
	count := 1
	mapp, guard := parseInput(lines)
	actualDirection := []int{0, 1}
	mapp[guard[1]][guard[0]] = '0'
	for i := 0; i < 100000000; i++ {
		newGuardX := guard[0] + actualDirection[0]
		newGuardY := guard[1] - actualDirection[1]
		if newGuardX == len(mapp[0]) || newGuardX < 0 || newGuardY == len(mapp) || newGuardY < 0 {
			break
		}
		if mapp[newGuardY][newGuardX] == '#' {
			nDX, nDY := rotateClockwise(actualDirection)
			actualDirection = []int{nDX, nDY}
		}
		if mapp[newGuardY][newGuardX] == '.' {
			guard = []int{newGuardX, newGuardY}
			mapp[guard[1]][guard[0]] = '0'
			count += 1
		}
		if mapp[newGuardY][newGuardX] == '0' {
			guard = []int{newGuardX, newGuardY}
		}

		//prettyPrint(mapp)
		//fmt.Println("----------------------------")

	}
	fmt.Println(count)
}

func two(lines []string) {
	mapp, guard := parseInput(lines)

	directions := [][]int{
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, 0},
	}

	isLoop := func(matrix [][]rune, start []int) bool {
		visited := make(map[[3]int]bool)
		pos := start
		dir := 0

		for {
			state := [3]int{pos[0], pos[1], dir}
			if visited[state] {
				return true
			}
			visited[state] = true

			nextX := pos[0] + directions[dir][0]
			nextY := pos[1] + directions[dir][1]

			if nextX < 0 || nextX >= len(matrix[0]) || nextY < 0 || nextY >= len(matrix) || matrix[nextY][nextX] == '#' {
				dir = (dir + 1) % 4
			} else {
				pos = []int{nextX, nextY}
			}

			if nextX < 0 || nextX >= len(matrix[0]) || nextY < 0 || nextY >= len(matrix) {
				break
			}
		}
		return false
	}

	validPlacements := 0
	for y := 0; y < len(mapp); y++ {
		for x := 0; x < len(mapp[0]); x++ {
			if mapp[y][x] == '.' {
				mapp[y][x] = '#'
				if isLoop(mapp, guard) {
					validPlacements++
				}
				mapp[y][x] = '.'
			}
		}
	}

	fmt.Println(validPlacements)
}

func main() {
	lines := ReadFile("./test.txt")
	one(lines)
	two(lines)

}
