package project

import "sort"

func FindOddInt(arr []int) int {
	sort.Ints(arr)
	count := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			count++
		} else {
			if count%2 != 0 {
				return arr[i-1]
			}
			count = 1
		}
	}
	return arr[len(arr)-1]
}
