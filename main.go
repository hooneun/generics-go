package main

import (
	"fmt"
)

type Integer interface {
	int8 | int16 | int32 | int64 | int
}

type Float interface {
	float32 | float64
}

func Add[T Integer | Float](a, b T) T {
	return a + b
}

func Filter[T any](slice []T, f func(T) bool) []T {
	var n []T
	for _, e := range slice {
		if f(e) {
			n = append(n, e)
		}
	}
	return n
}

func Map[T any](slice []T, f func(T) T) []T {
	var n []T

	for _, e := range slice {
		n = append(n, f(e))
	}

	return n
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	filterNumber := Filter(numbers, func(v int) bool {
		return v%2 == 0
	})

	fmt.Println(filterNumber)

	fmt.Println(Add(1, 2))
	fmt.Println(Add(1.2, 4.5))

	mapNumber := Map(numbers, func(v int) int {
		return v * 2
	})

	fmt.Println(mapNumber)
}
