package queue

import "testing"

func TestAddQueue(t *testing.T) {
	q := New(10)
	for i := 0; i < 10; i++ {
		if len(q.items) != i {
			t.Errorf("Expected len(q.items) to be %d, got %d", i, len(q.items))
		}
		if !q.Append(i) {
			t.Errorf("Expected q.Append(%d) to return true, got false", i)
		}
	}
}

func TestNextQueue(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		q.Append(i)
	}

	for i := 0; i < 3; i++ {
		item, ok := q.Next()
		if !ok {
			t.Errorf("Expected q.Next() to return true, got false")
		}
		if item != i {
			t.Errorf("Expected q.Next() to return %d, got %d", i, item)
		}
	}
	item, ok := q.Next()
	if ok {
		t.Errorf("should not be any mone more items in the queue %d", item)
	}
}
