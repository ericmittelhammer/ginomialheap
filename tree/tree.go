package tree

import "errors"

// Tree is the base tree type in a binomial heap
type Tree struct {
	child, sibling, parent *Tree
	degree                 int
	value                  int
}

// Merge two binomial trees according to standard merge algorithm
// 1. add the lower priority tree as a child of the higher priority tree
// 2. increment the degree
func Merge(p *Tree, q *Tree) (*Tree, error) {
	if p.degree != q.degree {
		return &Tree{}, errors.New("Merged trees must be of the same degree")
	}

	var parent, child *Tree

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

// pops the min (head) value from the tree
// returns the head as a single node tree,
// and a slice of trees that were its children
func popMin(head *Tree) (*Tree, []*Tree) {
	if head.degree == 0 {
		return head, nil
	}
	// the number of children that a node has is always the same as its degree
	size := head.degree

	// move each of the children into the slice,
	// while setting their pointers to nil
	// since they will funciton as independent trees
	// add them in reverse order so they will
	// appear in increasing degree order in the resulting array
	children := make([]*Tree, size)
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
