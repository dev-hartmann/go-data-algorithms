package main

import (
	"fmt"
	da "go-data-algorithms"
)

func main() {
	expectedOutcomeArray := []int{1, 2, 3, 4, 5, 6}
	actualInputArray := []int{6, 5, 4, 3, 2, 1}

	outcomeArray, err := da.InsertionSort(actualInputArray)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(outcomeArray)
	fmt.Println(expectedOutcomeArray)

	outcome2, err := da.QuickSort(actualInputArray)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(outcome2)
}
