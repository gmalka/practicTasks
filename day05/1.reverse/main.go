package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	fmt.Println("test 1 prev: ", arr1)
	reverse(arr1)
	fmt.Println("test 1 after: ", arr1)

	arr2 := []int{}
	fmt.Println("test 2 prev: ", arr2)
	reverse(arr2)
	fmt.Println("test 2 after: ", arr2)

	arr3 := []int{1}
	fmt.Println("test 3 prev: ", arr3)
	reverse(arr3)
	fmt.Println("test 3 after: ", arr3)

	arr4 := []int{1, 2}
	fmt.Println("test 4 prev: ", arr4)
	reverse(arr4)
	fmt.Println("test 4 after: ", arr4)
}

func reverse(arr []int) {
	for i, j := 0, len(arr) - 1; i < j; i, j = i + 1, j - 1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}