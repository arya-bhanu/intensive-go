package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello World")
	arr := make([]int, 0)
	var arr2 []int
	fmt.Println(reflect.DeepEqual(arr, []int{}))
	fmt.Println(reflect.DeepEqual(arr2, []int{}))
	fmt.Println(arr == nil)
	fmt.Println(arr2 == nil)

	var testarr = make([]int, 0, 6)
	testarr = append(testarr, 5)
	testarr = append(testarr, 5)
	testarr = append(testarr, 5)
	testarr = append(testarr, 5)
	testarr = append(testarr, 5)
	testarr = append(testarr, 5)
	testarr = append(testarr, 5)
	fmt.Println(len(testarr))
	fmt.Println(cap(testarr))
}

func FindMax(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	max := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > max {
			max = numbers[i]
		}
	}
	return max
}

func RemoveDuplicates(numbers []int) []int {
	singular := make([]int, 0)
	hash := make(map[int]struct{})
	for _, num := range numbers {
		if _, ok := hash[num]; ok {
			continue
		}
		singular = append(singular, num)
		hash[num] = struct{}{}
	}
	return singular
}

func ReverseSlice(slice []int) []int {
	reversed := make([]int, 0)
	for i := len(slice) - 1; i >= 0; i-- {
		reversed = append(reversed, slice[i])
	}
	return reversed
}

func FilterEven(numbers []int) []int {
	filtered := make([]int, 0)
	for _, num := range numbers {
		if num%2 == 0 {
			filtered = append(filtered, num)
		}
	}
	return filtered
}
