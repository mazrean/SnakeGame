package board

import (
	"reflect"
	"testing"
)

func TestStateIsGoal(t *testing.T) {
	type testStateIsGoal struct {
		state *State
		expect bool
		description string
	}

	values := []*testStateIsGoal{
		{
			state: &State{
				Goal: &Position{
					X: 0,
					Y: 0,
				},
				Snake: &Snake{
					&Position{
						X: 1,
						Y: 1,
					},
					&Position{
						X: 1,
						Y: 0,
					},
					&Position{
						X: 0,
						Y: 0,
					},
				},
			},
			expect: true,
			description: "goaled board",
		},
		{
			state: &State{
				Goal: &Position{
					X: 0,
					Y: 0,
				},
				Snake: &Snake{
					&Position{
						X: 2,
						Y: 1,
					},
					&Position{
						X: 1,
						Y: 1,
					},
					&Position{
						X: 1,
						Y: 0,
					},
				},
			},
			expect: false,
			description: "not goaled board",
		},
		{
			state: &State{
				Goal: &Position{
					X: 1,
					Y: 0,
				},
				Snake: &Snake{
					&Position{
						X: 1,
						Y: 1,
					},
					&Position{
						X: 1,
						Y: 0,
					},
					&Position{
						X: 0,
						Y: 0,
					},
				},
			},
			expect: false,
			description: "board which body is on board",
		},
	}

	for _,v := range values {
		res, err := v.state.IsGoal()
		if err != nil {
			t.Fatalf(v.description + " Error: %s", err.Error())
		}

		if res != v.expect {
			t.Fatal(v.description + " %t", res)
		}
	}
}

func TestStateAbleDirctions(t *testing.T) {
	type testState struct {
		state *State
		expect []Direction
		description string
	}

	values := []*testState {
		{
			state: &State{
				Board: &Board{
					Width: 3,
					Height: 3,
				},
				Snake: &Snake{
					&Position{
						X: 2,
						Y: 2,
					},
					&Position{
						X: 2,
						Y: 1,
					},
					&Position{
						X: 1,
						Y: 1,
					},
				},
			},
			expect: []Direction{UP, DOWN, LEFT},
			description: "body collision detection",
		},
		{
			state: &State{
				Board: &Board{
					Width: 3,
					Height: 3,
				},
				Snake: &Snake{
					&Position{
						X: 1,
						Y: 2,
					},
					&Position{
						X: 1,
						Y: 1,
					},
					&Position{
						X: 1,
						Y: 0,
					},
				},
			},
			expect: []Direction{RIGHT,LEFT},
			description: "upeer wall collision detection",
		},
		{
			state: &State{
				Board: &Board{
					Width: 3,
					Height: 3,
				},
				Snake: &Snake{
					&Position{
						X: 1,
						Y: 0,
					},
					&Position{
						X: 1,
						Y: 1,
					},
					&Position{
						X: 0,
						Y: 1,
					},
				},
			},
			expect: []Direction{UP,DOWN},
			description: "left wall collision detection",
		},
		{
			state: &State{
				Board: &Board{
					Width: 3,
					Height: 3,
				},
				Snake: &Snake{
					&Position{
						X: 1,
						Y: 2,
					},
					&Position{
						X: 1,
						Y: 1,
					},
					&Position{
						X: 2,
						Y: 1,
					},
				},
			},
			expect: []Direction{UP,DOWN},
			description: "right wall collision detection",
		},
		{
			state: &State{
				Board: &Board{
					Width: 3,
					Height: 3,
				},
				Snake: &Snake{
					&Position{
						X: 1,
						Y: 0,
					},
					&Position{
						X: 1,
						Y: 1,
					},
					&Position{
						X: 1,
						Y: 2,
					},
				},
			},
			expect: []Direction{RIGHT,LEFT},
			description: "lower wall collision detection",
		},
	}

	for _,v := range values {
		directions, err :=v.state.AbleDirections()
		if err != nil {
			t.Fatalf(v.description + " Error: %s", err.Error())
		}

		if len(directions) != len(v.expect) {
			t.Fatalf(v.description + " Unexpected Value: %#v", v.expect)
		}
		for _,dir := range directions {
			for i,val := range v.expect {
				if val == dir {
					break
				}
				if i == len(v.expect) {
					t.Fatalf(v.description + " Value Not Found")
				}
			}
		}
	}
}

func TestStateMove(t *testing.T)  {
	type testState struct {
		state *State
		direction Direction
		expect *Snake
		response *State
		nextDirection *Direction
		nextExpect *Snake
		description string
	}

	state := &State{
		Snake: &Snake{
			&Position{
				X: 0,
				Y: 2,
			},
			&Position{
				X: 1,
				Y: 2,
			},
			&Position{
				X: 1,
				Y: 1,
			},
		},
	}

	up := UP
	right := RIGHT
	values := []*testState{
		{
			state: state,
			direction: UP,
			expect: &Snake{
				&Position{
					X: 1,
					Y: 2,
				},
				&Position{
					X: 1,
					Y: 1,
				},
				&Position{
					X: 1,
					Y: 0,
				},
			},
			nextDirection: &right,
			nextExpect: &Snake{
				&Position{
					X: 1,
					Y: 1,
				},
				&Position{
					X: 1,
					Y: 0,
				},
				&Position{
					X: 2,
					Y: 0,
				},
			},
			description: "up move",
		},
		{
			state: state,
			direction: RIGHT,
			expect: &Snake{
				&Position{
					X: 1,
					Y: 2,
				},
				&Position{
					X: 1,
					Y: 1,
				},
				&Position{
					X: 2,
					Y: 1,
				},
			},
			nextDirection: &up,
			nextExpect: &Snake{
				&Position{
					X: 1,
					Y: 1,
				},
				&Position{
					X: 2,
					Y: 1,
				},
				&Position{
					X: 2,
					Y: 0,
				},
			},
			description: "right move",
		},
		{
			state: &State{
				Snake: &Snake{
					&Position{
						X: 0,
						Y: 2,
					},
					&Position{
						X: 1,
						Y: 2,
					},
					&Position{
						X: 1,
						Y: 1,
					},
				},
			},
			direction: LEFT,
			expect: &Snake{
				&Position{
					X: 1,
					Y: 2,
				},
				&Position{
					X: 1,
					Y: 1,
				},
				&Position{
					X: 0,
					Y: 1,
				},
			},
			description: "left move",
		},
		{
			state: &State{
				Snake: &Snake{
					&Position{
						X: 0,
						Y: 0,
					},
					&Position{
						X: 1,
						Y: 0,
					},
					&Position{
						X: 1,
						Y: 1,
					},
				},
			},
			direction: DOWN,
			expect: &Snake{
				&Position{
					X: 1,
					Y: 0,
				},
				&Position{
					X: 1,
					Y: 1,
				},
				&Position{
					X: 1,
					Y: 2,
				},
			},
			description: "down move",
		},
	}

	for _,v := range values {
		res,err := v.state.Move(v.direction)
		if err != nil {
			t.Fatalf(v.description + " Error(): %s", err.Error())
		}

		if res.Snake == v.state.Snake {
			t.Fatal(v.description + " Same Pointer")
		}
		if !reflect.DeepEqual(res.Snake, v.expect) {
			t.Fatalf(v.description + " Invalid Value: %#v", res.Snake)
		}

		v.response = res
	}

	for _,v := range values {
		if v.nextDirection != nil && v.nextExpect != nil {
			t.Logf("%+v %+v", v.nextDirection, v.nextExpect)
			res,err := v.response.Move(*v.nextDirection)
			if err != nil {
				t.Fatalf(v.description + " Error: %+v", err)
			}

			if res.Snake == v.response.Snake {
				t.Fatal(v.description + " Same Pointer(Next)")
			}
			if !reflect.DeepEqual(res.Snake, v.nextExpect) {
				t.Fatalf(v.description + " Invalid Value(Next): %#v", res.Snake)
			}
		}
	}
}

func TestStateString(t *testing.T) {
	type testState struct {
		state *State
		expect string
		description string
	}
	values := []*testState{
		{
			state: &State{
				Turn: 0,
				Board: &Board{
					Width: 3,
					Height: 3,
				},
				Goal: &Position{
					X: 0,
					Y: 2,
				},
				Snake: &Snake{
					&Position{
						X: 0,
						Y: 0,
					},
					&Position{
						X: 1,
						Y: 0,
					},
					&Position{
						X: 1,
						Y: 1,
					},
				},
			},
			expect: `Turn: 0
蛇蛇　
　頭　
ゴ　　
`,
			description: "nomal state",
		},
		{
			state: &State{
				Turn: 0,
				Board: &Board{
					Width: 3,
					Height: 3,
				},
				Goal: &Position{
					X: 1,
					Y: 1,
				},
				Snake: &Snake{
					&Position{
						X: 0,
						Y: 0,
					},
					&Position{
						X: 1,
						Y: 0,
					},
					&Position{
						X: 1,
						Y: 1,
					},
				},
			},
			expect: `Turn: 0
蛇蛇　
　頭　
　　　
`,
			description: "state which head is on goal",
		},
	}

	for _,v := range values {
		res := v.state.String()

		if res != v.expect {
			t.Fatalf(v.description + " Unexpected Value: %s", res)
		}
	}
}