package go_data_algorithms

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {

	expectedOutcomeArray := []int{1, 2, 3, 4, 5, 6}
	actualInputArray := []int{6, 5, 4, 3, 2, 1}

	outcomeArray := MergeSort(actualInputArray)

	fmt.Println(outcomeArray)
	fmt.Println(expectedOutcomeArray)

	if !testEq(outcomeArray, expectedOutcomeArray) {
		t.Error("expected result to be:", expectedOutcomeArray, "\n", "but was:", outcomeArray, "\n")
	}
}

func TestMergeSortNilInputArray(t *testing.T) {
	var actualInputArray []int

	outcomeArray := MergeSort(actualInputArray)

	if outcomeArray != nil {
		t.Error("outcome array should be nil")
	}
}

func TestMergeSortInputArrayWithOneElement(t *testing.T) {
	expectedOutcomeArray := []int{6}
	actualInputArray := []int{6}

	outcomeArray := MergeSort(actualInputArray)

	fmt.Println(outcomeArray)
	fmt.Println(expectedOutcomeArray)

	if !testEq(outcomeArray, expectedOutcomeArray) {
		t.Errorf("expected result to be:", expectedOutcomeArray, "but was:", outcomeArray)
	}
}
