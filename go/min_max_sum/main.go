package main

import "fmt"

func main() {
	input := [5]int32{1, 2, 3, 4, 5}
	fmt.Println(minMaxSum(input))
}

/*Given five positive integers, find the minimum and maximum values that can be calculated by summing exactly four of the five integers. Then print the respective minimum and maximum values as a single line of two space-separated long integers.

Example

The minimum sum is  and the maximum sum is . The function prints*/

func minMaxSum(arr [5]int32) string {
	var (
		min      = arr[0]
		max      = arr[0]
		minIndex = 0
		maxIndex = 0
		minSum   int
		maxSum   int
		i        int
	)
	for i = 1; i < len(arr); i++ {
		if arr[i] <= min {
			min = arr[i]
			minIndex = i
		}
		if arr[i] >= max {
			max = arr[i]
			maxIndex = i
		}
	}

	for i = 0; i < len(arr); i++ {
		if i != maxIndex {
			minSum += int(arr[i])
		}

		if i != minIndex {
			maxSum += int(arr[i])
		}
	}

	return fmt.Sprintf("%d %d", minSum, maxSum)
}
