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
