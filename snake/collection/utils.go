package collection

import "github.com/mazrean/SnakeGame/snake/board"

// Node 探索のノードの構造体
type Node struct {
	State *board.State
	Direction board.Direction
	Directions []*board.Direction
}