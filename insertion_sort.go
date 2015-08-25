package go_data_algorithms

import (
	"errors"
)

func InsertionSort(toBeSorted []int) ([]int, error) {
	//check for nil or empty array
	if len(toBeSorted) < 1 && toBeSorted == nil {
		return nil, errors.New("input array must not be nil or empty")
	}

	//if array has only one value, return array
	if len(toBeSorted) == 1 {
		return toBeSorted, nil
	}

	//loop over input array
	for index, value := range toBeSorted {
		//loop backwards from index position to get element n-1 and compare n and n-1
		for y := index; y > -1; y-- {
			//swap n with n-1 if n-1 < n
			if value < toBeSorted[y] {
				toBeSorted[y], toBeSorted[y+1] = toBeSorted[y+1], toBeSorted[y]
			}
		}
	}

	return toBeSorted, nil
}
