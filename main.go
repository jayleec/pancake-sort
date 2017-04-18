//#Pancake sorting golang version.
//#Date : April 18, 2017
//#Written by Lee Jae Kyung
//#Golang 1.7.4


package main

import "fmt"

func main(){
	// sort from back ( unsorted part|sorted part  )
	// 1) find the biggest number
	// 2) put turner behind
	//  (Max, unsorted part | sorted part)
	// 3) put turner previous of sorted part
	//  (unsorted part, Max | sorted part)

	S := []int{5,6,3,4,7,8,1}
	max := 0
	for i := len(S)-1; i > 0; i--{
		max = findMax(S, i)
		reverse(S, max)
		reverse(S, i)
	}
	fmt.Println(S)
}

//return max number's index
func findMax(A []int, high int) int{
	max, index := A[0], 0
	for i := 1; i<= high; i++{
		if max < A[i]{
			max = A[i]
			index = i
		}
	}
	return index
}

func reverse(A []int, turner int){
	for i, j := 0, turner; i<turner+1/2; i,j = i+1, j-1{
		A[i], A[j] = A[j], A[i]
	}
}