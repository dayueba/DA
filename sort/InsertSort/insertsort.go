package main

import "fmt"

func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		val := arr[i]
		j := i
		for ; j > 0 && val < arr[j-1]; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = val
	}

	fmt.Println(arr)
}

func main() {
	arr := []int{4, 2, 3, 1, 9, 6, 7}
	insertSort(arr)
}
