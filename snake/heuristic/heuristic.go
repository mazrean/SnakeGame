package heuristic

import "github.com/mazrean/SnakeGame/snake/board"

// HeuristicFunc ヒューリスティック関数
type HeuristicFunc func(*board.State, board.Direction) (int, error)