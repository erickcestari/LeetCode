package main

import (
	"fmt"
)

func boobleSortArray(nums []int) []int {
	n := len(nums)
	sorted := false

	for !sorted {
		sorted = true
		for i := 0; i < n-1; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				sorted = false
			}
		}
	}

	return nums
}

func takeMedian(nums []int) float64 {
	n := len(nums)
	mid := n / 2

	if n%2 == 1 {
		median := float64(nums[mid])
		return median
	} else {
		median := float64((nums[mid-1] + nums[mid])) / 2
		return median
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	numsMerged := append(nums1, nums2...)

	sortedArray := boobleSortArray(numsMerged)

	return takeMedian(sortedArray)
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	var median = findMedianSortedArrays(nums1, nums2)
	fmt.Println(median)
}
