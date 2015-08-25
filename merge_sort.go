package go_data_algorithms


func MergeSort(toBeSorted []int) []int {

	if len(toBeSorted) < 1 && toBeSorted == nil {
		return nil
	}

	//if array has only one value, return array
	if len(toBeSorted) < 2 {
		return toBeSorted
	}

		middle := (len(toBeSorted)) / 2
		return merge(MergeSort(toBeSorted[:middle]),MergeSort(toBeSorted[middle:]))
}

func merge(leftArray,rightArray []int) []int {

	i,j := 0,0
	size := len(leftArray) + len(rightArray)
	helper := make([]int, size)

	for k := 0; k < size; k++ {
		if i > len(leftArray)-1 && j <= len(rightArray)-1 {
			helper[k] = rightArray[j]
			j++
		} else if j > len(rightArray)-1 && i <= len(leftArray)-1 {
			helper[k] = leftArray[i]
			i++
		} else if leftArray[i] < rightArray[j] {
			helper[k] = leftArray[i]
			i++
		} else {
			helper[k] = rightArray[j]
			j++
		}
	}
	return helper
}