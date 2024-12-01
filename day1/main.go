package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"utils"
)

func main() {
	lines, _ := utils.ReadLines("input.txt")

	//parse
	list_a := []int{}
	list_b := []int{}
	for _, row := range lines {
		items := strings.Split(row, "   ")
		a, _ := strconv.Atoi(items[0])
		b, _ := strconv.Atoi(items[1])
		list_a = append(list_a, a)
		list_b = append(list_b, b)
	}
	sort.Ints(list_a)
	sort.Ints(list_b)

	//part 1
	sum := 0
	for i, a := range list_a {
		sum += utils.AbsInt(a - list_b[i])
	}
	fmt.Printf("Part 1 %v\n", sum)

	//part 2
	counts := map[int]int{}
	for _, b := range list_b {
		counts[b]++
	}
	sum2 := 0
	for _, a := range list_a {
		sum2 += a * counts[a]
	}
	fmt.Printf("Part 2 %v\n", sum2)
}
