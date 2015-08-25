package go_data_algorithms

import (
	"fmt"
	"testing"
)

func testEq(a, b []int) bool {
	if len(a) != len(b) {
		fmt.Println("expected:", len(a))
		fmt.Println("expected:", len(b))
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			fmt.Println("expected:", a[i])
			fmt.Println("expected:", b[i])
			return false
		}
	}

	return true
}

func TestInsertionSort(t *testing.T) {

	expectedOutcomeArray := []int{1, 2, 3, 4, 5, 6}
	actualInputArray := []int{6, 5, 4, 3, 2, 1}

	outcomeArray, err := InsertionSort(actualInputArray)
	if err != nil {
		t.Errorf("error should not occur")
	}

	fmt.Println(outcomeArray)
	fmt.Println(expectedOutcomeArray)

	if !testEq(outcomeArray, expectedOutcomeArray) {
		t.Error("expected result to be:", expectedOutcomeArray, "\n", "but was:", outcomeArray, "\n")
	}
}

func TestInsertionSortNilInputArray(t *testing.T) {
	var actualInputArray []int

	outcomeArray, err := InsertionSort(actualInputArray)

	if outcomeArray != nil {
		t.Log("outcome array should be nil")
	}

	if err != nil {
		t.Log("Error is:", err.Error())
	}
}

func TestInsertionSortInputArrayWithOneElement(t *testing.T) {
	expectedOutcomeArray := []int{6}
	actualInputArray := []int{6}

	outcomeArray, err := InsertionSort(actualInputArray)
	if err != nil {
		t.Errorf("error should not occur")
	}

	fmt.Println(outcomeArray)
	fmt.Println(expectedOutcomeArray)

	if !testEq(outcomeArray, expectedOutcomeArray) {
		t.Errorf("expected result to be:", expectedOutcomeArray, "\n", "but was:", outcomeArray, "\n")
	}
}
