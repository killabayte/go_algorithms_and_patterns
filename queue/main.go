package main

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(i int) {
}

func (q *Queue) Dequeue() int {
	return 0
}

func main() {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
}
