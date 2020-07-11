package heuristic

import (
	"fmt"

	"github.com/mazrean/SnakeGame/snake/board"
)

// Manhattan マンハッタン距離を使ったヒューリスティック関数
func Manhattan(n *board.State, d board.Direction) (int, error) {
	head, err := n.Snake.Head()
	if err != nil {
		return 0, fmt.Errorf("Get Head Error: %w", err)
	}

	goal := n.Goal

	return goal.ManhattanDistance(head, d), nil
}