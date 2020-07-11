package board

import (
	"reflect"
	"testing"
)

func TestSnakeMove(t *testing.T) {
	type testSanke struct {
		snake *Snake
		position *Position
		expect *Snake
		response *Snake
		nextPosition *Position
		nextExpect *Snake
		description string
	}
	snake := &Snake{
		{
			X: 0,
			Y: 0,
		},
		{
			X: 0,
			Y: 1,
		},
	}

	values := []*testSanke{
		{
			snake: snake,
			position: &Position{
				X: 1,
				Y: 1,
			},
			expect: &Snake{
				{
					X: 0,
					Y: 1,
				},
				{
					X: 1,
					Y: 1,
				},
			},
			nextPosition: &Position{
				X: 1,
				Y: 0,
			},
			nextExpect: &Snake{
				{
					X: 1,
					Y: 1,
				},
				{
					X: 1,
					Y: 0,
				},
			},
			description: "normal case",
		},
		{
			snake: snake,
			position: &Position{
				X: 0,
				Y: 2,
			},
			expect: &Snake{
				{
					X: 0,
					Y: 1,
				},
				{
					X: 0,
					Y: 2,
				},
			},
			nextPosition: &Position{
				X: 1,
				Y: 2,
			},
			nextExpect: &Snake{
				{
					X: 0,
					Y: 2,
				},
				{
					X: 1,
					Y: 2,
				},
			},
			description: "second move case",
		},
	}

	for _,v := range values {
		res, err := v.snake.Move(v.position)
		if err != nil {
			t.Fatalf(v.description + " Error: %w", err)
		}

		if res == v.snake {
			t.Fatal(v.description + " Same Pointer")
		}
		if !reflect.DeepEqual(res, v.expect) {
			t.Fatal(v.description + " Invalid Value")
		}
		v.response = res
	}

	for _,v := range values {
		res, err := v.response.Move(v.nextPosition)
		if err != nil {
			t.Fatalf(v.description + " Error: %w", err)
		}

		if res == v.response {
			t.Fatal(v.description + " Same Pointer(Next)")
		}
		if !reflect.DeepEqual(res, v.nextExpect) {
			t.Fatal(v.description + " Invalid Value(Next)")
		}
	}
}