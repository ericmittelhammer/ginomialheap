package heap

import "testing"

func TestCreate(t *testing.T) {
	heap := Create(5)

	if heap.Head.Value != 5 {
		t.Error("Heap Head incorrectly set")
	}

	if heap.Min != heap.Head {
		t.Error("Heap Min incorrectly set")
	}
}
