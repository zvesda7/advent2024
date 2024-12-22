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

func add(p1 Point, p2 Point) Point {
	return Point{p1.X + p2.X, p1.Y + p2.Y}
}

func inBounds(p Point, width, height int) bool {
	if p.X < 0 || p.Y < 0 {
		return false
	}
	if p.X >= width || p.Y >= height {
		return false
	}
	return true
}

func calcAntiNodes1(p1, p2 Point, width, height int) []Point {
	var points []Point
	a1 := Point{p2.X + p2.X - p1.X, p2.Y + p2.Y - p1.Y}
	a2 := Point{p1.X + p1.X - p2.X, p1.Y + p1.Y - p2.Y}
	if inBounds(a1, width, height) {
		points = append(points, a1)
	}
	if inBounds(a2, width, height) {
		points = append(points, a2)
	}
	return points
}

func calcAntiNodes2(p1, p2 Point, width, height int) []Point {
	var points []Point
	d1 := Point{p2.X - p1.X, p2.Y - p1.Y}
	d2 := Point{p1.X - p2.X, p1.Y - p2.Y}
	for inBounds(p1, width, height) {
		points = append(points, p1)
		p1 = add(p1, d1)
	}
	for inBounds(p2, width, height) {
		points = append(points, p2)
		p2 = add(p2, d2)
	}
	//fmt.Println(d1, d2, points)
	return points
}

func main() {
	lines, _ := utils.ReadLines("input.txt")

	height := len(lines)
	width := len(lines[0])

	freqs := map[byte][]Point{}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			freq := byte(lines[y][x])
			if freq != '.' {
				freqs[freq] = append(freqs[freq], Point{x, y})
			}
		}
	}

	antinodes1 := map[PointHash]bool{}
	antinodes2 := map[PointHash]bool{}
	for _, points := range freqs {
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				for _, node := range calcAntiNodes1(points[i], points[j], width, height) {
					antinodes1[node.hash()] = true
				}
				for _, node := range calcAntiNodes2(points[i], points[j], width, height) {
					antinodes2[node.hash()] = true
					//fmt.Println(node)
				}
			}
		}
	}
	fmt.Printf("Part 1 %v\n", len(antinodes1))
	fmt.Printf("Part 2 %v\n", len(antinodes2))
}
