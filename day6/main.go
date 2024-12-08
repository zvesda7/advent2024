package main

import (
	"fmt"
	"utils"
)

type Point struct {
	X, Y int
}

type PointHash int

func (p Point) hash() PointHash {
	return PointHash(p.X*1000 + p.Y)
}

func (p Point) hashWithDir(dir int) PointHash {
	return PointHash(p.X*10000 + p.Y*10 + (dir % 4))
}

func add(p1 Point, p2 Point) Point {
	return Point{p1.X + p2.X, p1.Y + p2.Y}
}

func inBounds(x, y, width, height int) bool {
	if x < 0 || y < 0 {
		return false
	}
	if x >= width || y >= height {
		return false
	}
	return true
}

var DIRS = []Point{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

func guardLoops(guard Point, height int, width int, walls map[PointHash]bool) bool {
	curDirI := 0
	walked := map[PointHash]bool{}
	for inBounds(guard.X, guard.Y, width, height) {
		if _, found := walked[guard.hashWithDir(curDirI)]; found {
			return true
		} else {
			walked[guard.hashWithDir(curDirI)] = true
		}

		test := add(guard, DIRS[curDirI%4])
		if wall, ok := walls[test.hash()]; ok && wall {
			curDirI++ //wall hit, turn right
		} else {
			guard = test
		}
	}
	return false
}

func main() {
	lines, _ := utils.ReadLines("input.txt")

	walls := map[PointHash]bool{}
	guardStart := Point{}
	height := len(lines)
	width := len(lines[0])
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if lines[y][x] == '#' {
				walls[Point{x, y}.hash()] = true
			} else if lines[y][x] == '^' {
				guardStart = Point{x, y}
			}
		}
	}

	guard := guardStart
	curDirI := 0
	walked := map[PointHash]bool{}
	for inBounds(guard.X, guard.Y, width, height) {
		walked[guard.hash()] = true

		test := add(guard, DIRS[curDirI%4])
		if wall, ok := walls[test.hash()]; ok && wall {
			curDirI++ //wall hit, turn right
		} else {
			guard = test
		}
	}
	fmt.Printf("Part 1 %v\n", len(walked))

	walked[guardStart.hash()] = false
	totalLooping := 0
	for p, _ := range walked {
		walls[p] = true
		if guardLoops(guardStart, height, width, walls) {
			totalLooping++
		}
		walls[p] = false
	}
	fmt.Printf("Part 2 %v\n", totalLooping)
}
