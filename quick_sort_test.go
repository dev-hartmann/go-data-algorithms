package go_data_algorithms

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {

	expectedOutcomeArray := []int{1, 2, 3, 4, 5, 6}
	actualInputArray := []int{6, 5, 4, 3, 2, 1}

	outcomeArray, err := QuickSort(actualInputArray)
	if err != nil {
		t.Errorf("error should not occur")
	}

	fmt.Println(outcomeArray)
	fmt.Println(expectedOutcomeArray)

	if !testEq(outcomeArray, expectedOutcomeArray) {
		t.Error("expected result to be:", expectedOutcomeArray, "\n", "but was:", outcomeArray, "\n")
	}
}

func TestQuickSortNilInputArray(t *testing.T) {
	var actualInputArray []int

	outcomeArray, err := QuickSort(actualInputArray)

	if outcomeArray != nil {
		t.Log("outcome array should be nil")
	}

	if err != nil {
		t.Log("Error is:", err.Error())
	}
}

func TestQuickSortInputArrayWithOneElement(t *testing.T) {
	expectedOutcomeArray := []int{6}
	actualInputArray := []int{6}

	outcomeArray, err := QuickSort(actualInputArray)
	if err != nil {
		t.Errorf("error should not occur")
	}

	fmt.Println(outcomeArray)
	fmt.Println(expectedOutcomeArray)

	if !testEq(outcomeArray, expectedOutcomeArray) {
		t.Errorf("expected result to be:", expectedOutcomeArray, "but was:", outcomeArray)
	}
}
