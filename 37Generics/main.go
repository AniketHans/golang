package main

import "fmt"

type Number interface {
	int | int32 | int64 | float32 | float64
}

func NumberSum[T Number](values []T) T {
	var s T;
	for _, val := range values {
		s += val
	}
	return s
}

func main() {
	numbers1 := []int{1, 2, 3, 4, 5}
	numbers2 := []int32{1, 2, 3, 4, 5}
	numbers3 := []int64{1, 2, 3, 4, 5}
	numbers4 := []float32{1.0, 2.0, 3.0, 4.0, 5.0}
	numbers5 := []float64{1.0, 2.0, 3.0, 4.0, 5.0}

	fmt.Println(NumberSum(numbers1))
	fmt.Println(NumberSum(numbers2))
	fmt.Println(NumberSum(numbers3))
	fmt.Println(NumberSum(numbers4))
	fmt.Println(NumberSum(numbers5))
}