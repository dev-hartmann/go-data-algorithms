package go_data_algorithms

import (
	"errors"
	"fmt"
	"math/rand"
)

func QuickSort(toBeSorted []int) ([]int, error) {

	fmt.Println("length of array:", len(toBeSorted))
	//check for nil or empty array
	if len(toBeSorted) < 1 && toBeSorted == nil {
		return nil, errors.New("input array must not be nil or empty")
	}

	//if array has only one value, return array
	if len(toBeSorted) < 2 {
		return toBeSorted, nil
	}

	left, right := 0, len(toBeSorted)-1
	pivot := rand.Int() % len(toBeSorted)

	toBeSorted[pivot], toBeSorted[right] = toBeSorted[right], toBeSorted[pivot]

	for i := range toBeSorted {
		if toBeSorted[i] < toBeSorted[right] {
			toBeSorted[i], toBeSorted[left] = toBeSorted[left], toBeSorted[i]
			left++
		}
	}

	toBeSorted[left], toBeSorted[right] = toBeSorted[right], toBeSorted[left]
	QuickSort(toBeSorted[:left])
	QuickSort(toBeSorted[left+1:])

	return toBeSorted, nil
}
