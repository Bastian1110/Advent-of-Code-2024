package main

import (
	"bufio"
	"fmt"
	"os"
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

func parseInput(lines []string) [][]rune {
	var mapp [][]rune
	for _, line := range lines {
		runeSlice := []rune(line)
		mapp = append(mapp, runeSlice)
	}
	return mapp
}

func mapOfAnthenas(mapp [][]rune) map[rune][][]int {
	res := make(map[rune][][]int)
	for i, line := range mapp {
		for j := range line {
			if mapp[i][j] != '.' {
				res[mapp[i][j]] = append(res[mapp[i][j]], []int{j, i})
			}
		}

	}
	return res
}

func distanceComponents(point1, point2 []int) (dx, dy int) {
	dx = point2[0] - point1[0]
	dy = point2[1] - point1[1]
	return dx, dy
}

func prettyPrint(matrix [][]rune) {
	for _, row := range matrix {
		fmt.Println(string(row))
	}
}

func one(lines []string) {
	result := 0
	mapp := parseInput(lines)
	anthenas := mapOfAnthenas(mapp)
	for _, v := range anthenas {
		for i := 0; i < len(v); i++ {
			for j := i; j < len(v); j++ {
				if i != j {
					fmt.Println("Points : ", v[i], v[j])
					dOX, dOY := distanceComponents(v[i], v[j])
					dTX, dTY := distanceComponents(v[j], v[i])
					fmt.Println("Distance One : ", dOX, dOY)
					fmt.Println("Distance Two: ", dTX, dTY)
					antiNodeOneX := v[i][0] - dOX
					antiNodeOneY := v[i][1] - dOY
					antiNodeTwoX := v[j][0] - dTX
					antiNodeTwoY := v[j][1] - dTY
					fmt.Println("Point 1 : ", antiNodeOneX, antiNodeOneY)
					fmt.Println("Point 2 : ", antiNodeTwoX, antiNodeTwoY)
					if antiNodeOneX >= 0 && antiNodeOneX < len(mapp[0]) && antiNodeOneY >= 0 && antiNodeOneY < len(mapp) {
						mapp[antiNodeOneY][antiNodeOneX] = '#'
					}
					if antiNodeTwoX >= 0 && antiNodeTwoX < len(mapp[0]) && antiNodeTwoY >= 0 && antiNodeTwoY < len(mapp) {
						mapp[antiNodeTwoY][antiNodeTwoX] = '#'
					}
					prettyPrint(mapp)
				}
			}
		}
	}
	for i := 0; i < len(mapp); i++ {
		for j := 0; j < len(mapp[0]); j++ {
			if mapp[i][j] == '#' {
				result += 1
			}

		}
	}
	fmt.Println(result)
}

func two(lines []string) {
	result := 0
	mapp := parseInput(lines)
	anthenas := mapOfAnthenas(mapp)
	for _, v := range anthenas {
		for i := 0; i < len(v); i++ {
			for j := i; j < len(v); j++ {
				if i != j {
					fmt.Println("Points : ", v[i], v[j])
					dOX, dOY := distanceComponents(v[i], v[j])
					dTX, dTY := distanceComponents(v[j], v[i])
					fmt.Println("Distance One : ", dOX, dOY)
					fmt.Println("Distance Two: ", dTX, dTY)
					first := false
					var lastOne []int
					var lastTwo []int

					oneDead := false
					twoDead := false
					for true {
						var antiNodeOneX int
						var antiNodeOneY int
						var antiNodeTwoX int
						var antiNodeTwoY int
						if !first {
							antiNodeOneX = v[i][0] - dOX
							antiNodeOneY = v[i][1] - dOY
							antiNodeTwoX = v[j][0] - dTX
							antiNodeTwoY = v[j][1] - dTY
							first = true
						} else {
							antiNodeOneX = lastOne[0] - dOX
							antiNodeOneY = lastOne[1] - dOY
							antiNodeTwoX = lastTwo[0] - dTX
							antiNodeTwoY = lastTwo[1] - dTY
						}
						fmt.Println("Point 1 : ", antiNodeOneX, antiNodeOneY)
						fmt.Println("Point 2 : ", antiNodeTwoX, antiNodeTwoY)
						if antiNodeOneX >= 0 && antiNodeOneX < len(mapp[0]) && antiNodeOneY >= 0 && antiNodeOneY < len(mapp) {
							mapp[antiNodeOneY][antiNodeOneX] = '#'
						} else {
							oneDead = true
						}
						if antiNodeTwoX >= 0 && antiNodeTwoX < len(mapp[0]) && antiNodeTwoY >= 0 && antiNodeTwoY < len(mapp) {
							mapp[antiNodeTwoY][antiNodeTwoX] = '#'
						} else {
							twoDead = true
						}
						lastOne = []int{antiNodeOneX, antiNodeOneY}
						lastTwo = []int{antiNodeTwoX, antiNodeTwoY}
						if oneDead && twoDead {
							break
						}
					}
					prettyPrint(mapp)
				}
			}
		}
	}
	for i := 0; i < len(mapp); i++ {
		for j := 0; j < len(mapp[0]); j++ {

			_, exists := anthenas[mapp[i][j]]
			if exists || mapp[i][j] == '#' {
				result += 1
			}

		}
	}
	fmt.Println(result)
}

func main() {
	lines := ReadFile("./test.txt")
	//one(lines)
	two(lines)

}
