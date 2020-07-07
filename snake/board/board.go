package board

import (
	"fmt"

	"github.com/mazrean/SnakeGame/snake/collection"
)

// Direction 移動の種類
type Direction int

const (
	// UP 上移動
	UP Direction = iota
	// DOWN 下移動
	DOWN
	// RIGHT 右移動
	RIGHT
	// LEFT 左移動
	LEFT
)

// Board 面の
type Board struct {
	Width int
	Height int
}

// State 盤面の構造体
type State struct {
	Board Board
	Goal *Position
	Snake Snake
}

// IsGoal ゴールに到達したかのチェック
func (s *State) IsGoal() (bool, error) {
	head, err := s.Snake.Head()
	if err != nil {
		return false, fmt.Errorf("Get Head Error: %w", err)
	}

	return s.Goal.equal(head), nil
}

// AbleDirections 移動可能な方向の配列を返す
func (s *State) AbleDirections() ([]Direction, error) {
	head,err := s.Snake.Head()
	if err != nil {
		return nil, fmt.Errorf("Get Head Error: %w", err)
	}

	headDirs := []struct{
			head *Position
			direction Direction
		}{
			{
				head: &Position{
					x: head.x,
					y: head.y+1,
				},
				direction: UP,
			},
			{
				head: &Position{
					x: head.x,
					y: head.y-1,
				},
				direction: DOWN,
			},
			{
				head: &Position{
					x: head.x+1,
					y: head.y,
				},
				direction: RIGHT,
			},
			{
				head: &Position{
					x: head.x-1,
					y: head.y,
				},
				direction: LEFT,
			},
		}

		directions := make([]Direction, 0, 4)
		for _,headDir := range headDirs {
			if headDir.head.x >= s.Board.Width || headDir.head.x < 0 || headDir.head.y >= s.Board.Height || headDir.head.y < 0 {
				continue
			}

			isCollide, err := s.Snake.Search(headDir.head)
			if err != nil {
				return nil, fmt.Errorf("Snake Search Error: %w", err)
			}

			if !isCollide {
				directions = append(directions, headDir.direction)
			}
		}

		return directions, nil
}

// Move 状態の遷移
func (s *State) Move(d Direction) (*State,error) {
	state := &(*s)

	nowHead,err := s.Snake.Head()
	if err != nil {
		return nil, fmt.Errorf("Get Head Error: %w", err)
	}

	newHead := &Position{
		x: nowHead.x,
		y: nowHead.y,
	}
	switch d {
	case UP:
		newHead.y--
	case DOWN:
		newHead.y++
	case RIGHT:
		newHead.x++
	case LEFT:
		newHead.x--
	}

	newSnakeQueue := collection.Queue(state.Snake)
	_, err = newSnakeQueue.Pop()
	if err != nil {
		return nil, fmt.Errorf("Queue Pop Error: %w", err)
	}
	err = newSnakeQueue.Push(newHead)
	if err != nil {
		return nil, fmt.Errorf("Queue Push Error: %w", err)
	}

	state.Snake = Snake(newSnakeQueue)

	return state, nil
}

func (s *State) String() string {
	array := make([][]string, s.Board.Height)
	for i := range array {
		lineSlice := make([]string, s.Board.Width)
		for i := range lineSlice {
			lineSlice[i] = "　"
		}
		array[i] = lineSlice
	}

	array[s.Goal.y][s.Goal.x] = "ゴ"

	for i,v := range s.Snake {
		p, ok := v.(*Position)
		if !ok {
			return ""
		}
		if i == len(s.Snake)-1 {
			array[p.y][p.x] = "頭"
		} else {
			array[p.y][p.x] = "蛇"
		}
	}

	str := make([]byte, 0, s.Board.Height*(s.Board.Width+1)*8)
	for _,v := range array {
		for _,val := range v {
			str = append(str, val...)
		}
		str = append(str, '\n')
	}

	return string(str)
}