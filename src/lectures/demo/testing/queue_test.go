package queue

import (
	queue "coursecontent/demo/testing"
	"testing"
)

func TestAddQueue(t *testing.T) {
	q := queue.New(3) // Use the imported package name
	for i := 0; i < 3; i++ {
		if len(q.items) != i { // This won't work if items is unexported
			t.Errorf("Incorrect queue element count: %v, want %v", len(q.items), i)
		}
		if !q.Append(i) {
			t.Errorf("fail to append item %v to queue", i)
		}
	}
}
