// Package sorts implements additional functions to manipulate UTF-8
// encoded strings, beyond what is provided in the standard "strings" package.
package sorts

//Fn fucntion interface
type Fn func([]int, Sort)

// Sort Enum
type Sort string

// Sorts enum definition
const (
	DESC Sort = "DESC"
	ASC  Sort = "ASC"
)

// SelectionSort function
func SelectionSort(arr []int, sortType Sort) {
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			var x = false
			if sortType == DESC {
				x = arr[i] < arr[j]
			} else {
				x = arr[i] > arr[j]
			}
			if x {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
}

//BubbleSort fucntion
func BubbleSort(arr []int, sortType Sort) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			var x = false
			if sortType == DESC {
				x = arr[j] < arr[j+1]
			} else {
				x = arr[j] > arr[j+1]
			}
			if x {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
}

// InsertionSort function
func InsertionSort(arr []int, sortType Sort) {
	for i := 1; i < len(arr); i++ {
		v := arr[i]
		j := i
		if sortType == ASC {
			for ; j >= 1 && arr[j-1] > v; j-- {
				arr[j] = arr[j-1]
			}
		} else {
			for ; j >= 1 && arr[j-1] < v; j-- {
				arr[j] = arr[j-1]
			}
		}
		arr[j] = v
	}
}

//QuickSort Function
func QuickSort(arr []int, sortType Sort) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, start int, end int) {
	if start > end {
		return
	}
	mid := findMid(arr, start, end)
	quickSort(arr, start, mid-1)
	quickSort(arr, mid+1, end)
}

func findMid(arr []int, start int, end int) int {
	if end < start {
		return start
	}
	v := arr[end]
	startPos, endPos := start, end-1
	for startPos <= endPos {
		if arr[startPos] > v {
			swap(arr, startPos, endPos)
			endPos--
		} else {
			startPos++
		}
	}
	swap(arr, startPos, end)
	return startPos
}

func swap(arr []int, i int, j int) {
	t := arr[i]
	arr[i] = arr[j]
	arr[j] = t
}
