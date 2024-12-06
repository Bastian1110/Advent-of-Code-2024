package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(path string) [][]rune {
	var res [][]rune
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Cannot open the file 😬")
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		res = append(res, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
		return nil
	}

	return res
}

func searchXmas(matrix [][]rune, word string, startX, startY, dirX, dirY int) bool {
	wordLength := len(word)
	for i := 0; i < wordLength; i++ {
		newX := startX + i*dirX
		newY := startY + i*dirY

		if newX < 0 || newX >= len(matrix) || newY < 0 || newY >= len(matrix[0]) {
			return false
		}
		if matrix[newX][newY] != rune(word[i]) {
			return false
		}
	}
	return true
}

func one(matrix [][]rune) {
	directions := [][2]int{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0},
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
	}

	wordCount := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			for _, dir := range directions {
				if searchXmas(matrix, "XMAS", i, j, dir[0], dir[1]) {
					wordCount++
				}
			}
		}
	}
	fmt.Println(wordCount)
}

func isMAS(grid [][]rune, startRow, startCol, dx, dy int) bool {
	rows, cols := len(grid), len(grid[0])

	if startRow < 0 || startRow >= rows ||
		startCol < 0 || startCol >= cols ||
		startRow+dx < 0 || startRow+dx >= rows ||
		startCol+dy < 0 || startCol+dy >= cols ||
		startRow+2*dx < 0 || startRow+2*dx >= rows ||
		startCol+2*dy < 0 || startCol+2*dy >= cols {
		return false
	}

	if grid[startRow][startCol] == 'M' &&
		grid[startRow+dx][startCol+dy] == 'A' &&
		grid[startRow+2*dx][startCol+2*dy] == 'S' {
		return true
	}

	if grid[startRow][startCol] == 'S' &&
		grid[startRow+dx][startCol+dy] == 'A' &&
		grid[startRow+2*dx][startCol+2*dy] == 'M' {
		return true
	}

	return false
}

func two(grid [][]rune) int {
	rows, cols := len(grid), len(grid[0])
	xmasCount := 0

	// Possible X shape directions
	directions := [][2]int{
		{-1, -1}, {1, 1},
		{-1, 1}, {1, -1},
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			for _, dir := range directions {
				dx, dy := dir[0], dir[1]
				centerRow := row + dx
				centerCol := col + dy

				if centerRow >= 0 && centerRow < rows &&
					centerCol >= 0 && centerCol < cols {
					if isMAS(grid, row, col, dx, dy) &&
						isMAS(grid, row+2*dx, col+2*dy, -dx, -dy) {
						if centerRow >= 0 && centerRow < rows &&
							centerCol >= 0 && centerCol < cols &&
							grid[centerRow][centerCol] == 'A' {
							xmasCount++
						}
					}
				}
			}
		}
	}

	return xmasCount
}
func main() {
	lines := ReadFile("./test.txt")
	one(lines)
	fmt.Println(two(lines))

}
