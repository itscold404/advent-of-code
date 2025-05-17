package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Parse file into two slices
func parseInputs() ([]int, []int) {
	var list1 []int
	var list2 []int

	file, ferr := os.Open("input.txt")

	if ferr != nil {
		panic(ferr)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}

		nums := strings.Fields(line)
		num1, err1 := strconv.Atoi(nums[0])
		num2, err2 := strconv.Atoi(nums[1])

		if err1 != nil {
			panic(err1)
		}

		if err2 != nil {
			panic(err2)
		}

		list1 = append(list1, num1)
		list2 = append(list2, num2)

		if err == io.EOF {
			break
		}
	}

	defer file.Close()

	return list1, list2
}

// Find the distance between two values
func findDistance(int1 int, int2 int) int {
	diff := int1 - int2
	if diff < 0 {
		return -diff
	}

	return diff
}

func main() {
	list1, list2 := parseInputs()

	if len(list1) != len(list2) {
		panic("mismatch length")
	}

	// PART 1:
	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})
	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})

	var diff int
	for i := range list1 {
		diff += findDistance(list1[i], list2[i])
	}
	fmt.Print("Part 1: ", diff, "\n")

	//PART 2:
	// make map to store each occurance of a number
	var count map[int]int = make(map[int]int)
	for _, val := range list2 {
		count[val] += 1
	}

	var simScore int = 0
	for _, val := range list1 {
		simScore += val * count[val]
	}
	fmt.Print("Part 2: ", simScore, "\n")
}
