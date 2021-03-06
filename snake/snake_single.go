package snake

import (
	"errors"
	"fmt"

	"github.com/mazrean/SnakeGame/snake/board"
	col "github.com/mazrean/SnakeGame/snake/collection"
	"github.com/mazrean/SnakeGame/snake/heuristic"
)

// NOT_FOUND 答えが見つからなかったときのエラー
var NOT_FOUND = errors.New("Answer Not Found")
var nodeCount int64 = 0

// SingleSnake シングルスレッドのヘビゲームソルバー
func SingleSnake(searchType string, s *board.State, deps ...int) (*[]board.Direction, int64, error) {
	if len(deps)==0 {
		defer func(){
			nodeCount = 0
		}()
	}
	var collection col.Collection
	switch searchType {
	case "bfs":
		collection = new(col.Queue)
	case "dfs":
		collection = new(col.Stack)
	case "iddfs":
		for i := 1;true;i++ {
			directions, _, err := SingleSnake("dfs", s, i)
			if err != nil && err != NOT_FOUND {
				return nil, 0, fmt.Errorf("DFS Error: %w", err)
			}
			if directions != nil {
				return directions, nodeCount, nil
			}
		}
	case "A*":
		collection = col.NewPriorityQueue(heuristic.Manhattan)
	default:
		return nil, 0, errors.New("Invalid Search Type")
	}

	isGoal, err := Check(s, collection)
	if err != nil {
		return nil, 0, fmt.Errorf("Check Is Goal Error: %w", err)
	}
	if isGoal {
		directions := []board.Direction{}
		return &directions, nodeCount, nil
	}
	isEmpty := false
	var nd *col.Node
	var directions *[]board.Direction
	for !isEmpty {
		nd, err = collection.Pop()
		if err != nil {
			return nil, 0, errors.New("Collection Pop Error")
		}

		isGoal, directions, err = NodeTask(nd.State, nd.Direction, collection, deps...)
		if err != nil {
			return nil, 0, fmt.Errorf("NodeTask Error: %w", err)
		}

		if isGoal {
			return directions, nodeCount, nil
		}

		isEmpty, err = collection.Empty()
		if err != nil {
			return nil, nodeCount, fmt.Errorf("Collection Empty Error: %w", err)
		}
	}

	return nil, nodeCount, NOT_FOUND
}