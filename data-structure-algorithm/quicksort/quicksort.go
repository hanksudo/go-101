package main

import "log"

func Quicksort(arr []int, low, high int) {
	if low >= high {
		return
	}

	idx := partition(arr, low, high)
	Quicksort(arr, low, idx-1)
	Quicksort(arr, idx+1, high)
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[high], arr[i+1] = arr[i+1], arr[high]
	return i + 1
}

func main() {
	nums := []int{3, 7, 2, 9, 4}
	Quicksort(nums, 0, len(nums)-1)
	log.Println(nums)
}
