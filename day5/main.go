package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

type Deps map[int]map[int]bool

func checkOrder(deps *Deps, update []int) (bool, int, int) {
	for i := 0; i < len(update)-1; i++ {
		for j := i + 1; j < len(update); j++ {
			n0 := update[i]
			n1 := update[j]
			if x, ok := (*deps)[n0]; ok {
				if _, ok2 := x[n1]; ok2 {
					return false, i, j
				}
			}
		}
	}
	return true, 0, 0
}

func main() {
	lines, _ := utils.ReadLines("input.txt")

	deps := Deps{}

	y := 0
	for ; y < len(lines); y++ {
		if lines[y] == "" {
			break
		}
		pair := strings.Split(lines[y], "|")
		n0, _ := strconv.Atoi(pair[0])
		n1, _ := strconv.Atoi(pair[1])
		if _, ok := deps[n1]; !ok {
			deps[n1] = map[int]bool{}
		}
		deps[n1][n0] = true
	}

	sum := 0
	badUpdates := [][]int{}
	for y++; y < len(lines); y++ {
		updateStrs := strings.Split(lines[y], ",")
		update := []int{}
		for _, s := range updateStrs {
			n, _ := strconv.Atoi(s)
			update = append(update, n)
		}
		good, _, _ := checkOrder(&deps, update)
		if good {
			sum += update[len(update)/2]
		} else {
			badUpdates = append(badUpdates, update)
		}
	}
	fmt.Printf("Part 1 %v\n", sum)

	swapDone := true
	for swapDone {
		swapDone = false
		for i := 0; i < len(badUpdates); i++ {
			good, n0, n1 := checkOrder(&deps, badUpdates[i])
			if !good {
				swapDone = true
				temp := badUpdates[i][n0]
				badUpdates[i][n0] = badUpdates[i][n1]
				badUpdates[i][n1] = temp
			}
		}
	}

	sum2 := 0
	for _, update := range badUpdates {
		sum2 += update[len(update)/2]
	}
	fmt.Printf("Part 2 %v\n", sum2)
}
