package snake

import (
	"errors"
	"fmt"

	"github.com/mazrean/SnakeGame/snake/board"
	col "github.com/mazrean/SnakeGame/snake/collection"
)

// SingleSnake シングルスレッドのヘビゲームソルバー
func SingleSnake(searchType string, s *board.State) ([]board.Direction, error) {
	cpuf, err := os.Create("cpu.prof")
	if err != nil {
		return nil, fmt.Errorf("CPU File Create Error: %w", err)
	}
	defer cpuf.Close()

	memf, err := os.Create("mem.prof")
	if err != nil {
		return nil, fmt.Errorf("Memory File Create Error: %w", err)
	}

	pprof.StartCPUProfile(cpuf)
	defer pprof.StopCPUProfile()

	defer func() {
		pprof.Lookup("heap").WriteTo(memf, 0)
		memf.Close()
	}()

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
	var nd *col.Node
	var directions []board.Direction
	for !isEmpty {
		nd, err = collection.Pop()
		if err != nil {
			return nil, errors.New("Collection Pop Error")
		}

		isGoal, directions, err = NodeTask(nd.State, nd.Direction, collection)
		if err != nil {
			return nil, fmt.Errorf("NodeTask Error: %w", err)
		}

		if isGoal {
			return directions, nil
		}

		isEmpty, err = collection.Empty()
		if err != nil {
			return nil, fmt.Errorf("Collection Empty Error: %w", err)
		}
	}

	return nil, errors.New("Answer Not Found")
}