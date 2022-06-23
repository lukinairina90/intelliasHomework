package main

import "fmt"

func main() {
	arr := []int{4, 1, 4, -4, 6, 3, 8, 8, 1, 1, 1, 1, 1, 1, 4, 5, 6, 0, 0, 0, 4, 3, 3, 33}
	lenght := len(arr)
	var result []int

	mapa := make(map[int]int)

	for i := 0; i < lenght; i++ {
		mapa[arr[i]] = i
	}

	for i, _ := range mapa {
		result = append(result, i)
	}
	fmt.Println(mapa)
	fmt.Println(result)
}
