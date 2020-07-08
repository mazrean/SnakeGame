package snake

import (
	"errors"
	"fmt"

	"github.com/mazrean/SnakeGame/snake/board"
	col "github.com/mazrean/SnakeGame/snake/collection"
)

// SingleSnake シングルスレッドのヘビゲームソルバー
func SingleSnake(searchType string, s *board.State) ([]board.Direction, error) {
	var collection col.Collection
	switch searchType {
	case "bfs":
		collection = new(col.Queue)
	default:
		return nil, errors.New("Invalid Search Type")
	}

	isGoal, err := Check(s, collection)
	if err != nil {
		return nil, fmt.Errorf("Check Is Goal Error: %w", err)
	}
	if isGoal {
		return []board.Direction{}, nil
	}
	isEmpty := false
	var interfaceNode interface{}
	var nd *node
	var ok bool
	for !isEmpty {
		interfaceNode, err = collection.Pop()
		if err != nil {
			return nil, fmt.Errorf("Collection Pop Error: %w", err)
		}
		nd, ok = interfaceNode.(*node)
		if !ok {
			return nil, errors.New("Unexpected Type Parse Error")
		}

		isGoal, err = NodeTask(nd.state, nd.direction, collection)
		if err != nil {
			return nil, fmt.Errorf("NodeTask Error: %w", err)
		}

		if isGoal {
			return nd.state.Directions, nil
		}
	}

	return nil, errors.New("Answer Not Found")
}