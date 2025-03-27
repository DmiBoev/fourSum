package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 0, -1, 0, -2, 2}
	target := 0
	fmt.Print(fourSum(a, target))

}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums) // Сортируем массив
	var matrix [][]int
	n := len(nums)

	for i := 0; i < n-3; i++ {
		// Пропускаем дубликаты для i
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			// Пропускаем дубликаты для j
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left := j + 1
			right := n - 1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					matrix = append(matrix, []int{nums[i], nums[j], nums[left], nums[right]})
					// Пропускаем дубликаты для left и right
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return matrix
}
