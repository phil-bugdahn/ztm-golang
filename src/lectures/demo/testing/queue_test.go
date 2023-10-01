package queue

import "testing"

func TestAddQueue(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		if len(q.items) != i {
			t.Errorf("Incorrect queue element count")
		}
		if !q.Append(i) {
			t.Errorf("Failed to append item %v", i)
		}
	}
	if q.Append(4) {
		t.Errorf("should not be able to add to a full queue")
	}
}

func TestNext(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		q.Append(i)
	}

	for i := 0; i < 3; i++ {
		item, ok := q.Next()
		if !ok {
			t.Errorf("Should be able to retrieve at idx %v", i)
		}
		if item != i {
			t.Errorf("Item at %v was actually %v", i, item)
		}
	}

	item, ok := q.Next()
	if ok {
		t.Errorf("should be empty, got %v", item)
	}
}