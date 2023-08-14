package main

import "fmt"

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	findBestSum(arr)
}

func findBestSum(arr []int) {
	bestSum := 0
	currSum := 0

	for i := 0; i < len(arr); i++ {
		currSum += arr[i]
		if currSum <= 0 {
			currSum = 0
		}
		if currSum > bestSum {
			bestSum = currSum
		}
	}

	fmt.Print(bestSum)
}
