package gobasics

import "fmt"

func TestSliceArray() {
	sliceA := make([]int, 5)
	fmt.Println(len(sliceA))
	fmt.Println(sliceA)
	fmt.Println(cap(sliceA))
}
