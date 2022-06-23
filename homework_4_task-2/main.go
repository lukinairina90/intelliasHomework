package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := "2000 1 9 3 1000 4 -5 -100 -400 300 1 6"
	var result string
	var inputNum []int

	inputSlice := strings.Split(input, " ")

	for _, val := range inputSlice {
		convStr, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		inputNum = append(inputNum, convStr)
	}

	sort.Ints(inputNum)

	//fmt.Println("inputNum = ", inputNum)

	if len(inputNum) == 1 {
		result = strconv.Itoa(inputNum[0])
	} else {
		result = strconv.Itoa(inputNum[len(inputNum)-1]) + " " + strconv.Itoa(inputNum[0])
	}

	fmt.Printf("result = %q\n", result)

}
