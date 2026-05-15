package job_queue

type Job struct {
	value    string
	priority int
	index    int
}

type JobQueue []*Job

func (q JobQueue) Len() int {
	return len(q)
}

func (q JobQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return q[i].priority > q[j].priority
}

func (q JobQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *JobQueue) Push(x any) {
	n := len(*q)
	item := x.(*Job)
	item.index = n
	*q = append(*q, item)
}

func (q *JobQueue) Pop() any {
	old := *q
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*q = old[0 : n-1]
	return item
}
