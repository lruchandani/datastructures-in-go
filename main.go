package main

import (
	"fmt"

	queue "github.com/lruchandani/datastructures-in-go/queues"
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
		fmt.Println("Sorting = 1")
		fmt.Println("Queue = 2")
		var choice = 0
		fmt.Scanf("%d", &choice)
		switch choice {
		case 1:
			sorting()
		case 2:
			queuing()
		default:
			fmt.Print("Invalid Choice ", choice)
			return
		}
	}
}

func queuing() {
	q := queue.MyQueue{}

	for {
		fmt.Println("Add to Queue = 1")
		fmt.Println("Print Queue = 2")
		fmt.Println("Size Queue = 3")
		var choice = 0
		fmt.Scanf("%d", &choice)
		switch choice {
		case 1:
			fmt.Println("Enter element to add:")
			var e = 0
			fmt.Scan(&e)
			q.Add(e)
		case 2:
			q.Print()
		case 3:
			fmt.Print("Size is ", q.Size())
		default:
			fmt.Print("Inavlid Choice")
			return

		}
	}
}
func sorting() {
	for {
		fmt.Println("Selection Sort (ASC) = 1")
		fmt.Println("Selection Sort (DSC) = 2")
		fmt.Println("Bubble Sort (ASC) = 3")
		fmt.Println("Bubble Sort (DESC) = 4")
		fmt.Println("InsertionSort Sort = 5")
		fmt.Println("Quick Sort = 6")
		fmt.Println("Exit : Anything else")
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
		case 5:
			sortMethod(sorts.InsertionSort, takeInput(), sorts.DESC)
		case 6:
			sortMethod(sorts.QuickSort, takeInput(), sorts.ASC)
		default:
			return
		}

	}
}

func sortMethod(f sorts.Fn, arr []int, sortType sorts.Sort) {
	f(arr, sortType)
	fmt.Printf("%v\n", arr)
}
