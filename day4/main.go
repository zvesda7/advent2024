package main

import (
	"fmt"
	"utils"
)

type Coord int

type Point struct {
	X, Y int
}

func (p Point) hash() int {
	return p.X*1000 + p.Y
}
func add(p1 Point, p2 Point) Point {
	return Point{p1.X + p2.X, p1.Y + p2.Y}
}

func main() {
	lines, _ := utils.ReadLines("input.txt")

	points := map[int]byte{}
	for y, row := range lines {
		for x, char := range row {
			points[Point{y, x}.hash()] = byte(char)
		}
	}

	dirs := []Point{
		{0, -1},
		{1, -1},
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
	}

	match := "XMAS"
	total := 0
	for y, row := range lines {
		for x, _ := range row {
			p := Point{x, y}
			for _, dir := range dirs {
				testp := p
				good := true
				for i := 0; i < len(match); i++ {
					if match[i] != points[testp.hash()] {
						good = false
						break
					}
					testp = add(testp, dir)
				}
				if good {
					total += 1
				}
			}
		}
	}
	fmt.Printf("Part 1 %v\n", total)

	matches := map[string]bool{
		"MSAMS": true,
		"SSAMM": true,
		"MMASS": true,
		"SMASM": true,
	}
	total2 := 0
	for y, row := range lines {
		for x, _ := range row {
			p := Point{x, y}
			s := string(points[add(p, Point{-1, -1}).hash()])
			s += string(points[add(p, Point{1, -1}).hash()])
			s += string(points[add(p, Point{0, 0}).hash()])
			s += string(points[add(p, Point{-1, 1}).hash()])
			s += string(points[add(p, Point{1, 1}).hash()])
			if matches[s] {
				total2 += 1
			}

		}
	}
	fmt.Printf("Part 2 %v\n", total2)

}
