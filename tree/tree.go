package tree

import "errors"

// Tree is the base tree type in a binomial heap
type Tree struct {
	child, sibling, parent *Tree
	degree                 int
	value                  int
}

// Merge two binomial trees according to standard merge logic
// add the lower priority tree as a child of the higher priority tree
// and increment the degree
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
