package tree

import "errors"

// BinomialTree is the base tree type in a BinomialHeap
type BinomialTree struct {
	child, sibling, parent *BinomialTree
	degree                 int
	value                  int
}

// Merge two binomial trees according to standard merge algorithm
// 1. add the lower priority tree as a child of the higher priority tree
// 2. increment the degree
func Merge(p *BinomialTree, q *BinomialTree) (*BinomialTree, error) {
	if p.degree != q.degree {
		return &BinomialTree{}, errors.New("Merged trees must be of the same degree")
	}

	var parent, child *BinomialTree

	if p.value < q.value {
		parent = p
		child = q
	} else {
		parent = q
		child = p
	}

	child.sibling = parent.child
	child.parent = parent
	parent.child = child
	parent.degree = parent.degree + 1
	return parent, nil
}

// pops the head value from the tree
// returns it as a single node tree,
// and a slice of trees that were its children
func detatchHead(head *BinomialTree) (*BinomialTree, []*BinomialTree) {
	if head.degree == 0 {
		return head, nil
	}
	// numder of children is always the same as degree
	size := head.degree

	// move each of the children into the slice,
	// while setting their pointers to nil
	// since they will function as independent trees
	// add them in reverse order so they will
	// appear in increasing degree order in the resulting array
	children := make([]*BinomialTree, size)
	child := head.child
	i := size - 1
	for child != nil {
		cur := child
		child = cur.sibling
		cur.parent = nil
		cur.sibling = nil
		children[i] = cur
		i = i - 1
	}

	// nil out the head's pointers & set degree to 0
	head.child = nil
	head.sibling = nil
	head.degree = 0

	return head, children
}
