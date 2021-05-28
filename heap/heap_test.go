package heap

import (
	"testing"
)

func TestInsert(t *testing.T) {
	maxHeap := MaxHeap{}
	maxHeap.Insert(10)
	maxHeap.Insert(30)
	maxHeap.Insert(20)
	maxHeap.Insert(40)

	if maxHeap.array[0] != 40 {
		t.Error("heapifyUp process inaccurate")
	}
}

func TestExtract(t *testing.T) {
	maxHeap := MaxHeap{}
	buildHeap := []int{20, 10, 30, 15, 13, 7, 9}

	for _, v := range buildHeap {
		maxHeap.Insert(v)
	}

	extracted, _ := maxHeap.Extract()
	if extracted != 30 {
		t.Errorf("expected largest int, but got %d instead", extracted)
	}
	if maxHeap.array[0] != 20 {
		t.Errorf("expected new max to be 20, but got %d", maxHeap.array[0])
	}

	extracted, _ = maxHeap.Extract()
	if extracted != 20 {
		t.Errorf("expected largest int, but got %d instead", extracted)
	}
	if maxHeap.array[0] != 15 {
		t.Errorf("expected new max to be 20, but got %d", maxHeap.array[0])
	}
}

func TestEmptyExtract(t *testing.T) {
	maxHeap := MaxHeap{}

	_, err := maxHeap.Extract()

	if err == nil {
		t.Error("expected error when extracting empty heap but got none")
	}
}