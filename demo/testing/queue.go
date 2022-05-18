package queue

type Queue struct {
	items    []int
	capacity int
}

func New(Capacity int) *Queue {
	return &Queue{
		items:    make([]int, 0, Capacity),
		capacity: Capacity,
	}
}

func (q *Queue) IsFull() bool {
	return len(q.items) == q.capacity
}

func (q *Queue) Append(item int) bool {
	if q.IsFull() {
		return false
	}
	q.items = append(q.items, item)
	return true
}

func (q *Queue) Next() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}
