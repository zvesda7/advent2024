package utils

import (
	"bufio"
	"os"
)

func Test() string {
	return "Test Worked"
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func AbsInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
