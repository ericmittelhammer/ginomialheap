package node

import "errors"

// Node is the root of a Binomial Tree
type Node struct {
	FirstChild *Node // Only need a pointer to the first child. Traversal always starts with this Node
	Next       *Node // Sibling Node.  Either as part of the LinkedList in top-level heap or as part of a subtree
	//Parent     *Node //
	Degree int
	Value  int
}

// Merge two binomial trees according to standard merge algorithm
// 1. add the lower priority tree as a child of the higher priority tree
// 2. increment the degree
// Returns the node that is now at the head of this tree
func Merge(p *Node, q *Node) (*Node, error) {
	if p.Degree != q.Degree {
		return nil, errors.New("Merged trees must be of the same degree")
	}

	var parent, child *Node

	if p.Value < q.Value {
		parent = p
		child = q
	} else {
		parent = q
		child = p
	}

	child.Next = parent.FirstChild
	//child.Parent = parent
	parent.FirstChild = child
	parent.Degree = parent.Degree + 1
	return parent, nil
}

// pops the head value from the tree
// sets its degree to 0 and nils its child
// returns its first child (which is now itself a Binomial Heap)
func (head *Node) detatchHead() (*Node, error) {
	if head.Next != nil {
		return nil, errors.New("head.Next != nil.  This node is either in a heap or is a subtree.  Can only detach top-level trees after they have been removed from the heap")
	}
	if head.Degree == 0 {
		return head, nil
	}

	child := head.FirstChild
	head.Degree = 0
	head.FirstChild = nil

	return child, nil
}
