package main

import (
	"fmt"
	"sync"
)

func main() {
	slice := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}

	wg := sync.WaitGroup{}
	wg.Add(len(slice))

	for i, val := range slice {
		go func(val []int, i int) {
			defer wg.Done()
			getSum(val, i)
		}(val, i)
	}
	wg.Wait()
}

func getSum(slice []int, i int) {
	sum := 0
	for _, val := range slice {
		sum += val
	}
	fmt.Printf("sum for slice %v:  %v\n", i+1, sum)
}
