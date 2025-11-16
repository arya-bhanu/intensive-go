package main

import (
	"fmt"
	"sort"
)

func main() {
	denominations := []int{1, 5, 10, 25, 50}

	// Test amounts
	amounts := []int{87, 42, 99, 33, 7}

	for _, amount := range amounts {
		// Find minimum number of coins
		minCoins := MinCoins(amount, denominations)

		// Find coin combination
		coinCombo := CoinCombination(amount, denominations)

		// Print results
		fmt.Printf("Amount: %d cents\n", amount)
		fmt.Printf("Minimum coins needed: %d\n", minCoins)
		fmt.Printf("Coin combination: %v\n", coinCombo)
		fmt.Println("---------------------------")
	}
}

func MinCoins(amount int, denominations []int) int {
	counter := 0
	mapped := CoinCombination(amount, denominations)
	for _, val := range mapped {
		counter += val
	}

	return counter
}

func CoinCombination(amount int, denominations []int) map[int]int {
	maps := make(map[int]int)
	sort.Ints(denominations)
	i := len(denominations) - 1
	for {
		if amount == 0 || i < 0 {
			break
		}
		val := denominations[i]
		if val <= amount {
			amount -= val
			currCount, ok := maps[val]
			if ok {
				currCount++
				maps[val] = currCount
			} else {
				maps[val] = 1
			}
		} else {
			i--
		}
	}
	return maps
}
