package main

import "fmt"

func selectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		minIndex := i

		for j := i; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		if minIndex != i {
			arr[minIndex], arr[i] = arr[i], arr[minIndex]
		}
	}

	fmt.Println(arr)
}

func main() {
	arr := []int{4, 2, 3, 1, 6, 7, 9}
	selectSort(arr)
}
