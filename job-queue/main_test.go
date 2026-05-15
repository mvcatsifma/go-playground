package job_queue

import (
	"container/heap"
	"testing"
)

// TODO: add benchmark test
// TODO: add thread safety test
func TestJobQueue_Len(t *testing.T) {
	tests := []struct {
		name string
		j    JobQueue
		want int
	}{
		{"len", make(JobQueue, 0), 0},
		{"len", make(JobQueue, 1), 1},
		{"len", make(JobQueue, 9), 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJobQueue_Push(t *testing.T) {
	jq := make(JobQueue, 0)
	heap.Push(&jq, &Job{})

	if jq.Len() != 1 {
		t.Error("expected len to be 1")
		t.Fail()
	}
}

func TestJobQueue_Less(t *testing.T) {
	jq := FixtureJobQueue()
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		q    JobQueue
		args args
		want bool
	}{
		{"less", jq, args{0, 1}, true},
		{"more", jq, args{1, 0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJobQueue_Swap(t *testing.T) {
	jq := FixtureJobQueue()
	jq.Swap(0, 1)
	j := jq.Pop().(*Job)
	if j.value != "/foo" {
		t.Error("expected value to be '/foo'")
		t.Fail()
	}
}

func TestJobQueue_Pop(t *testing.T) {
	jq := FixtureJobQueue()
	heap.Push(&jq, &Job{value: "/bax", priority: 5})
	j := heap.Pop(&jq).(*Job)
	if j.value != "/bax" {
		t.Error("expected value to be '/bax'")
		t.Fail()
	}
}

func FixtureJobQueue() JobQueue {
	items := map[string]int{
		"/foo": 3, "/bar": 2, "/baz": 4,
	}
	jq := make(JobQueue, len(items))
	i := 0
	for value, priority := range items {
		jq[i] = &Job{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&jq)
	return jq
}
