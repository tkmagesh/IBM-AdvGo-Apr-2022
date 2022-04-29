package main

import "fmt"

func sumInts(nos []int) int {
	result := 0
	for _, no := range nos {
		result += no
	}
	return result
}

func sumFloats(nos []float32) float32 {
	result := float32(0)
	for _, no := range nos {
		result += no
	}
	return result
}

/*
func sum[T int | float32](nos []T) T {
	var result T
	for _, no := range nos {
		result += no
	}
	return result
}
*/

type Numbers interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func sum[T Numbers](nos []T) T {
	var result T
	for _, no := range nos {
		result += no
	}
	return result
}

func Filter[T Numbers](items []T, predicate func(item T) bool) []T {
	result := []T{}
	for _, p := range items {
		if predicate(p) {
			result = append(result, p)
		}
	}
	return result
}

func main() {
	ints := []int{3, 1, 4, 2, 5}
	//fmt.Println(sumInts(ints))
	fmt.Println(sum(ints))

	floats := []float32{3, 1, 4, 2, 5}
	//fmt.Println(sumFloats(floats))
	fmt.Println(sum(floats))

	evenNos := Filter(ints, func(no int) bool {
		return no%2 == 0
	})
	fmt.Println(evenNos)
}
