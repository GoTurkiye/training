package main

type MyItem int

type Queue struct {
	list []MyItem
}

func NewQueue() *Queue {
	return &Queue{
		list: []MyItem{},
	}
}

func (q *Queue) Enqueue(value MyItem) {
	q.list = append(q.list, value)
}

func (q *Queue) Dequeue() MyItem {
	first := q.list[0]
	q.list = q.list[1:]
	return first
}

func main() {
	q := NewQueue()
	q.Enqueue(3)
	q.Enqueue(5)
	q.Enqueue(7)
	// 7 5 3
	println(q.Dequeue())
}