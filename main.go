package main

import (
	"fmt"

	"github.com/lruchandani/datastructures-in-go/sorts"
)

func takeInput() []int {
	fmt.Println("Enter size of array:")
	var noOfInts = 0
	fmt.Scan(&noOfInts)
	fmt.Println("Enter array")
	arr := make([]int, noOfInts)
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}
	return arr
}

func main() {
	for {
		fmt.Println("Selection Sort (ASC) = 1")
		fmt.Println("Selection Sort (DSC) = 2")
		fmt.Println("Bubble Sort (ASC) = 3")
		fmt.Println("Bubble Sort (DESC) = 4")
		fmt.Println("Exit : 5")
		fmt.Println("Choose Input type :")
		var choice = 0
		fmt.Scanf("%d", &choice)
		switch choice {
		case 1:
			sortMethod(sorts.SelectionSort, takeInput(), sorts.ASC)
		case 2:
			sortMethod(sorts.SelectionSort, takeInput(), sorts.DESC)
		case 3:
			sortMethod(sorts.BubbleSort, takeInput(), sorts.ASC)
		case 4:
			sortMethod(sorts.BubbleSort, takeInput(), sorts.DESC)
		default:
			return
		}

	}
}

func sortMethod(f sorts.Fn, arr []int, sortType sorts.Sort) {
	f(arr, sortType)
	fmt.Printf("%v\n", arr)
}