package queue

// An FIFO queue
type Queue []interface{}

// Pushes the element into the queue
func (q *Queue) Push(value interface{}) {
	*q = append(*q, value)
}

// Pops element from head.
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

//Returns if the queue is empty or not.
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
