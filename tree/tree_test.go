package tree

import "testing"

func TestCreation(t *testing.T) {
	a := &Tree{value: 55, degree: 0}
	if a.value != 55 {
		t.Error("incorrect inital value set")
	}
	if a.degree != 0 {
		t.Error("incorrect initial degree set")
	}
}

func TestMerge(t *testing.T) {
	a := &Tree{value: 1}
	b := &Tree{value: 2}
	merged, err := Merge(a, b)
	if err != nil {
		t.Error("error when merging", err)
	}
	if merged.value != 1 {
		t.Error("expected min value at head after merge")
	}
	if merged != a {
		t.Error("incorrect merge")
	}
	if merged.child.value != 2 {
		t.Error("expected lower priority child value after merge")
	}
	if merged.child != b {
		t.Error("incorrect merge")
	}
	if merged.child.parent != a {
		t.Error("incorrect parent pointer")
	}
}

func TestMergeFail(t *testing.T) {
	a := &Tree{value: 1, degree: 0}
	b := &Tree{value: 2, degree: 1}

	_, err2 := Merge(a, b)
	if err2 == nil {
		t.Error("Error expected. should not be able to merge trees of different degrees.")
	}
}

func TestMergeWithChildren(t *testing.T) {
	a := &Tree{value: 1}
	b := &Tree{value: 2}
	c, _ := Merge(a, b)

	x := &Tree{value: 3}
	y := &Tree{value: 4}
	z, _ := Merge(x, y)

	head, _ := Merge(z, c)

	if head != a {
		t.Error("priority not preserved in merge")
	}

	if !(head.child == x && head.child.child == y) {
		t.Error("children not set correctly in merge")
	}

	if !(head.child.sibling == b && head.child.sibling.sibling == nil) {
		t.Error("siblings not set correctly during merge")
	}

	if !(head.child.child.parent == head.child && head.child.parent == head && head.child.sibling.parent == head) {
		t.Error("parents not set correclty during merge")
	}
}

func TestPopMin(t *testing.T) {
	a := &Tree{value: 1}
	b := &Tree{value: 2}
	c, _ := Merge(a, b)

	d := &Tree{value: 3}
	e := &Tree{value: 4}
	f, _ := Merge(d, e)

	g, _ := Merge(c, f)

	s := &Tree{value: 5}
	u := &Tree{value: 6}
	v, _ := Merge(s, u)
	w := &Tree{value: 7}
	x := &Tree{value: 8}
	y, _ := Merge(w, x)

	z, _ := Merge(v, y)

	tree, _ := Merge(g, z)

	min, children := popMin(tree)

	if min.value != 1 {
		t.Error("min val was not popped")
	}

	if !(min.child == nil && min.sibling == nil) {
		t.Error("head pointers were not nilled correctly")
	}

	if min.degree != 0 {
		t.Error("head degree not set to 0")
	}

	if len(children) != 3 {
		t.Error("wrong number of children returned")
	}

	i := 0
	for i < len(children) {
		child := children[i]
		if child.degree != i {
			t.Error("children not added in degree order")
		}
		if !(child.parent == nil && child.sibling == nil) {
			t.Error("children pointers not correctly set to nil")
		}
		i = i + 1
	}

}
