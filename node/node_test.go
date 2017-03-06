package node

import (
	"log"
	"math/rand"
	"testing"
)

func MakeNode(degree int) *Node {

	if degree == 0 {
		val := rand.Intn(10000)
		return &Node{Value: val, Degree: 0}
	}

	node1 := MakeNode(degree - 1)
	node2 := MakeNode(degree - 1)
	res, _ := Merge(node1, node2)
	return res
}

func TestCreation(t *testing.T) {
	a := &Node{Value: 55, Degree: 0}
	if a.Value != 55 {
		t.Error("incorrect inital value set")
	}
	if a.Degree != 0 {
		t.Error("incorrect initial degree set")
	}
}

func TestMerge(t *testing.T) {
	a := &Node{Value: 1}
	b := &Node{Value: 2}
	merged, err := Merge(a, b)
	if err != nil {
		t.Error("error when merging", err)
	}
	if merged.Value != 1 {
		t.Error("expected min value at head after merge")
	}
	if merged != a {
		t.Error("incorrect merge")
	}
	if merged.FirstChild.Value != 2 {
		t.Error("expected lower priority child value after merge")
	}
	if merged.FirstChild != b {
		t.Error("incorrect merge")
	}
	// if merged.Child.Parent != a {
	// 	t.Error("incorrect parent pointer")
	// }
}

func TestMergeEquiv(t *testing.T) {
	a := &Node{Value: 2}
	b := &Node{Value: 2}
	merged, _ := Merge(a, b)

	if merged != a {
		t.Error("Arguemt ordering not preserved")
	}

	if merged.FirstChild != b {
		t.Error("Arguemt ordering not preserved")
	}
	// if merged.Child.Parent != a {
	// 	t.Error("incorrect parent pointer")
	// }
}

func TestMergeFail(t *testing.T) {
	a := &Node{Value: 1, Degree: 0}
	b := &Node{Value: 2, Degree: 1}

	_, err2 := Merge(a, b)
	if err2 == nil {
		t.Error("Error expected. should not be able to merge trees of different degrees.")
	}
}

func TestMergeWithChildren(t *testing.T) {
	a := &Node{Value: 1}
	b := &Node{Value: 2}
	c, _ := Merge(a, b)

	x := &Node{Value: 3}
	y := &Node{Value: 4}
	z, _ := Merge(x, y)

	head, _ := Merge(z, c)

	if head != a {
		t.Error("priority not preserved in merge")
	}

	if !(head.FirstChild == b) {
		t.Error("children not set correctly in merge")
	}

	if !(head.FirstChild.Next == x) {
		t.Error("child tree not added after higest order")
	}

	// if !(head.Child.Child.Parent == head.Child && head.Child.Parent == head && head.Child.Sibling.Parent == head) {
	// 	t.Error("parents not set correclty during merge")
	// }
}

func TestDetatchHead(t *testing.T) {
	a := &Node{Value: 1}
	b := &Node{Value: 2}
	c, _ := Merge(a, b)

	d := &Node{Value: 3}
	e := &Node{Value: 4}
	f, _ := Merge(d, e)

	g, _ := Merge(c, f)

	s := &Node{Value: 5}
	u := &Node{Value: 6}
	v, _ := Merge(s, u)
	w := &Node{Value: 7}
	x := &Node{Value: 8}
	y, _ := Merge(w, x)

	z, _ := Merge(v, y)

	tree, _ := Merge(g, z)

	firstChild, _ := tree.detatchHead()

	if tree.FirstChild != nil {
		t.Error("head was not detached from child")
	}

	if tree.Degree != 0 {
		t.Error("head degree not set to 0")
	}

	if firstChild != b {
		t.Error("did not return FirstChild")
	}

}
func isTree(n *Node) {
	children, err := n.detatchHead()
	if err != nil {
		log.Fatal(err)
	}
	iter := children
	for iter != nil {
		cur := iter
		if n.Value > cur.Value {
			log.Fatal("parent value ", n.Value, "greater than child value", cur.Value)
		}
		iter = cur.Next
		cur.Next = nil
		isTree(cur)
	}
}

func TestTreeSpec(t *testing.T) {
	for i := 0; i < 100; i++ {
		degree := rand.Intn(10) // test trees up to degree 10 (1024 nodes)
		n := MakeNode(degree)
		isTree(n)
	}
}

func TestUnion(t *testing.T) {
	// create two heaps.

	// 0 -> 3
	zero := &Node{Value: 5}

	one, _ := Merge(&Node{Value: 10}, &Node{Value: 11})

	twoInner1, _ := Merge(&Node{Value: 15}, &Node{Value: 16})
	twoInner2, _ := Merge(&Node{Value: 17}, &Node{Value: 18})
	two, err := Merge(twoInner1, twoInner2)
	if err != nil {
		t.Log(err)
	}

	// three := &Node{Value: 5}
	// for i := 0; i < 7; i++ {
	// 	val := rand.Int() + 5
	// 	three, _ = Merge(three, &Node{Value: val})
	// }

	zero.Next = two

	//one.Next = three

	result := Union(zero, one)

	count := 0
	lastDegree := -1

	cur := result
	for cur != nil {
		if cur.Degree <= lastDegree {
			t.Error("Degrees out of order")
		}
		count = count + 1
		lastDegree = cur.Degree
		cur = cur.Next
	}

	if count != 3 {
		t.Error("Expected 3 nodes to be in result. Instead there were", count, result)
	}
}
