package heap

import "errors"

type Heap interface {
	Insert(int)
	Extract() (int, error)
}

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(val int) {
	// Append value to the back
	h.array = append(h.array, val)

	// Heapify it up to its new position
	h.maxHeapifyUp(len(h.array)-1)
}

func (h *MaxHeap) Extract() (int, error) {
	if len(h.array) == 0 {
		return -1, errors.New("cannot extract empty heap")
	}

	extracted := h.array[0]

	lastIndex := len(h.array) - 1

	// Bring last element to front
	h.array[0] = h.array[lastIndex]

	// Downsize array by 1
	h.array = h.array[:lastIndex]

	h.maxHeapifyDown(0)

	return extracted, nil
}

func (h *MaxHeap) maxHeapifyUp(index int) {
	// Loop till false: Compare if parent is smaller than child
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	// loop while index has at least one child
	for l <= lastIndex {
		// only left child present, no right child
		if l == lastIndex {
			childToCompare = l
			// left child bigger than right child
		} else if h.array[l] > h.array[r] {
			childToCompare = l
			// right child bigger than left child
		} else {
			childToCompare = r
		}

		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			break
		}
	}
}

func parent(i int) int {
	return (i-1) / 2
}

func left(i int) int {
	return (i*2) + 1
}

func right(i int) int {
	return (i*2) + 2
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}