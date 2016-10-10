package heap

import "github.com/ericmittelhammer/ginomialheap/tree"

type BinomialHeap struct {
	// head of the heap.  will be the tree with the lowest degree in the heap
	Head *tree.BinomialTree
	// shortcut pointer to the tree with the smallest head element.
	// ensures O(1) lookup
	Min *tree.BinomialTree
}

// utility method used to build the new heap
// there are two cases
// 1. h1 and h2 have the same degree.  merge them and add to the tail
//      if the tail is of the same degree, merge as well
// 2. h1 and h2 have different degrees. choose the smaller and add
//      to tail, unless tail is of the same degree, in which case merge
func AttachTo(tail *tree.BinomialTree, h1 *tree.BinomialTree, h2 *tree.BinomialTree) {
	if h1 == nil {
		tail.Sibling = h2
	} else if h2 == nil {
		tail.Sibling = h1
	} else if h1.Degree == h2.Degree {
		h1tail := h1.Sibling
		h2tail := h2.Sibling
		h1.Sibling = nil
		h2.Sibling = nil
		merged, _ := tree.Merge(h1, h2)
		if tail == nil {
			tail = merged
		} else if tail.Degree == merged.Degree {
			tail, _ = tree.Merge(tail, merged)
		} else {
			tail.Sibling = merged
		}
		AttachTo(tail, h1tail, h2tail)
	} else {
		var toAttach *tree.BinomialTree
		if h1.Degree < h2.Degree {
			toAttach = h1
			h1 = h1.Sibling
			toAttach.Sibling = nil
		} else {
			toAttach = h2
			h2 = h2.Sibling
			toAttach.Sibling = nil
		}
		if tail == nil {
			tail = toAttach
		} else if tail.Degree == toAttach.Degree {
			tail, _ = tree.Merge(tail, toAttach)
		} else {
			tail.Sibling = toAttach
		}
		AttachTo(tail, h1, h2)
	}
}

// merge two heaps according to the merge algorithm
// choose the lowest degree tree between the head of the two heaps
// merging trees when there are more than one of the same rank.
func Merge(heap1 *BinomialHeap, heap2 *BinomialHeap) *BinomialHeap {
	var min *tree.BinomialTree
	if heap1.Min.Value < heap2.Min.Value {
		min = heap1.Min
	} else {
		min = heap2.Min
	}
	result := &BinomialHeap{Head: nil, Min: min}
	AttachTo(result.Head, heap1.Head, heap2.Head)
	return result
}

func Create(initVal int) *BinomialHeap {
    t := &tree.BinomialTree{Value: initVal, Degree: 0}
    return &BinomialHeap{Head: t, Min: t}
}