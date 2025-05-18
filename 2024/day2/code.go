package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Parse file into two slices
func parseInputs() [][]int {
	var list1 [][]int

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

		var temp []int
		for _, val := range nums {
			num1, err1 := strconv.Atoi(val)

			if err1 != nil {
				panic(err1)
			}
			temp = append(temp, num1)
		}
		list1 = append(list1, temp)

		if err == io.EOF {
			break
		}
	}

	defer file.Close()

	return list1
}

func isNotSafe(a int, b int) bool {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}

	return (diff == 0) || (diff > 3)
}

func isSafe(s []int) bool {
	var safe bool = true
	var prev int = s[0]
	var inc bool = false
	for i, val := range s {
		if i == 0 {
			continue
		} else if i == 1 {
			if prev < val {
				inc = true
			}
		}

		if (inc && (prev > val)) ||
			(!inc && (prev < val)) ||
			(isNotSafe(prev, val)) {
			safe = false
			break
		}

		prev = val
	}

	return safe
}

func main() {
	list := parseInputs()

	// Part 1
	var numSafe int = 0
	for _, s := range list {
		safe := isSafe(s)

		if safe {
			numSafe += 1
		}
	}

	fmt.Print("Part 1: ", numSafe, "\n")

	// Part 2
	// var numSafe2 int = 0
	// for _, s := range list {
	// 	// fmt.Print(s)
	// 	var safe1 bool = true
	// 	var prev int = s[0]
	// 	var inc bool = false
	// 	var canRemove bool = true

	// 	// Attempt to check if need to remove all levels
	// 	// except the first one
	// 	for i, val := range s {
	// 		if i == 0 {
	// 			continue
	// 		} else if i == 1 {
	// 			if prev < val {
	// 				inc = true
	// 			}
	// 		}

	// 		if (inc && (prev > val)) ||
	// 			(!inc && (prev < val)) ||
	// 			(isNotSafe(prev, val)) {
	// 			if canRemove {
	// 				canRemove = false

	// 				if (i+1 < len(s)) && s[i-1] < val {
	// 					inc = true
	// 				} else if (i+1 < len(s)) && s[i-1] > val {
	// 					inc = false
	// 				}
	// 				continue
	// 			} else {
	// 				safe1 = false
	// 				break
	// 			}
	// 		} else {
	// 			prev = val
	// 		}
	// 	}

	// 	var safe2 bool = true
	// 	prev = s[1]
	// 	inc = false
	// 	// Check if first number can be removed
	// 	for i, val := range s {
	// 		if i == 0 || i == 1 {
	// 			continue
	// 		} else if i == 2 {
	// 			if prev < val {
	// 				inc = true
	// 			}
	// 		}

	// 		if (inc && (prev > val)) ||
	// 			(!inc && (prev < val)) ||
	// 			(isNotSafe(prev, val)) {
	// 			safe2 = false
	// 			break
	// 		}

	// 		prev = val
	// 	}

	// 	if safe1 || safe2 {
	// 		numSafe2 += 1
	// 	}
	// }

	var numSafe2 int = 0
	for _, s := range list {
		safe := false

		for i := range len(s) {
			var temp []int
			for j := range len(s) {
				if i != j {
					temp = append(temp, s[j])
				}
			}

			check := isSafe(temp)

			if check {
				safe = true
			}
		}
		if safe {
			numSafe2 += 1
		}

	}

	fmt.Print("Part 2: ", numSafe2, "\n")
}
