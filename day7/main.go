package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"utils"
)

func findHit(nums []int, find int) bool {
	for i := 0; i < (1 << (len(nums) - 1)); i++ {
		if calcTest(nums, i) == find {
			return true
		}
	}
	return false
}

func calcTest(nums []int, mask int) int {
	total := nums[0]
	for i := 1; i < len(nums); i++ {
		if (mask & 1) > 0 {
			total += nums[i]
		} else {
			total *= nums[i]
		}
		mask >>= 1
	}
	return total
}

func findHit3(nums []int, find int) bool {
	cnt := len(nums) - 1
	iters := int(math.Pow(3, float64(cnt)))

	for i := 0; i < iters; i++ {
		if calcTest3(nums, i) == find {
			return true
		}
	}
	return false
}

func calcTest3(nums []int, mask int) int {
	total := nums[0]
	for i := 1; i < len(nums); i++ {
		if mask%3 == 0 {
			total += nums[i]
		} else if mask%3 == 1 {
			total *= nums[i]
		} else if mask%3 == 2 {
			total, _ = strconv.Atoi(strconv.Itoa(total) + strconv.Itoa(nums[i]))

		}
		mask /= 3
	}
	return total
}

func main() {
	lines, _ := utils.ReadLines("input.txt")

	total1 := 0
	total2 := 0
	for _, line := range lines {
		split1 := strings.Split(line, ": ")
		tv, _ := strconv.Atoi(split1[0])
		split2 := strings.Split(split1[1], " ")
		nums := []int{}
		for _, n := range split2 {
			ni, _ := strconv.Atoi(n)
			nums = append(nums, ni)
		}
		if findHit(nums, tv) {
			total1 += tv
		}
		if findHit3(nums, tv) {
			total2 += tv
		}
	}
	fmt.Printf("Part 1 %v\n", total1)
	fmt.Printf("Part 2 %v\n", total2)

}
