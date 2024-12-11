package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
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

func parseInput(line string) []int {
	var res []int
	for _, v := range line {
		num, _ := strconv.Atoi(string(v))
		res = append(res, num)
	}
	return res
}

func calculateMap(blocks []int) []string {
	var res []string
	id := 0
	for i := 0; i < len(blocks)-1; i += 2 {
		for j := 0; j < blocks[i]; j++ {
			res = append(res, strconv.Itoa(id))
		}
		for j := 0; j < blocks[i+1]; j++ {
			res = append(res, ".")
		}
		id += 1
	}
	if len(blocks)%2 != 0 {
		for j := 0; j < blocks[len(blocks)-1]; j++ {
			res = append(res, strconv.Itoa(id))
		}
	}
	return res
}

func getLastIntIndex(x []string) int {
	for i := len(x) - 1; i >= 0; i-- {
		if _, err := strconv.Atoi(x[i]); err == nil {
			return i
		}
	}
	return -1
}

func order(blocks []string) {
	lastInt := getLastIntIndex(blocks)
	for i := range blocks {
		if blocks[i] == "." {
			blocks[i] = blocks[lastInt]
			blocks[lastInt] = "."
			lastInt = getLastIntIndex(blocks)
		}
	}
	if _, err := strconv.Atoi(blocks[len(blocks)-1]); err == nil {
		i := slices.Index(blocks, ".")
		blocks[i] = blocks[len(blocks)-1]
		blocks[len(blocks)-1] = "."

	}
	fmt.Println(blocks)
	count := 0
	for i, v := range blocks {
		b, _ := strconv.Atoi(v)
		count += i * b
	}
	fmt.Println(count)
}

func one(line string) {
	input := parseInput(line)
	mapp := calculateMap(input)
	order(mapp)
}

func getConsecutive(blocks []string, index int) int {
	count := 0
	for i := index; i < len(blocks)-1; i++ {
		if blocks[i] == blocks[i+1] {
			count++
		} else {
			break
		}
	}
	return count
}

func firstIndexAndCountOfLastConsecutiveInts(slice []string) (int, int) {
	lastIndex := -1
	count := 0
	n := len(slice)

	for i := n - 1; i >= 0; {
		val, err := strconv.Atoi(slice[i])
		if err == nil {
			sequenceStart := i
			for i >= 0 {
				nextVal, err := strconv.Atoi(slice[i])
				if err != nil || nextVal != val {
					break
				}
				i--
			}
			if lastIndex == -1 {
				lastIndex = i + 1
				count = sequenceStart - i
			}
		} else {
			i--
		}
	}

	return lastIndex, count
}

func removeElement(slice []string, index int) []string {
	if index < 0 || index >= len(slice) {
		fmt.Println("Index out of bounds")
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func orderTwo(blocks []string) {
	dotBlocks := FindAllDotBlocks(blocks)
	numberBlocks := FindBlocks(blocks)

	sort.Slice(dotBlocks, func(i, j int) bool {
		return dotBlocks[i].Index < dotBlocks[j].Index
	})
	sort.Slice(numberBlocks, func(i, j int) bool {
		return len(numberBlocks[i].Value) < len(numberBlocks[j].Value)
	})

	for _, dotBlock := range dotBlocks {
		for i, numberBlock := range numberBlocks {
			if len(numberBlock.Value) >= dotBlock.Count {
				blocks = blocks[:dotBlock.Index] + blocks[dotBlock.Index+dotBlock.Count:]
				for _, char := range numberBlock.Value {
					blocks = append(blocks[:dotBlock.Index], string(char))
					blocks = append(blocks[dotBlock.Index+1:]...)
				}
				numberBlocks = removeElement(numberBlocks, i)
				break
			}
		}
	}

	fmt.Println(blocks)

	count := 0
	for i, v := range blocks {
		b, _ := strconv.Atoi(v)
		count += i * b
	}
	fmt.Println(count)
}

type Block struct {
	Value string
	Index int
	Count int
}

func FindBlocks(slice []string) []Block {
	var blocks []Block
	n := len(slice)

	for i := 0; i < n; {
		val := slice[i]
		if _, err := strconv.Atoi(val); err != nil {
			i++
			continue
		}

		start := i
		for i < n && slice[i] == val {
			i++
		}
		blocks = append(blocks, Block{
			Value: val,
			Index: start,
			Count: i - start,
		})
	}
	return blocks
}

type DotBlock struct {
	Index int
	Count int
}

func FindAllDotBlocks(slice []string) []DotBlock {
	var dotBlocks []DotBlock
	n := len(slice)

	for i := 0; i < n; {
		if slice[i] != "." {
			i++
			continue
		}
		start := i
		for i < n && slice[i] == "." {
			i++
		}

		dotBlocks = append(dotBlocks, DotBlock{
			Index: start,
			Count: i - start,
		})
	}

	return dotBlocks
}

func two(line string) {
	input := parseInput(line)
	mapp := calculateMap(input)
	orderTwo(mapp)
}

func main() {
	lines := ReadFile("./test.txt")
	//one(lines[0])
	two(lines[0])

}
