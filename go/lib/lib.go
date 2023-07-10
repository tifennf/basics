package lib

import "math/rand"

const upper_bound = 1000

func GenerateRandArr(len int) []int {

	res := make([]int, len)

	for i := 0; i < len; i++ {
		res[i] = rand.Intn(upper_bound)
	}

	return res
}

func Min(arr []int, len int) int {
	// Pre: len(arr) != 0
	var min int = arr[0]

	for i := 1; i < len; i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}

	return min
}

func Max(arr []int, len int) int {
	// Pre: len(arr) != 0
	var max int = arr[0]

	for i := 1; i < len; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}

func Average(arr []int, len int) float64 {
	s := 0.0

	for i := 0; i < len; i++ {
		s += float64(arr[i])
	}

	return s / float64(len)

}

func Sort(arr []int, len int) {
	// tri à bulles
	for i := 0; i < len; i++ {
		for j := 0; j < len-(i+1); j++ {
			if arr[j] > arr[j+1] {
				// transpose both
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
			}
		}
	}
}
