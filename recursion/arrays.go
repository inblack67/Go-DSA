package recursion

import "fmt"

// default index=0
func PrintArray(arr []int, index int) {
	if index == len(arr) {
		return
	}
	fmt.Println(arr[index])
	PrintArray(arr, index+1)
}

// default index=len(arr) - 1
func PrintArrayInReverse(arr []int, index int) {
	if index == -1 {
		return
	}
	fmt.Println(arr[index])
	PrintArrayInReverse(arr, index-1)
}

// default index=0
func MaxOfArray(arr []int, index int) int {
	if index == 0 {
		return arr[0]
	}
	curr := arr[index-1]
	prev := MaxOfArray(arr, index-1)
	if curr > prev {
		return curr
	} else {
		return prev
	}
}
