package tree

import "testing"

func TestCreation(t *testing.T) {
	a := &BinomialTree{Value: 55, Degree: 0}
	if a.Value != 55 {
		t.Error("incorrect inital value set")
	}
	if a.Degree != 0 {
		t.Error("incorrect initial degree set")
	}
}

func TestMerge(t *testing.T) {
	a := &BinomialTree{Value: 1}
	b := &BinomialTree{Value: 2}
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
	if merged.Child.Value != 2 {
		t.Error("expected lower priority child value after merge")
	}
	if merged.Child != b {
		t.Error("incorrect merge")
	}
	if merged.Child.Parent != a {
		t.Error("incorrect parent pointer")
	}
}

func TestMergeFail(t *testing.T) {
	a := &BinomialTree{Value: 1, Degree: 0}
	b := &BinomialTree{Value: 2, Degree: 1}

	_, err2 := Merge(a, b)
	if err2 == nil {
		t.Error("Error expected. should not be able to merge trees of different degrees.")
	}
}

func TestMergeWithChildren(t *testing.T) {
	a := &BinomialTree{Value: 1}
	b := &BinomialTree{Value: 2}
	c, _ := Merge(a, b)

	x := &BinomialTree{Value: 3}
	y := &BinomialTree{Value: 4}
	z, _ := Merge(x, y)

	head, _ := Merge(z, c)

	if head != a {
		t.Error("priority not preserved in merge")
	}

	if !(head.Child == x && head.Child.Child == y) {
		t.Error("children not set correctly in merge")
	}

	if !(head.Child.Sibling == b && head.Child.Sibling.Sibling == nil) {
		t.Error("siblings not set correctly during merge")
	}

	if !(head.Child.Child.Parent == head.Child && head.Child.Parent == head && head.Child.Sibling.Parent == head) {
		t.Error("parents not set correclty during merge")
	}
}

func TestDetatchHead(t *testing.T) {
	a := &BinomialTree{Value: 1}
	b := &BinomialTree{Value: 2}
	c, _ := Merge(a, b)

	d := &BinomialTree{Value: 3}
	e := &BinomialTree{Value: 4}
	f, _ := Merge(d, e)

	g, _ := Merge(c, f)

	s := &BinomialTree{Value: 5}
	u := &BinomialTree{Value: 6}
	v, _ := Merge(s, u)
	w := &BinomialTree{Value: 7}
	x := &BinomialTree{Value: 8}
	y, _ := Merge(w, x)

	z, _ := Merge(v, y)

	tree, _ := Merge(g, z)

	min, children := detatchHead(tree)

	if min.Value != 1 {
		t.Error("min val was not popped")
	}

	if !(min.Child == nil && min.Sibling == nil) {
		t.Error("head pointers were not nilled correctly")
	}

	if min.Degree != 0 {
		t.Error("head degree not set to 0")
	}

	if len(children) != 3 {
		t.Error("wrong number of children returned")
	}

	i := 0
	for i < len(children) {
		child := children[i]
		if child.Degree != i {
			t.Error("children not added in degree order")
		}
		if !(child.Parent == nil && child.Sibling == nil) {
			t.Error("children pointers not correctly set to nil")
		}
		i = i + 1
	}

}
