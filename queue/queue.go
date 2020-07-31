package queue

type Queue []interface{}

func (q *Queue) Push(value interface{}) {
	*q = append(*q, value)
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
