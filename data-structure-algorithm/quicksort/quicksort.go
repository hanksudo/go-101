package main

import "log"

func quicksort(arr []int, low, high int) {
	if low >= high {
		return
	}
	pivot := arr[(low+high)/2]
	idx := partition(arr, low, high, pivot)
	quicksort(arr, low, idx-1)
	quicksort(arr, idx, high)
}

func partition(arr []int, low, high, pivot int) int {
	for low <= high {
		for arr[low] < pivot {
			low++
		}
		for arr[high] > pivot {
			high--
		}
		if low <= high {
			arr[low], arr[high] = arr[high], arr[low]
			low++
			high--
		}
	}
	return low
}

func main() {
	nums := []int{3, 7, 2, 9, 4}
	quicksort(nums, 0, len(nums)-1)
	log.Println(nums)
}
