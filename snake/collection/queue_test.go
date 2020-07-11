package collection

import (
	"reflect"
	"testing"

	"github.com/mazrean/SnakeGame/snake/board"
)

var nodes = []*Node{
	{
		State: &board.State{},
		Direction: board.UP,
		Directions: []*board.Direction{},
	},
	{
		State: &board.State{},
		Direction: board.DOWN,
		Directions: []*board.Direction{},
	},
	{
		State: &board.State{},
		Direction: board.RIGHT,
		Directions: []*board.Direction{},
	},
	{
		State: &board.State{},
		Direction: board.LEFT,
		Directions: []*board.Direction{},
	},
}

func TestQueuePush(t *testing.T) {
	type testQueue struct {
		queue *Queue
		value *Node
		expect error
		description string
	}

	values := []*testQueue{
		{
			queue: &Queue{nodes[0],nodes[1],nodes[2]},
			value: nodes[3],
			expect: nil,
			description: "normal queue",
		},
		{
			queue: &Queue{},
			value: nodes[0],
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
		exptVal *Node
		exptErr bool
		description string
	}

	values := []*testQueue{
		{
			queue: &Queue{nodes[0],nodes[1],nodes[2]},
			exptSize: 2,
			exptVal: nodes[0],
			exptErr: false,
			description: "normal queue",
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

		if res != v.exptVal {
			t.Fatalf(v.description + " Pointer Invalid: %+v", res)
		}
		if !reflect.DeepEqual(res, v.exptVal) {
			t.Fatalf(v.description + " Value Invalid: %+v", res)
		}
	}
}

func TestQueuePeek(t *testing.T) {
	type testQueue struct {
		queue *Queue
		exptSize int
		exptVal *Node
		exptErr bool
		description string
	}

	values := []*testQueue{
		{
			queue: &Queue{nodes[0],nodes[1],nodes[2]},
			exptSize: 3,
			exptVal: nodes[0],
			exptErr: false,
			description: "normal queue",
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
			t.Fatalf("Unexpect Error: %+v", err)
		}

		if len(*v.queue) != v.exptSize {
			t.Fatal(v.description + " size")
		}

		if res != v.exptVal {
			t.Fatalf(v.description + " Pointer Invalid: %+v", res)
		}
		if !reflect.DeepEqual(res, v.exptVal) {
			t.Fatalf(v.description + " Value Invalid: %+v", res)
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
			queue: &Queue{nodes[0],nodes[1]},
			expect: 2,
			description: "normal normal queue",
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
			queue: &Queue{nodes[0],nodes[1]},
			expect: false,
			description: "normal normal queue",
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