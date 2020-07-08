package snake

import "github.com/mazrean/SnakeGame/snake/board"

type node struct {
	state *board.State
	direction board.Direction
	directions []*board.Direction
}