package queue

import "testing"

func TestAddQueue(t *testing.T) {
	q := New(3)

	for i := 0; i < 3; i++ {
		if len(q.items) != i {
			t.Errorf("The queue has an incorrect element count: %v, want %v", len(q.items), i)
		}
		if !q.Append(i) {
			t.Errorf("The queue has a failed append attempt: item %v", i)
		}
	}
	if q.Append(4) {
		t.Errorf("The queue should not be able to add to a full queue")
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
			t.Errorf("The queue should have returned a next item to iterate")
		}
		if item != i {
			t.Errorf("The item should be in a correct order - got: %v, want: %v", item, i)
		}
	}
}
