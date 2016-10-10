package heap

import "testing"

func TestCreate(t *testing.T) {
	heap := Create()

	if heap.Head != nil {
		t.Error("Heap Head incorrectly set")
	}

	if heap.Min != nil {
		t.Error("Heap Min incorrectly set")
	}
}
