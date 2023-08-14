package main

import "fmt"

func main() {
	arr := []int{1, 5, 7, 3, 4, 6, 8, 0}
	PickSort(arr)
}

func BubbleSort(arr []int) {
	for i := 0; i+1 < len(arr); i++ {
		for j := 0; j+1 < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				k := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = k
			}
		}
	}

	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
	}
}

func PickSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := arr[i]
		min_i := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < min {
				min = arr[j]
				min_i = j
			}
		}
		if min_i != i {
			k := arr[i]
			arr[i] = min
			arr[min_i] = k
		}
	}

	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
	}
}
