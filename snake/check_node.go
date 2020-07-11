package snake

import (
	"fmt"

	"github.com/mazrean/SnakeGame/snake/board"
	"github.com/mazrean/SnakeGame/snake/collection"
)

// NodeTask 各Nodeで行う処理
func NodeTask(s *board.State, d board.Direction, c collection.Collection, deps ...int) (bool, *[]board.Direction, error) {
	state, err := s.Move(d)
	if err != nil {
		return false, nil, fmt.Errorf("Snake Move Error: %w", err)
	}

	isGoal, err :=Check(state, c, deps...)
	if err != nil {
		return false, nil, err
	}

	return isGoal, state.Directions, nil
}

// Check Nodeのチェック
func Check(s *board.State, c collection.Collection, deps ...int) (bool, error) {
	isGoal, err := s.IsGoal()
	if err != nil {
		return false, fmt.Errorf("Check Is Goal Error: %w", err)
	}
	if isGoal {
		return true, nil
	}

	if len(deps) != 0 && deps[0]<=s.Turn {
		return false, nil
	}
	directions, err := s.AbleDirections()
	if err != nil {
		return false, fmt.Errorf("Get Able Directions Error: %w", err)
	}

	for _,d := range directions {
		node := &collection.Node{
			State: s,
			Direction: d,
		}
		err = c.Push(node)
		if err != nil {
			return false, fmt.Errorf("Collection Push Error: %w", err)
		}
	}

	return false, nil
}