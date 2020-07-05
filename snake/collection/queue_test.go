package collection

import (
	"testing"
)

func TestQueuePush(t *testing.T) {
	type testQueue struct {
		queue *Queue
		value interface{}
		expect error
		description string
	}

	values := []*testQueue{
		{
			queue: &Queue{0,1,2},
			value: 3,
			expect: nil,
			description: "int queue",
		},
		{
			queue: &Queue{"a","b","c"},
			value: "d",
			expect: nil,
			description: "string queue",
		},
		{
			queue: &Queue{},
			value: 0,
			expect: nil,
			description: "empty queue",
		},
	}

	for _,v := range values {
		err := v.queue.Push(v.value)
		if err != v.expect {
			t.Fatal(v.description)
		}
	}
}

func TestQueuePop(t *testing.T) {
	type testQueue struct {
		queue *Queue
		exptSize int
		exptVal interface{}
		exptErr bool
		description string
	}

	values := []*testQueue{
		{
			queue: &Queue{0,1,2},
			exptSize: 2,
			exptVal: 0,
			exptErr: false,
			description: "int queue",
		},
		{
			queue: &Queue{"a","b","c"},
			exptSize: 2,
			exptVal: "a",
			exptErr: false,
			description: "string queue",
		},
		{
			queue: &Queue{},
			exptSize: 0,
			exptVal: nil,
			exptErr: true,
			description: "empty queue",
		},
	}

	for _,v := range values {
		res, err := v.queue.Pop()
		if (err != nil) != v.exptErr {
			t.Fatalf("Unexpect Error: %s", err.Error())
		}

		if len(*v.queue) != v.exptSize {
			t.Fatal(v.description + " size")
		}

		switch x := v.exptVal.(type) {
		case int:
			resVal, ok := res.(int)
			if !ok || resVal != x {
				t.Fatalf(v.description + " %d", resVal)
			}
		case string:
			resVal, ok := res.(string)
			if !ok || resVal != x {
				t.Fatalf(v.description + " %s", resVal)
			}
		}
	}
}

func TestQueuePeek(t *testing.T) {
	type testQueue struct {
		queue *Queue
		exptSize int
		exptVal interface{}
		exptErr bool
		description string
	}

	values := []*testQueue{
		{
			queue: &Queue{0,1,2},
			exptSize: 3,
			exptVal: 0,
			exptErr: false,
			description: "int queue",
		},
		{
			queue: &Queue{"a","b","c"},
			exptSize: 3,
			exptVal: "a",
			exptErr: false,
			description: "string queue",
		},
		{
			queue: &Queue{},
			exptSize: 0,
			exptVal: nil,
			exptErr: true,
			description: "empty queue",
		},
	}

	for _,v := range values {
		res, err := v.queue.Peek()
		if (err != nil) != v.exptErr {
			t.Fatalf("Unexpect Error: %s", err.Error())
		}

		if len(*v.queue) != v.exptSize {
			t.Fatal(v.description + " size")
		}

		switch x := v.exptVal.(type) {
		case int:
			resVal, ok := res.(int)
			if !ok || resVal != x {
				t.Fatalf(v.description + " %d", resVal)
			}
		case string:
			resVal, ok := res.(string)
			if !ok || resVal != x {
				t.Fatalf(v.description + " %s", resVal)
			}
		}
	}
}

func TestQueueSize(t *testing.T) {
	type testQueue struct {
		queue *Queue
		expect int
		description string
	}

	values := []*testQueue{
		{
			queue: &Queue{0,1},
			expect: 2,
			description: "normal int queue",
		},
		{
			queue: &Queue{},
			expect: 0,
			description: "empty queue",
		},
	}

	for _,v := range values {
		res, err := v.queue.Size()
		if err != nil {
			t.Fatalf(v.description + " Error: %s", err.Error())
		}

		if res != v.expect {
			t.Fatalf(v.description + " %d", res)
		}
	}
}

func TestQueueEmpty(t *testing.T) {
	type testQueue struct {
		queue *Queue
		expect bool
		description string
	}

	values := []*testQueue{
		{
			queue: &Queue{0,1},
			expect: false,
			description: "normal int queue",
		},
		{
			queue: &Queue{},
			expect: true,
			description: "empty queue",
		},
	}

	for _,v := range values {
		res, err := v.queue.Empty()
		if err != nil {
			t.Fatalf(v.description + " Error: %s", err.Error())
		}

		if res != v.expect {
			t.Fatalf(v.description + " %t", res)
		}
	}
}