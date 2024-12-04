package main

import (
	"fmt"
	"strings"
	"utils"
)

type Mul struct {
	a int
	b int
}

func parseString(stream *string, pos *int, val string) bool {
	if *pos+len(val) <= len(*stream) && (*stream)[*pos:*pos+len(val)] == val {
		(*pos) += len(val)
		return true
	}
	return false
}

func parseChar(stream *string, pos *int, char byte) bool {
	if *pos < len(*stream) && (*stream)[*pos] == char {
		(*pos)++
		return true
	}
	return false
}

func parseNum(stream *string, pos *int) (bool, int) {
	isNum := false
	num := 0

	for ; *pos < len(*stream); (*pos)++ {
		c := (*stream)[*pos]
		if c < '0' || c > '9' {
			break
		}
		isNum = true
		num = num*10 + int(c-'0')
	}
	return isNum, num
}

func parseFunc2(stream *string, pos *int, name string) (bool, int, int) {
	p := *pos
	if parseString(stream, &p, name+"(") {
		g1, n1 := parseNum(stream, &p)
		if g1 {
			if parseChar(stream, &p, ',') {
				g2, n2 := parseNum(stream, &p)
				if g2 {
					if parseChar(stream, &p, ')') {
						*pos = p
						return true, n1, n2
					}
				}
			}
		}
	}
	return false, 0, 0
}

func main() {
	lines, _ := utils.ReadLines("input.txt")

	stream := strings.Join(lines, "")

	total := 0
	pos := 0
	for pos < len(stream) {
		if success, n1, n2 := parseFunc2(&stream, &pos, "mul"); success {
			total += n1 * n2
		} else {
			pos++
		}
	}
	fmt.Printf("Part 1 %v\n", total)

	total2 := 0
	enabled := true
	pos = 0
	for pos < len(stream) {
		if parseString(&stream, &pos, "do()") {
			enabled = true
		} else if parseString(&stream, &pos, "don't()") {
			enabled = false
		} else if success, n1, n2 := parseFunc2(&stream, &pos, "mul"); success {
			if enabled {
				total2 += n1 * n2
			}
		} else {
			pos++
		}
	}

	fmt.Printf("Part 2 %v\n", total2)
}
