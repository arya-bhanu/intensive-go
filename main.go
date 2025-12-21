package main

import (
	"strings"
	"time"
)

// SlowSort sorts a slice of integers using a very inefficient algorithm (bubble sort)
// TODO: Optimize this function to be more efficient
func SlowSort(data []int) []int {
	// Make a copy to avoid modifying the original
	result := make([]int, len(data))
	copy(result, data)

	// Bubble sort implementation
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result)-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}

	return result
}

// OptimizedSort is your optimized version of SlowSort
// It should produce identical results but perform better
func OptimizedSort(data []int) []int {
	// TODO: Implement a more efficient sorting algorithm
	// Hint: Consider using sort package or a more efficient algorithm
	return quickSort(data, 0, len(data)-1) // Replace this with your optimized implementation
}

func quickSort(data []int, low, high int) []int {
	if low < high {
		// sorted partition
		arr, pivot := partition(data, low, high)

		// sorted left partition
		data = quickSort(arr, low, pivot-1)

		// sorted right partition
		data = quickSort(data, pivot+1, high)
	}
	return data
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] > pivot {
			continue
		}
		arr[j], arr[i] = arr[i], arr[j]
		i++
	}
	// locate the pivot
	arr[high], arr[i] = arr[i], arr[high]
	return arr, i
}

// InefficientStringBuilder builds a string by repeatedly concatenating
// TODO: Optimize this function to be more efficient
func InefficientStringBuilder(parts []string, repeatCount int) string {
	result := ""

	for i := 0; i < repeatCount; i++ {
		for _, part := range parts {
			result += part
		}
	}

	return result
}

// OptimizedStringBuilder is your optimized version of InefficientStringBuilder
// It should produce identical results but perform better
func OptimizedStringBuilder(parts []string, repeatCount int) string {
	var sb strings.Builder
	for _, part := range parts {
		sb.WriteString(part)
	}

	str := sb.String()
	sb.Reset()

	for range repeatCount {
		sb.WriteString(str)
	}
	// TODO: Implement a more efficient string building method
	// Hint: Consider using strings.Builder or bytes.Buffer
	return sb.String() // Replace this with your optimized implementation
}

// ExpensiveCalculation performs a computation with redundant work
// It computes the sum of all fibonacci numbers up to n
// TODO: Optimize this function to be more efficient
func ExpensiveCalculation(n int) int {
	if n <= 0 {
		return 0
	}

	sum := 0
	for i := 1; i <= n; i++ {
		sum += fibonacci(i)
	}

	return sum
}

// Helper function that computes the fibonacci number at position n
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	calculated := fibonacci(n-1) + fibonacci(n-2)
	return calculated
}

func optfibonacci(n int, c map[int]int) int {
	if n <= 1 {
		return n
	}
	if val, ok := c[n]; ok {
		return val
	}
	calculated := optfibonacci(n-1, c) + optfibonacci(n-2, c)
	c[n] = calculated
	return calculated
}

// OptimizedCalculation is your optimized version of ExpensiveCalculation
// It should produce identical results but perform better
func OptimizedCalculation(n int) int {
	// TODO: Implement a more efficient calculation method
	// Hint: Consider memoization or avoiding redundant calculations
	if n <= 0 {
		return 0
	}

	sum := 0
	cache := make(map[int]int)
	for i := 1; i <= n; i++ {
		sum += optfibonacci(i, cache)
	}

	return sum // Replace this with your optimized implementation
}

// HighAllocationSearch searches for all occurrences of a substring and creates a map with their positions
// TODO: Optimize this function to reduce allocations
func HighAllocationSearch(text, substr string) map[int]string {
	result := make(map[int]string)

	// Convert to lowercase for case-insensitive search
	lowerText := strings.ToLower(text)
	lowerSubstr := strings.ToLower(substr)

	for i := 0; i < len(lowerText); i++ {
		// Check if we can fit the substring starting at position i
		if i+len(lowerSubstr) <= len(lowerText) {
			// Extract the potential match
			potentialMatch := lowerText[i : i+len(lowerSubstr)]

			// Check if it matches
			if potentialMatch == lowerSubstr {
				// Store the original case version
				result[i] = text[i : i+len(substr)]
			}
		}
	}

	return result
}

// OptimizedSearch is your optimized version of HighAllocationSearch
// It should produce identical results but perform better with fewer allocations
func OptimizedSearch(text, substr string) map[int]string {
	result := make(map[int]string)

	if substr == "" {
		return result
	}

	loweredText := strings.ToLower(text)
	loweredSubstr := strings.ToLower(substr)

	if len(loweredSubstr) > len(loweredText) {
		return result
	}

	for i := 0; i < len(loweredText); i++ {
		if loweredSubstr[0] == loweredText[i] {
			if len(text)-i >= len(substr) {
				match := loweredText[i : i+len(loweredSubstr)]
				if match == loweredSubstr {
					result[i] = text[i : i+len(loweredSubstr)]
					i = i + len(loweredSubstr) - 1
				}
			}
		}
	}

	return result
}

// A function to simulate CPU-intensive work for benchmarking
// You don't need to optimize this; it's just used for testing
func SimulateCPUWork(duration time.Duration) {
	start := time.Now()
	for time.Since(start) < duration {
		// Just waste CPU cycles
		for i := 0; i < 1000000; i++ {
			_ = i
		}
	}
}
